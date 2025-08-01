package types

// Summary represents the structure of the JSON response from the kubelet summary API.
type Summary struct {
	Node NodeSummary  `json:"node"`
	Pods []PodSummary `json:"pods"`
}

// NodeSummary contains information about the node in the summary.
type NodeSummary struct {
	NodeName string `json:"nodeName"`
}

// PodSummary contains information about each pod in the summary.
type PodSummary struct {
	PodRef struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
		UID       string `json:"uid"`
	} `json:"podRef"`
	Containers       []ContainerSummary `json:"containers"`
	EphemeralStorage struct {
		AvailableBytes uint64 `json:"availableBytes"`
		CapacityBytes  uint64 `json:"capacityBytes"`
		UsedBytes      uint64 `json:"usedBytes"`
		Inodes         uint64 `json:"inodes"`
		InodesFree     uint64 `json:"inodesFree"`
		InodesUsed     uint64 `json:"inodesUsed"`
	} `json:"ephemeral-storage"`
}

// ContainerSummary contains information about each container in the pod summary.
type ContainerSummary struct {
	Name   string `json:"name"`
	Memory struct {
		AvailableBytes uint64 `json:"availableBytes"`
		CapacityBytes  uint64 `json:"capacityBytes"`
		UsageBytes     uint64 `json:"usageBytes"`
	} `json:"memory"`
	Rootfs struct {
		AvailableBytes uint64 `json:"availableBytes"`
		CapacityBytes  uint64 `json:"capacityBytes"`
		UsedBytes      uint64 `json:"usedBytes"`
		Inodes         uint64 `json:"inodes"`
		InodesFree     uint64 `json:"inodesFree"`
		InodesUsed     uint64 `json:"inodesUsed"`
	} `json:"rootfs"`
	Logs struct {
		AvailableBytes uint64 `json:"availableBytes"`
		CapacityBytes  uint64 `json:"capacityBytes"`
		UsedBytes      uint64 `json:"usedBytes"`
		Inodes         uint64 `json:"inodes"`
		InodesFree     uint64 `json:"inodesFree"`
		InodesUsed     uint64 `json:"inodesUsed"`
	} `json:"logs"`
}
