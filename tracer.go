package tracer

import (
	"net/http"

	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// NewHandler instanciates a new tracer HTTP handler.
func NewHandler(o ...otelhttp.Option) func(http.Handler) http.Handler {
	opts := []otelhttp.Option{
		otelhttp.WithPropagators(propagator.CloudTraceFormatPropagator{}),
		otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
			return r.URL.Path
		}),
	}

	opts = append(opts, o...)

	return func(next http.Handler) http.Handler {
		return otelhttp.NewHandler(next, "", opts...)
	}
}
