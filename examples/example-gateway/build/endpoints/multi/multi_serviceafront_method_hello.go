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

package multiendpoint

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi/workflow"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi/module"
)

// ServiceAFrontHelloHandler is the handler for "/multi/serviceA_f/hello"
type ServiceAFrontHelloHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewServiceAFrontHelloHandler creates a handler
func NewServiceAFrontHelloHandler(deps *module.Dependencies) *ServiceAFrontHelloHandler {
	handler := &ServiceAFrontHelloHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope, deps.Default.Tracer,
		"multi", "helloA",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *ServiceAFrontHelloHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Register(
		"GET", "/multi/serviceA_f/hello",
		h.endpoint,
	)
}

// HandleRequest handles "/multi/serviceA_f/hello".
func (h *ServiceAFrontHelloHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	// log endpoint request to downstream services
	if ce := req.Logger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		req.Logger.Debug("endpoint request to downstream", zfields...)
	}

	w := workflow.NewServiceAFrontHelloWorkflow(h.Dependencies)
	if span := req.GetSpan(); span != nil {
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	response, cliRespHeaders, err := w.Handle(ctx, req.Header)
	if err != nil {
		res.SendError(500, "Unexpected server error", err)
		return

	}

	bytes, err := json.Marshal(response)
	if err != nil {
		res.SendError(500, "Unexpected server error", errors.Wrap(err, "Unable to marshal resp json"))
		return
	}
	res.WriteJSONBytes(200, cliRespHeaders, bytes)
}
