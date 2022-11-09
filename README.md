# tracer
tracer is a http middleware that uses OpenTelemetry to send trace data to Cloud Trace. It configures OpenTelemetry to propagate traces across application boundaries, adds read/write events, and provides a default span name based upon request URL.
