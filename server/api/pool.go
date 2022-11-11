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
	// storagePools
	StoragePools []*StoragePool `json:"storagePools"`
	// page 信息
	Page *Pagination `json:"page,omitempty"`
}

// NodeDiskListByPool
type NodeDiskListByPool struct {
	// StoragePoolName 存储池名称
	StoragePoolName string `json:"storagePoolName,omitempty"`
	// nodeName 节点名称
	NodeName string `json:"nodeName,omitempty"`
	// localDisks 节点磁盘列表
	LocalDisks []*LocalDisk `json:"localDisks,omitempty"`
	// page 信息
	Page *Pagination `json:"page,omitempty"`
}

// StorageNodeListByPool
type StorageNodeListByPool struct {
	// StoragePoolName 存储池名称
	StoragePoolName string `json:"storagePoolName,omitempty"`
	// StorageNodes
	StorageNodes []*StorageNode `json:"storageNodes,omitempty"`
	// page 信息
	Page *Pagination `json:"page,omitempty"`
}
