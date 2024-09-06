package chonker

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/wazazaby/vimebu/v2"
)

// StatsForNerds exposes Prometheus metrics for chonker requests.
// Metric names are prefixed with "chonker_".
// Metrics are labeled with request host URLs.
//
// The following metrics are exposed for a request to https://example.com:
//
// chonker_http_requests_fetching{host="example.com"}
// chonker_http_requests_total{host="example.com"}
// chonker_http_requests_total{host="example.com",range="false"}
// chonker_http_request_chunks_fetching{host="example.com",stage="do"}
// chonker_http_request_chunks_fetching{host="example.com",stage="copy"}
// chonker_http_request_chunks_total{host="example.com"}
// chonker_http_request_chunk_duration_seconds{host="example.com"}
// chonker_http_request_chunk_bytes{host="example.com"}
//
// You can surface these metrics in your application using the
// [metrics.RegisterSet] function.
//
// [metrics.RegisterSet]: https://pkg.go.dev/github.com/VictoriaMetrics/metrics#RegisterSet
var StatsForNerds = metrics.NewSet()

type hostMetrics struct {
	// requestsFetching is the number of currently active requests to a host.
	requestsFetching *metrics.Gauge
	// requestsTotal is the total number of requests completed to a host.
	requestsTotal *metrics.Counter
	// requestsTotalSansRange is the total number of requests completed to a host
	// that did not use range requests.
	requestsTotalSansRange *metrics.Counter
	// requestChunksFetching is the number of currently active request chunks to a host.
	requestChunksFetchingStageDo   *metrics.Gauge
	requestChunksFetchingStageCopy *metrics.Gauge
	// requestChunksTotal is the total number of request chunks completed to a host.
	requestChunksTotal *metrics.Counter
	// requestChunkDurationSeconds measures the duration of request chunks to a host.
	requestChunkDurationSeconds *metrics.Histogram
	// requestChunkBytes measures the number of bytes fetched in request chunks to a host.
	requestChunkBytes *metrics.Histogram
}

func getHostMetrics(host string) *hostMetrics {
	return &hostMetrics{
		requestsFetching: vimebu.Metric("chonker_http_requests_fetching").
			LabelString("host", host).
			GetOrCreateGaugeInSet(StatsForNerds, nil),
		requestsTotal: vimebu.Metric("chonker_http_requests_total").
			LabelString("host", host).
			GetOrCreateCounterInSet(StatsForNerds),
		requestsTotalSansRange: vimebu.Metric("chonker_http_requests_total").
			LabelString("host", host).
			LabelString("range", "false").
			GetOrCreateCounterInSet(StatsForNerds),
		requestChunksFetchingStageDo: vimebu.Metric("chonker_http_request_chunks_fetching").
			LabelString("host", host).
			LabelString("stage", "do").
			GetOrCreateGaugeInSet(StatsForNerds, nil),
		requestChunksFetchingStageCopy: vimebu.Metric("chonker_http_request_chunks_fetching").
			LabelString("host", host).
			LabelString("stage", "copy").
			GetOrCreateGaugeInSet(StatsForNerds, nil),
		requestChunksTotal: vimebu.Metric("chonker_http_request_chunks_total").
			LabelString("host", host).
			GetOrCreateCounterInSet(StatsForNerds),
		requestChunkDurationSeconds: vimebu.Metric("chonker_http_request_chunk_duration_seconds").
			LabelString("host", host).
			GetOrCreateHistogramInSet(StatsForNerds),
		requestChunkBytes: vimebu.Metric("chonker_http_request_chunk_bytes").
			LabelString("host", host).
			GetOrCreateHistogramInSet(StatsForNerds),
	}
}
