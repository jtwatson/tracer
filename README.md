# tracer

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/jtwatson/tracer)

**tracer** is a http middleware that uses [OpenTelemetry](https://opentelemetry.io/docs/instrumentation/go/) to send trace data to [Cloud Trace](https://cloud.google.com/trace/). It configures OpenTelemetry to propagate traces across application boundaries, adds read/write events, and provides a default span name based upon request URL.
