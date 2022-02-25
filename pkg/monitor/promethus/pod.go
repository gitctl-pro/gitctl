package promethus

import (
	"fmt"
	"strings"
)

var seriesKeys = []string{
	"cluster",
	"namespace",
	"create_by_name",
	"host_ip",
	"pod",
	"pod_ip",
	"node",
}

var queryKeys = []string{
	"cluster",
	"namespace",
	"pod",
	"container",
}

var podInfoSeries = "kube_pod_info{create_by_name=~'$name-*', $query}"

var metrics = map[string]string{
	"container_network_transmit":               `sort_desc(sum by (pod) (irate(container_network_transmit_bytes_total{$query}[5m])))`,
	"container_network_receive":                `sort_desc(sum by (pod) (irate(container_network_receive_bytes_total{$query}[5m])))`,
	"container_receive_errors_total":           `(sum(irate(container_network_receive_errors_total{$query}[5m])) by (pod))`,
	"container_receive_packets_total":          `(sum(irate(container_network_receive_packets_total{$query}[5m])) by (pod))`,
	"container_receive_packets_dropped_total":  `(sum(irate(container_network_receive_packets_dropped_total{$query}[5m])) by (pod))`,
	"container_transmit_errors_total":          `(sum(irate(container_network_transmit_errors_total{$query}[5m])) by (pod))`,
	"container_transmit_packets_total":         `(sum(irate(container_network_transmit_packets_total{$query}[5m])) by (pod))`,
	"container_transmit_packets_dropped_total": `(sum(irate(container_network_transmit_packets_dropped_total{$query}[5m])) by (pod))`,
	"container_memory_usage":                   `sum by(pod) (container_memory_usage_bytes{$query})`,
	"container_memory_cache":                   `sum by(pod) (container_memory_cache{$query})`,
	"container_memory_requests":                `sum by(pod) (kube_pod_container_resource_requests_memory_bytes{$query})`,
	"container_memory_limits":                  `sum by(pod) (kube_pod_container_resource_limits_memory_bytes{$query})`,
	"container_cpu_usage":                      `sum by(pod) (irate(container_cpu_usage_seconds_total{$query})`,
	"container_cpu_requests":                   `sum by(pod) (kube_pod_container_resource_requests_cpu_cores{$query})`,
	"container_cpu_limits":                     `sum by(pod) (kube_pod_container_resource_limits_cpu_cores{$query})`,
}

func PodQueryExp(metric string, querys map[string]string) string {
	query := encodeQuery(queryKeys, querys)
	exp := strings.NewReplacer("$query", query).Replace(metrics[metric])
	return exp
}

func PodSeriesExp(name string, querys map[string]string) string {
	query := encodeQuery(seriesKeys, querys)
	exp := strings.NewReplacer("$name", name, "$query", query).Replace(podInfoSeries)
	return exp
}

func encodeQuery(keys []string, querys map[string]string) string {
	quote := func(k, v string) string {
		return fmt.Sprintf("%s='%s'")
	}
	var buf strings.Builder
	for key, value := range querys {
		for _, k := range keys {
			if key == k && len(value) > 0 {
				buf.WriteString(quote(key, value))
				if buf.Len() > 0 {
					buf.WriteByte(',')
				}
			}
		}
	}
	return buf.String()
}
