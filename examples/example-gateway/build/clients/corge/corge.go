// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package corgeclient

import (
	"context"
	"errors"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/uber/tchannel-go"
	"github.com/uber/zanzibar/config"
	zanzibar "github.com/uber/zanzibar/runtime"
	"github.com/uber/zanzibar/runtime/ruleengine"

	"go.uber.org/zap"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/corge/module"
	clientsIDlClientsCorgeCorge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/corge/corge"
)

// Client defines corge client interface.
type Client interface {
	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsIDlClientsCorgeCorge.Corge_EchoString_Args,
	) (string, map[string]string, error)
}

// NewClient returns a new TChannel client for service corge.
func NewClient(deps *module.Dependencies) Client {
	serviceName := deps.Default.Config.MustGetString("clients.corge.serviceName")
	var routingKey string
	if deps.Default.Config.ContainsKey("clients.corge.routingKey") {
		routingKey = deps.Default.Config.MustGetString("clients.corge.routingKey")
	}
	var requestUUIDHeaderKey string
	if deps.Default.Config.ContainsKey("tchannel.clients.requestUUIDHeaderKey") {
		requestUUIDHeaderKey = deps.Default.Config.MustGetString("tchannel.clients.requestUUIDHeaderKey")
	}
	sc := deps.Default.Channel.GetSubChannel(serviceName, tchannel.Isolated)

	ip := deps.Default.Config.MustGetString("sidecarRouter.default.tchannel.ip")
	port := deps.Default.Config.MustGetInt("sidecarRouter.default.tchannel.port")
	sc.Peers().Add(ip + ":" + strconv.Itoa(int(port)))

	/*Ex:
	{
	  "clients.rider-presentation.alternates": {
		"routingConfigs": [
		  {
			"headerName": "x-test-env",
			"headerValue": "*",
			"serviceName": "testservice"
		  },
		  {
			"headerName": "x-container",
			"headerValue": "container*",
			"serviceName": "relayer"
		  }
		],
		"servicesDetail": {
		  "testservice": {
			"ip": "127.0.0.1",
			"port": 5000
		  },
		  "relayer": {
			"ip": "127.0.0.1",
			"port": 12000
		  }
		}
	  }
	}*/
	var re ruleengine.RuleEngine
	var headerPatterns []string
	altChannelMap := make(map[string]*tchannel.SubChannel)
	headerPatterns, re = initializeDynamicChannel(deps, headerPatterns, altChannelMap, re)

	timeoutVal := int(deps.Default.Config.MustGetInt("clients.corge.timeout"))
	timeout := time.Millisecond * time.Duration(
		timeoutVal,
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		deps.Default.Config.MustGetInt("clients.corge.timeoutPerAttempt"),
	)

	methodNames := map[string]string{
		"Corge::echoString": "EchoString",
	}

	circuitBreakerDisabled := configureCicruitBreaker(deps, timeoutVal)

	client := zanzibar.NewTChannelClientContext(
		deps.Default.Channel,
		deps.Default.Logger,
		deps.Default.ContextMetrics,
		deps.Default.ContextExtractor,
		&zanzibar.TChannelClientOption{
			ServiceName:          serviceName,
			ClientID:             "corge",
			MethodNames:          methodNames,
			Timeout:              timeout,
			TimeoutPerAttempt:    timeoutPerAttempt,
			RoutingKey:           &routingKey,
			RuleEngine:           re,
			HeaderPatterns:       headerPatterns,
			RequestUUIDHeaderKey: requestUUIDHeaderKey,
			AltChannelMap:        altChannelMap,
		},
	)

	return &corgeClient{
		client:                 client,
		circuitBreakerDisabled: circuitBreakerDisabled,
	}
}

func initializeDynamicChannel(deps *module.Dependencies, headerPatterns []string, altChannelMap map[string]*tchannel.SubChannel, re ruleengine.RuleEngine) ([]string, ruleengine.RuleEngine) {
	if deps.Default.Config.ContainsKey("clients.corge.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.corge.alternates", &alternateServiceDetail)

		ruleWrapper := ruleengine.RuleWrapper{}
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			ruleValue := []string{routingConfig.ServiceName}
			rd := routingConfig.RoutingDelegate
			if rd != nil {
				ruleValue = append(ruleValue, *rd)
			}
			rawRule := ruleengine.RawRule{Patterns: []string{textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName),
				strings.ToLower(routingConfig.HeaderValue)}, Value: ruleValue}
			headerPatterns = append(headerPatterns, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
			ruleWrapper.Rules = append(ruleWrapper.Rules, rawRule)

			scAlt := deps.Default.Channel.GetSubChannel(routingConfig.ServiceName, tchannel.Isolated)
			serviceRouting, ok := alternateServiceDetail.ServicesDetailMap[routingConfig.ServiceName]
			if !ok {
				panic("service routing mapping incorrect for service: " + routingConfig.ServiceName)
			}
			scAlt.Peers().Add(serviceRouting.IP + ":" + strconv.Itoa(serviceRouting.Port))
			altChannelMap[routingConfig.ServiceName] = scAlt
		}

		re = ruleengine.NewRuleEngine(ruleWrapper)
	}
	return headerPatterns, re
}

