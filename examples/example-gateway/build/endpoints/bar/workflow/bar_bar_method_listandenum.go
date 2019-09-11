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

package workflow

import (
	"context"

	"github.com/uber/zanzibar/codegen"
	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	"go.uber.org/zap"
)

// BarListAndEnumWorkflow defines the interface for BarListAndEnum workflow
type BarListAndEnumWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBarBar.Bar_ListAndEnum_Args,
	) (string, zanzibar.Header, error)
}

// NewBarListAndEnumWorkflow creates a workflow
func NewBarListAndEnumWorkflow(deps *module.Dependencies) BarListAndEnumWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.bar.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.bar.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, routingConfig.HeaderName)
		}
	}

	return &barListAndEnumWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// barListAndEnumWorkflow calls thrift client Bar.ListAndEnum
type barListAndEnumWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w barListAndEnumWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ListAndEnum_Args,
) (string, zanzibar.Header, error) {
	clientRequest := convertToListAndEnumClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	for _, whitelistedHeader := range w.whitelistedDynamicHeaders {
		transformedHeaderName := codegen.CamelCase(whitelistedHeader)
		headerVal, ok := reqHeaders.Get(transformedHeaderName)
		if ok {
			clientHeaders[transformedHeaderName] = headerVal
		}
	}

	clientRespBody, _, err := w.Clients.Bar.ListAndEnum(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertListAndEnumBarException(
				errValue,
			)

			return "", nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Bar"),
			)

			return "", nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertBarListAndEnumClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToListAndEnumClientRequest(in *endpointsBarBar.Bar_ListAndEnum_Args) *clientsBarBar.Bar_ListAndEnum_Args {
	out := &clientsBarBar.Bar_ListAndEnum_Args{}

	out.DemoIds = make([]string, len(in.DemoIds))
	for index1, value2 := range in.DemoIds {
		out.DemoIds[index1] = string(value2)
	}
	out.DemoType = (*clientsBarBar.DemoType)(in.DemoType)

	return out
}

func convertListAndEnumBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}

func convertBarListAndEnumClientResponse(in string) string {
	out := in

	return out
}
