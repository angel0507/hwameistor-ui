package api

import (
	"time"

	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
)

// LocalPool is storage pool struct
type LocalPool struct {
	// Supported pool name: HDD_POOL, SSD_POOL, NVMe_POOL 存储池
	Name string `json:"name,omitempty"`
}

// Volume
type Volume struct {
	// local volume name 名称
	Name string `json:"name,omitempty"`

	// local volume state 状态
	State State `json:"state,omitempty"`

	// replica number 副本数
	ReplicaNumber int64 `json:"replicaNumber,omitempty"`

	// VolumeGroup is the group name of the local volumes. It is designed for the scheduling and allocating. 磁盘组
	VolumeGroup string `json:"volumegroup,omitempty"`

	// size 容量
	RequiredCapacityBytes int64 `json:"requiredCapacityBytes,omitempty"`

	// PersistentVolumeClaimNamespace is the namespace of the associated PVC 命名空间
	PersistentVolumeClaimNamespace string `json:"pvcNamespace,omitempty"`

	// PersistentVolumeClaimName is the name of the associated PVC 绑定PVC
	PersistentVolumeClaimName string `json:"pvcName,omitempty"`

	// Convertible 转换高可用模式
	Convertible bool `json:"convertible,omitempty"`

	// createTime 创建时间
	CreateTime time.Time `json:"createTime,omitempty"`
}

// VolumeList
type VolumeList struct {
	// page
	Page int32 `json:"page,omitempty"`
	// pageSize
	PageSize int32 `json:"pageSize,omitempty"`
	// volumes
	Volumes []*Volume `json:"volumes,omitempty"`
}

// VolumeReplica
type VolumeReplica struct {
	// replica name
	Name string `json:"name,omitempty"`

	// replica state
	State State `json:"state,omitempty"`

	// Synced is the sync state of the volume replica, which is important in HA volume 同步状态
	Synced bool `json:"synced,omitempty"`

	// NodeName is the assigned node where the volume replica is located 节点
	NodeName string `json:"nodeName,omitempty"`

	// RequiredCapacityBytes 容量
	RequiredCapacityBytes int64 `json:"requiredCapacityBytes,omitempty"`

	// StoragePath is a real path of the volume replica, like /dev/sdg.
	StoragePath string `json:"storagePath,omitempty"`

	// DevicePath is a link path of the StoragePath of the volume replica,
	// e.g. /dev/LocalStorage_PoolHDD/pvc-fbf3ffc3-66db-4dae-9032-bda3c61b8f85
	DevicePath string `json:"devicePath,omitempty"`
}

// VolumeReplicaList
type VolumeReplicaList struct {
	// volume name
	VolumeName string `json:"volumeName,omitempty"`
	// VolumeReplicas
	VolumeReplicas []*VolumeReplica `json:"volumeReplicas,omitempty"`
}

// VolumeOperationList
type VolumeOperationList struct {
	// node name
	NodeName string `json:"nodeName,omitempty"`
	// VolumeOperations
	VolumeOperations []*VolumeOperation `json:"VolumeOperations,omitempty"`
}

// only migrate operation now
type VolumeOperation struct {
	// volume name
	VolumeName string `json:"volumeName,omitempty"`
	// page
	Page int32 `json:"page,omitempty"`
	// pageSize
	PageSize int32 `json:"pageSize,omitempty"`
	// VolumeMigrateOperations
	VolumeMigrateOperations []*VolumeMigrateOperation `json:"VolumeMigrateOperations,omitempty"`
}

// VolumeMigrateOperation
type VolumeMigrateOperation struct {
	// VolumeMigrateName 迁移CRD名称
	Name string `json:"name"`

	// State 迁移状态
	State State `json:"state,omitempty"`

	// VolumeName 迁移卷名称
	VolumeName string `json:"VolumeName"`

	// SourceNode 迁移源节点
	SourceNode string `json:"sourceNode"`

	// TargetNode 迁移目的节点
	TargetNode string `json:"targetNode"`

	// StartTime 迁移开始时间
	StartTime time.Time `json:"startTime,omitempty"`

	// EndTime 迁移结束时间
	EndTime time.Time `json:"endTime,omitempty"`
}

// LocalVolumeMigrateSpec defines the desired state of LocalVolumeMigrate
type LocalVolumeMigrateSpec struct {
	// volumeName
	VolumeName string `json:"volumeName"`

	// sourceNode
	SourceNode string `json:"sourceNode"`

	// targetNodesSuggested
	TargetNodesSuggested []string `json:"targetNodesSuggested"`

	// migrateAllVols
	MigrateAllVols bool `json:"migrateAllVols,omitempty"`

	// abort
	Abort bool `json:"abort,omitempty"`
}