func configureCicruitBreaker(deps *module.Dependencies, timeoutVal int) bool {
	// circuitBreakerDisabled sets whether circuit-breaker should be disabled
	circuitBreakerDisabled := false
	if deps.Default.Config.ContainsKey("clients.corge.circuitBreakerDisabled") {
		circuitBreakerDisabled = deps.Default.Config.MustGetBoolean("clients.corge.circuitBreakerDisabled")
	}
	// sleepWindowInMilliseconds sets the amount of time, after tripping the circuit,
	// to reject requests before allowing attempts again to determine if the circuit should again be closed
	sleepWindowInMilliseconds := 5000
	if deps.Default.Config.ContainsKey("clients.corge.sleepWindowInMilliseconds") {
		sleepWindowInMilliseconds = int(deps.Default.Config.MustGetInt("clients.corge.sleepWindowInMilliseconds"))
	}
	// maxConcurrentRequests sets how many requests can be run at the same time, beyond which requests are rejected
	maxConcurrentRequests := 20
	if deps.Default.Config.ContainsKey("clients.corge.maxConcurrentRequests") {
		maxConcurrentRequests = int(deps.Default.Config.MustGetInt("clients.corge.maxConcurrentRequests"))
	}
	// errorPercentThreshold sets the error percentage at or above which the circuit should trip open
	errorPercentThreshold := 20
	if deps.Default.Config.ContainsKey("clients.corge.errorPercentThreshold") {
		errorPercentThreshold = int(deps.Default.Config.MustGetInt("clients.corge.errorPercentThreshold"))
	}
	// requestVolumeThreshold sets a minimum number of requests that will trip the circuit in a rolling window of 10s
	// For example, if the value is 20, then if only 19 requests are received in the rolling window of 10 seconds
	// the circuit will not trip open even if all 19 failed.
	requestVolumeThreshold := 20
	if deps.Default.Config.ContainsKey("clients.corge.requestVolumeThreshold") {
		requestVolumeThreshold = int(deps.Default.Config.MustGetInt("clients.corge.requestVolumeThreshold"))
	}
	if !circuitBreakerDisabled {
		hystrix.ConfigureCommand("corge", hystrix.CommandConfig{
			MaxConcurrentRequests:  maxConcurrentRequests,
			ErrorPercentThreshold:  errorPercentThreshold,
			SleepWindow:            sleepWindowInMilliseconds,
			RequestVolumeThreshold: requestVolumeThreshold,
			Timeout:                timeoutVal,
		})
	}
	return circuitBreakerDisabled
}

// corgeClient is the TChannel client for downstream service.
type corgeClient struct {
	client                 *zanzibar.TChannelClient
	circuitBreakerDisabled bool
}

// EchoString is a client RPC call for method "Corge::echoString"
func (c *corgeClient) EchoString(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsIDlClientsCorgeCorge.Corge_EchoString_Args,
) (string, map[string]string, error) {
	var result clientsIDlClientsCorgeCorge.Corge_EchoString_Result
	var resp string

	logger := c.client.Loggers["Corge::echoString"]

	var success bool
	respHeaders := make(map[string]string)
	var err error
	if c.circuitBreakerDisabled {
		success, respHeaders, err = c.client.Call(
			ctx, "Corge", "echoString", reqHeaders, args, &result,
		)
	} else {
		// We want hystrix ckt-breaker to count errors only for system issues
		var clientErr error
		err = hystrix.DoC(ctx, "corge", func(ctx context.Context) error {
			success, respHeaders, clientErr = c.client.Call(
				ctx, "Corge", "echoString", reqHeaders, args, &result,
			)
			if _, isSysErr := clientErr.(tchannel.SystemError); !isSysErr {
				// Declare ok if it is not a system-error
				return nil
			}
			return clientErr
		}, nil)
		if err == nil {
			// ckt-breaker was ok, bubble up client error if set
			err = clientErr
		}
	}

	if err == nil && !success {
		switch {
		default:
			err = errors.New("corgeClient received no result or unknown exception for EchoString")
		}
	}
	if err != nil {
		logger.Warn("Client failure: TChannel client call returned error", zap.Error(err))
		return resp, respHeaders, err
	}

	resp, err = clientsIDlClientsCorgeCorge.Corge_EchoString_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Client failure: unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}
