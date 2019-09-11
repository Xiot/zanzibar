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

	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceGetProfileWorkflow defines the interface for SimpleServiceGetProfile workflow
type SimpleServiceGetProfileWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBazBaz.SimpleService_GetProfile_Args,
	) (*endpointsBazBaz.GetProfileResponse, zanzibar.Header, error)
}

// NewSimpleServiceGetProfileWorkflow creates a workflow
func NewSimpleServiceGetProfileWorkflow(deps *module.Dependencies) SimpleServiceGetProfileWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.baz.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.baz.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, routingConfig.HeaderName)
		}
	}

	return &simpleServiceGetProfileWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// simpleServiceGetProfileWorkflow calls thrift client Baz.GetProfile
type simpleServiceGetProfileWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w simpleServiceGetProfileWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_GetProfile_Args,
) (*endpointsBazBaz.GetProfileResponse, zanzibar.Header, error) {
	clientRequest := convertToGetProfileClientRequest(r)

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

	clientRespBody, _, err := w.Clients.Baz.GetProfile(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertGetProfileAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceGetProfileClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToGetProfileClientRequest(in *endpointsBazBaz.SimpleService_GetProfile_Args) *clientsBazBaz.SimpleService_GetProfile_Args {
	out := &clientsBazBaz.SimpleService_GetProfile_Args{}

	if in.Request != nil {
		out.Request = &clientsBazBaz.GetProfileRequest{}
		out.Request.Target = clientsBazBaz.UUID(in.Request.Target)
	} else {
		out.Request = nil
	}

	return out
}

func convertGetProfileAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}

func convertSimpleServiceGetProfileClientResponse(in *clientsBazBaz.GetProfileResponse) *endpointsBazBaz.GetProfileResponse {
	out := &endpointsBazBaz.GetProfileResponse{}

	out.Payloads = make([]*endpointsBazBaz.Profile, len(in.Payloads))
	for index1, value2 := range in.Payloads {
		if value2 != nil {
			out.Payloads[index1] = &endpointsBazBaz.Profile{}
			if in.Payloads[index1].Recur1 != nil {
				out.Payloads[index1].Recur1 = &endpointsBazBaz.Recur1{}
				out.Payloads[index1].Recur1.Field1 = make(map[endpointsBazBaz.UUID]*endpointsBazBaz.Recur2, len(in.Payloads[index1].Recur1.Field1))
				for key3, value4 := range in.Payloads[index1].Recur1.Field1 {
					if value4 != nil {
						out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)] = &endpointsBazBaz.Recur2{}
						if in.Payloads[index1].Recur1.Field1[key3].Field21 != nil {
							out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)].Field21 = &endpointsBazBaz.Recur3{}
							if in.Payloads[index1] != nil && in.Payloads[index1].Recur1 != nil && in.Payloads[index1].Recur1.Field1[key3] != nil && in.Payloads[index1].Recur1.Field1[key3].Field21 != nil {
								out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)].Field21.Field31 = endpointsBazBaz.UUID(in.Payloads[index1].Recur1.Field1[key3].Field21.Field31)
							}
						} else {
							out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)].Field21 = nil
						}
						if in.Payloads[index1].Recur1.Field1[key3].Field22 != nil {
							out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)].Field22 = &endpointsBazBaz.Recur3{}
							if in.Payloads[index1] != nil && in.Payloads[index1].Recur1 != nil && in.Payloads[index1].Recur1.Field1[key3] != nil && in.Payloads[index1].Recur1.Field1[key3].Field22 != nil {
								out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)].Field22.Field31 = endpointsBazBaz.UUID(in.Payloads[index1].Recur1.Field1[key3].Field22.Field31)
							}
						} else {
							out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)].Field22 = nil
						}
					} else {
						out.Payloads[index1].Recur1.Field1[endpointsBazBaz.UUID(key3)] = nil
					}
				}
			} else {
				out.Payloads[index1].Recur1 = nil
			}
		} else {
			out.Payloads[index1] = nil
		}
	}

	return out
}
