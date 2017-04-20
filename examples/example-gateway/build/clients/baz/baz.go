// Package bazClient is generated code used to make or handle TChannel calls using Thrift.
package bazClient

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/uber/zanzibar/runtime"

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanBaz is the interface that defines the server handler and client interface.
type TChanBaz interface {
	Call(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBazBaz.SimpleService_Call_Args,
	) (
		*clientsBazBaz.BazResponse,
		map[string]string,
		error,
	)

	Simple(
		ctx context.Context,
		reqHeaders map[string]string,
	) (
		map[string]string,
		error,
	)

	SimpleFuture(
		ctx context.Context,
		reqHeaders map[string]string,
	) (
		map[string]string,
		error,
	)
}

// NewClient returns a new TChannel client for service baz.
func NewClient(gateway *zanzibar.Gateway) *BazClient {
	serviceName := gateway.Config.MustGetString("clients.baz.serviceName")
	sc := gateway.Channel.GetSubChannel(serviceName)

	ip := gateway.Config.MustGetString("clients.baz.ip")
	port := gateway.Config.MustGetInt("clients.baz.port")
	sc.Peers().Add(ip + ":" + strconv.Itoa(int(port)))

	timeout := time.Millisecond * time.Duration(
		gateway.Config.MustGetInt("clients.baz.timeout"),
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		gateway.Config.MustGetInt("clients.baz.timeoutPerAttempt"),
	)

	client := zanzibar.NewTChannelClient(gateway.Channel,
		&zanzibar.TChannelClientOption{
			ServiceName:       serviceName,
			Timeout:           timeout,
			TimeoutPerAttempt: timeoutPerAttempt,
		},
	)

	return &BazClient{
		thriftService: "SimpleService",
		client:        client,
	}
}

// BazClient is the TChannel client for downstream service.
type BazClient struct {
	thriftService string
	client        zanzibar.TChanClient
}

// Call ...
func (c *BazClient) Call(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Call_Args,
) (*clientsBazBaz.BazResponse, map[string]string, error) {
	var result clientsBazBaz.SimpleService_Call_Result

	success, respHeaders, err := c.client.Call(
		ctx, c.thriftService, "Call", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("received no result or unknown exception for Call")
		}
	}
	if err != nil {
		return nil, nil, err
	}

	resp, err := clientsBazBaz.SimpleService_Call_Helper.UnwrapResponse(&result)

	return resp, respHeaders, err
}

// Simple ...
func (c *BazClient) Simple(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_Simple_Result

	args := &clientsBazBaz.SimpleService_Simple_Args{}
	success, respHeaders, err := c.client.Call(
		ctx, c.thriftService, "Simple", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.SimpleErr != nil:
			err = result.SimpleErr
		default:
			err = errors.New("received no result or unknown exception for SimpleService")
		}
	}
	if err != nil {
		return nil, err
	}

	return respHeaders, err
}

// SimpleFuture ...
func (c *BazClient) SimpleFuture(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error) {
	var result clientsBazBaz.SimpleService_SimpleFuture_Result

	args := &clientsBazBaz.SimpleService_SimpleFuture_Args{}
	success, respHeaders, err := c.client.Call(
		ctx, c.thriftService, "SimpleFuture", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		case result.SimpleErr != nil:
			err = result.SimpleErr
		case result.NewErr != nil:
			err = result.NewErr
		default:
			err = errors.New("received no result or unknown exception for SimpleService")
		}
	}
	if err != nil {
		return nil, err
	}

	return respHeaders, err
}
