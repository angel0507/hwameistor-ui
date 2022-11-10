package api

import (
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
)

// StorageNode
type StorageNode struct {
	// name 节点名字
	Name string `json:"name,omitempty"`
	// ip 节点IP
	IP string `json:"ip,omitempty"`
	// node state 节点状态
	NodeState string `json:"nodeState,omitempty"`
	// driver status 驱动状态
	DriverStatus State `json:"driverStatus,omitempty"`
	// totalDiskCount 总磁盘数
	TotalDiskCount int64 `json:"totalDiskCount,omitempty"`
	// usedDiskCount 已纳管磁盘数
	UsedDiskCount int64 `json:"usedDiskCount,omitempty"`
	// freeCapacityBytes LSN可分配存储容量
	FreeCapacityBytes int64 `json:"freeCapacityBytes,omitempty"`
	// totalHDDCapacityBytes HDD存储总容量
	TotalHDDCapacityBytes int64 `json:"totalHDDCapacityBytes,omitempty"`
	// totalSSDCapacityBytes SSD存储总容量
	TotalSSDCapacityBytes int64 `json:"totalSSDCapacityBytes,omitempty"`
	// usedHDDCapacityBytes HDD已经使用存储量
	UsedHDDCapacityBytes int64 `json:"usedHDDCapacityBytes,omitempty"`
	// usedSSDCapacityBytes SSD已经使用存储量
	UsedSSDCapacityBytes int64 `json:"usedSSDCapacityBytes,omitempty"`
	// IsRAID 是否Raid
	IsRAID bool `json:"isRaid,omitempty"`
}

// LocalDiskListByNode
type LocalDiskListByNode struct {
	// nodeName 节点名称
	NodeName string `json:"nodeName,omitempty"`
	// page
	Page int32 `json:"page,omitempty"`
	// pageSize
	PageSize int32 `json:"pageSize,omitempty"`
	// localDisks 节点磁盘列表
	LocalDisks []*LocalDisk `json:"localDisks,omitempty"`
}

// StorageNodeList
type StorageNodeList struct {
	// page
	Page int32 `json:"page,omitempty"`
	// pageSize
	PageSize int32 `json:"pageSize,omitempty"`
	// StorageNodes
	StorageNodes []*StorageNode `json:"storageNodes,omitempty"`
}

func ToStorageNodeResource(lsn apisv1alpha1.LocalStorageNode) *StorageNode {
	r := &StorageNode{}

	r.Name = lsn.Name
	r.DriverStatus = State(lsn.Status.State)

	return r
}
