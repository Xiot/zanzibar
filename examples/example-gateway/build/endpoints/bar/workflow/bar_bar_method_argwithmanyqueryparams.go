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

// BarArgWithManyQueryParamsWorkflow defines the interface for BarArgWithManyQueryParams workflow
type BarArgWithManyQueryParamsWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBarBar.Bar_ArgWithManyQueryParams_Args,
	) (*endpointsBarBar.BarResponse, zanzibar.Header, error)
}

// NewBarArgWithManyQueryParamsWorkflow creates a workflow
func NewBarArgWithManyQueryParamsWorkflow(deps *module.Dependencies) BarArgWithManyQueryParamsWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients.bar.alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients.bar.alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, routingConfig.HeaderName)
		}
	}

	return &barArgWithManyQueryParamsWorkflow{
		Clients:                   deps.Client,
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// barArgWithManyQueryParamsWorkflow calls thrift client Bar.ArgWithManyQueryParams
type barArgWithManyQueryParamsWorkflow struct {
	Clients                   *module.ClientDependencies
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// Handle calls thrift client.
func (w barArgWithManyQueryParamsWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBarBar.Bar_ArgWithManyQueryParams_Args,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {
	clientRequest := convertToArgWithManyQueryParamsClientRequest(r)

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

	clientRespBody, _, err := w.Clients.Bar.ArgWithManyQueryParams(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Client failure: could not make client request",
				zap.Error(errValue),
				zap.String("client", "Bar"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertBarArgWithManyQueryParamsClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToArgWithManyQueryParamsClientRequest(in *endpointsBarBar.Bar_ArgWithManyQueryParams_Args) *clientsBarBar.Bar_ArgWithManyQueryParams_Args {
	out := &clientsBarBar.Bar_ArgWithManyQueryParams_Args{}

	out.AStr = string(in.AStr)
	out.AnOptStr = (*string)(in.AnOptStr)
	out.ABool = bool(in.ABool)
	out.AnOptBool = (*bool)(in.AnOptBool)
	out.AInt8 = int8(in.AInt8)
	out.AnOptInt8 = (*int8)(in.AnOptInt8)
	out.AInt16 = int16(in.AInt16)
	out.AnOptInt16 = (*int16)(in.AnOptInt16)
	out.AInt32 = int32(in.AInt32)
	out.AnOptInt32 = (*int32)(in.AnOptInt32)
	out.AInt64 = int64(in.AInt64)
	out.AnOptInt64 = (*int64)(in.AnOptInt64)
	out.AFloat64 = float64(in.AFloat64)
	out.AnOptFloat64 = (*float64)(in.AnOptFloat64)
	out.AUUID = clientsBarBar.UUID(in.AUUID)
	out.AnOptUUID = (*clientsBarBar.UUID)(in.AnOptUUID)
	out.AListUUID = make([]clientsBarBar.UUID, len(in.AListUUID))
	for index1, value2 := range in.AListUUID {
		out.AListUUID[index1] = clientsBarBar.UUID(value2)
	}
	out.AnOptListUUID = make([]clientsBarBar.UUID, len(in.AnOptListUUID))
	for index3, value4 := range in.AnOptListUUID {
		out.AnOptListUUID[index3] = clientsBarBar.UUID(value4)
	}
	out.AStringList = make([]string, len(in.AStringList))
	for index5, value6 := range in.AStringList {
		out.AStringList[index5] = string(value6)
	}
	out.AnOptStringList = make([]string, len(in.AnOptStringList))
	for index7, value8 := range in.AnOptStringList {
		out.AnOptStringList[index7] = string(value8)
	}
	out.AUUIDList = make([]clientsBarBar.UUID, len(in.AUUIDList))
	for index9, value10 := range in.AUUIDList {
		out.AUUIDList[index9] = clientsBarBar.UUID(value10)
	}
	out.AnOptUUIDList = make([]clientsBarBar.UUID, len(in.AnOptUUIDList))
	for index11, value12 := range in.AnOptUUIDList {
		out.AnOptUUIDList[index11] = clientsBarBar.UUID(value12)
	}
	out.ATs = clientsBarBar.Timestamp(in.ATs)
	out.AnOptTs = (*clientsBarBar.Timestamp)(in.AnOptTs)

	return out
}

func convertBarArgWithManyQueryParamsClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key3, value4 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key3] = int32(value4)
	}
	out.BinaryField = []byte(in.BinaryField)
	var convertBarResponseHelper5 func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse)
	convertBarResponseHelper5 = func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse) {
		if in != nil {
			out = &endpointsBarBar.BarResponse{}
			out.StringField = string(in.StringField)
			out.IntWithRange = int32(in.IntWithRange)
			out.IntWithoutRange = int32(in.IntWithoutRange)
			out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
			for key6, value7 := range in.MapIntWithRange {
				out.MapIntWithRange[endpointsBarBar.UUID(key6)] = int32(value7)
			}
			out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
			for key8, value9 := range in.MapIntWithoutRange {
				out.MapIntWithoutRange[key8] = int32(value9)
			}
			out.BinaryField = []byte(in.BinaryField)
			out.NextResponse = convertBarResponseHelper5(in.NextResponse)
		} else {
			out = nil
		}
		return
	}
	out.NextResponse = convertBarResponseHelper5(in.NextResponse)

	return out
}
