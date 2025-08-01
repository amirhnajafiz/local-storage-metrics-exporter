package metrics

import "github.com/prometheus/client_golang/prometheus"

// constant values for namespace and subsystem
const (
	NS = "KES"
	SS = "exporter"
)

// Metrics holds the Prometheus metrics for the exporter.
type Metrics struct {
	ephemeralStorageAvailableBytes *prometheus.GaugeVec
	ephemeralStorageCapacityBytes  *prometheus.GaugeVec
	ephemeralStorageUsageBytes     *prometheus.GaugeVec
	ephemeralStorageInodes         *prometheus.GaugeVec
	ephemeralStorageInodesFree     *prometheus.GaugeVec
	ephemeralStorageInodesUsed     *prometheus.GaugeVec
	containerMemoryAvailableBytes  *prometheus.GaugeVec
	containerMemoryCapacityBytes   *prometheus.GaugeVec
	containerMemoryUsageBytes      *prometheus.GaugeVec
	containerRootfsAvailableBytes  *prometheus.GaugeVec
	containerRootfsCapacityBytes   *prometheus.GaugeVec
	containerRootfsUsageBytes      *prometheus.GaugeVec
	containerRootfsInodes          *prometheus.GaugeVec
	containerRootfsInodesFree      *prometheus.GaugeVec
	containerRootfsInodesUsed      *prometheus.GaugeVec
	containerLogsAvailableBytes    *prometheus.GaugeVec
	containerLogsCapacityBytes     *prometheus.GaugeVec
	containerLogsUsageBytes        *prometheus.GaugeVec
	containerLogsInodes            *prometheus.GaugeVec
	containerLogsInodesFree        *prometheus.GaugeVec
	containerLogsInodesUsed        *prometheus.GaugeVec
}

// NewMetrics initializes and registers the Prometheus metrics for the exporter.
func NewMetrics() (*Metrics, error) {
	// create Prometheus metrics using the newGaugeVec helper function in utils.go
	return &Metrics{
		ephemeralStorageAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_available_bytes",
			Help:      "Ephemeral storage available in bytes",
		}, []string{"pod", "namespace", "node", "uid"}),
		ephemeralStorageCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_capacity_bytes",
			Help:      "Ephemeral storage capacity in bytes",
		}, []string{"pod", "namespace", "node", "uid"}),
		ephemeralStorageUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_usage_bytes",
			Help:      "Ephemeral storage usage in bytes",
		}, []string{"pod", "namespace", "node", "uid"}),
		ephemeralStorageInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_inodes",
			Help:      "Ephemeral storage inodes",
		}, []string{"pod", "namespace", "node", "uid"}),
		ephemeralStorageInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_inodes_free",
			Help:      "Ephemeral storage free inodes",
		}, []string{"pod", "namespace", "node", "uid"}),
		ephemeralStorageInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "ephemeral_storage_inodes_used",
			Help:      "Ephemeral storage used inodes",
		}, []string{"pod", "namespace", "node", "uid"}),
		containerMemoryAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_memory_available_bytes",
			Help:      "Container memory available in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerMemoryCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_memory_capacity_bytes",
			Help:      "Container memory capacity in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerMemoryUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_memory_usage_bytes",
			Help:      "Container memory usage in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerRootfsAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_rootfs_available_bytes",
			Help:      "Container root filesystem available in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerRootfsCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_rootfs_capacity_bytes",
			Help:      "Container root filesystem capacity in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerRootfsUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_rootfs_usage_bytes",
			Help:      "Container root filesystem usage in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerRootfsInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_rootfs_inodes",
			Help:      "Container root filesystem inodes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerRootfsInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_rootfs_inodes_free",
			Help:      "Container root filesystem free inodes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerRootfsInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_rootfs_inodes_used",
			Help:      "Container root filesystem used inodes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerLogsAvailableBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_logs_available_bytes",
			Help:      "Container logs available in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerLogsCapacityBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_logs_capacity_bytes",
			Help:      "Container logs capacity in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerLogsUsageBytes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_logs_usage_bytes",
			Help:      "Container logs usage in bytes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerLogsInodes: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_logs_inodes",
			Help:      "Container logs inodes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerLogsInodesFree: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_logs_inodes_free",
			Help:      "Container logs free inodes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
		containerLogsInodesUsed: newGaugeVec(prometheus.GaugeOpts{
			Namespace: NS,
			Subsystem: SS,
			Name:      "container_logs_inodes_used",
			Help:      "Container logs used inodes",
		}, []string{"pod", "namespace", "node", "uid", "container"}),
	}, nil
}
