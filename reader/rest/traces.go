package rest

import (
	"context"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/bygui86/go-timescaledb/reader/logging"
)

func retrieveSpanAndCtx(request *http.Request, operationName string) (opentracing.Span, context.Context) {
	clientSpanContext, clientSpanErr := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
	if clientSpanErr != nil {
		logging.SugaredLog.Debugf("%s", clientSpanErr.Error())
	}

	// Create the span referring to the RPC client if available.
	// If clientSpanContext == nil, a root span will be created.
	span := opentracing.StartSpan(operationName, ext.RPCServerOption(clientSpanContext))
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	return span, ctx
}
