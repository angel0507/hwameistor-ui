package api

import (
	"time"
)

type StoragePool struct {
	// Supported pool name: HDD_POOL, SSD_POOL, NVMe_POOL 存储池名称
	Name string `json:"name,omitempty"`

	// Supported class: HDD, SSD, NVMe 磁盘类型
	Class string `json:"class"`

	// HighAvailability 是否高可用
	HighAvailability bool `json:"highAvailability"`

	// TotalCapacityBytes 存储池对应存储总容量
	TotalCapacityBytes int64 `json:"totalCapacityBytes"`

	// UsedCapacityBytes 存储池已经使用存储容量
	UsedCapacityBytes int64 `json:"usedCapacityBytes"`

	// NodesNum 节点数
	NodeNum int64 `json:"nodesNum"`

	// createTime 创建时间
	CreateTime time.Time `json:"createTime,omitempty"`
}

// StoragePoolList
type StoragePoolList struct {
	// page
	Page int32 `json:"page,omitempty"`
	// pageSize
	PageSize int32 `json:"pageSize,omitempty"`
	// storagePools
	StoragePools []*StoragePool `json:"storagePools"`
}

// NodeDiskListByPool
type NodeDiskListByPool struct {
	// page
	Page int32 `json:"page,omitempty"`
	// pageSize
	PageSize int32 `json:"pageSize,omitempty"`
	// StoragePoolName 存储池名称
	StoragePoolName string `json:"storagePoolName,omitempty"`
	// LocalDiskListByNodes
	LocalDiskListByNodes []LocalDiskListByNode `json:"localDiskListByNodes,omitempty"`
}

// StorageNodeListByPool
type StorageNodeListByPool struct {
	// StoragePoolName 存储池名称
	StoragePoolName string `json:"storagePoolName,omitempty"`
	// StorageNodes
	StorageNodes []*StorageNode `json:"storageNodes,omitempty"`
}