// LocalVolumeMigrateStatus defines the observed state of LocalVolumeMigrate
type LocalVolumeMigrateStatus struct {
	// record the volume's replica number, it will be set internally
	OriginalReplicaNumber int64 `json:"originalReplicaNumber,omitempty"`
	// record the node where the specified replica is migrated to
	TargetNode string `json:"targetNode,omitempty"`

	// State of the operation, e.g. submitted, started, completed, abort, ...
	State State `json:"state,omitempty"`
	// error message to describe some states
	Message string `json:"message,omitempty"`
}

type LocalVolumeMigrate struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}

	Spec   LocalVolumeMigrateSpec   `json:"spec,omitempty"`
	Status LocalVolumeMigrateStatus `json:"status,omitempty"`
}

type LocalVolumeReplica struct {
	Kind     string `yaml:"kind"`
	Metadata struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}

	Spec   LocalVolumeReplicaSpec   `json:"spec,omitempty"`
	Status LocalVolumeReplicaStatus `json:"status,omitempty"`
}

// LocalVolumeReplicaSpec defines the desired state of LocalVolumeReplica
type LocalVolumeReplicaSpec struct {
	// VolumeName is the name of the volume, e.g. pvc-fbf3ffc3-66db-4dae-9032-bda3c61b8f85
	VolumeName string `json:"volumeName,omitempty"`

	// PoolName is the name of the storage pool, e.g. LocalStorage_PoolHDD, LocalStorage_PoolSSD, etc..
	PoolName string `json:"poolName,omitempty"`

	// NodeName is the assigned node where the volume replica is located
	NodeName string `json:"nodeName,omitempty"`

	RequiredCapacityBytes int64 `json:"requiredCapacityBytes,omitempty"`

	Delete bool `json:"delete,omitempty"`
}

// LocalVolumeReplicaStatus defines the observed state of LocalVolumeReplica
type LocalVolumeReplicaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster

	// StoragePath is a real path of the volume replica, like /dev/sdg.
	StoragePath string `json:"storagePath,omitempty"`

	// DevicePath is a link path of the StoragePath of the volume replica,
	// e.g. /dev/LocalStorage_PoolHDD/pvc-fbf3ffc3-66db-4dae-9032-bda3c61b8f85
	DevicePath string `json:"devPath,omitempty"`

	// Disks is a list of physical disks where the volume replica is spread cross, especially for striped LVM volume replica
	Disks []string `json:"disks,omitempty"`

	// AllocatedCapacityBytes is the real allocated capacity in bytes
	AllocatedCapacityBytes int64 `json:"allocatedCapacityBytes,omitempty"`

	// State is the phase of volume replica, e.g. Creating, Ready, NotReady, ToBeDeleted, Deleted
	State State `json:"state,omitempty"`

	// Synced is the sync state of the volume replica, which is important in HA volume
	Synced bool `json:"synced,omitempty"`

	// HAState is state for ha replica, replica.Status.State == Ready only when HAState is Consistent of nil
	HAState *HAState `json:"haState,omitempty"`

	// InUse is one of volume replica's states, which indicates the replica is used by a Pod or not
	InUse bool `json:"inuse,omitempty"`
}

// HAState is state for ha replica
type HAState struct {
	// Consistent, Inconsistent, replica is ready only when consistent
	State State `json:"state"`
	// Reason is why this state happened
	Reason string `json:"reason,omitempty"`
}

func ToVolumeResource(lv apisv1alpha1.LocalVolume) *Volume {
	tmplv := &Volume{}
	tmplv.Name = lv.Name
	tmplv.RequiredCapacityBytes = lv.Spec.RequiredCapacityBytes
	tmplv.VolumeGroup = lv.Spec.VolumeGroup
	tmplv.State = State(lv.Status.State)
	tmplv.PersistentVolumeClaimName = lv.Spec.PersistentVolumeClaimName
	tmplv.PersistentVolumeClaimNamespace = lv.Spec.PersistentVolumeClaimNamespace
	tmplv.RequiredCapacityBytes = lv.Spec.RequiredCapacityBytes
	tmplv.ReplicaNumber = lv.Spec.ReplicaNumber
	tmplv.CreateTime = lv.CreationTimestamp.Time

	return tmplv
}

func ToVolumeReplicaResource(lvr apisv1alpha1.LocalVolumeReplica) *VolumeReplica {
	r := &VolumeReplica{}
	r.Name = lvr.Name
	return r
}

//func ToVolumeOperationResource(lvm apisv1alpha1.VolumeOperation) *VolumeOperation {
//	r := &VolumeOperation{}
//
//	r.Name = lvm.Name
//	r.VolumeName = lvm.Spec.VolumeName
//	r.SourceNode = lvm.Spec.SourceNodesNames[0]
//	r.TargetNode = lvm.Spec.TargetNodesNames[0]
//	r.State = State(lvm.Status.State)
//	r.StartTime = lvm.CreationTimestamp.Time
//	// todo r.EndTime
//
//	return r
//}
