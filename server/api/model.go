package api

type State string

const (
	// purpose of the following CRDs is for operational job
	OperationStateSubmitted           State = "Submitted"
	OperationStateMigrateAddReplica   State = "AddReplica"
	OperationStateMigrateSyncReplica  State = "SyncReplica"
	OperationStateMigratePruneReplica State = "PruneReplica"
	OperationStateInProgress          State = "InProgress"
	OperationStateCompleted           State = "Completed"
	OperationStateToBeAborted         State = "ToBeAborted"
	OperationStateAborting            State = "Cancelled"
	OperationStateAborted             State = "Aborted"
	OperationStateFailed              State = "Failed"

	VolumeStateToBeUnmount State = "ToBeMounted"
	VolumeStateEmpty       State = ""
	VolumeStateCreated     State = "Created"
	VolumeStateCreating    State = "Creating"
	VolumeStateReady       State = "Ready"
	VolumeStateNotReady    State = "NotReady"
	VolumeStateToBeDeleted State = "ToBeDeleted"
	VolumeStateDeleted     State = "Deleted"

	VolumeReplicaStateInvalid     State = "Invalid"
	VolumeReplicaStateCreating    State = "Creating"
	VolumeReplicaStateReady       State = "Ready"
	VolumeReplicaStateNotReady    State = "NotReady"
	VolumeReplicaStateToBeDeleted State = "ToBeDeleted"
	VolumeReplicaStateDeleted     State = "Deleted"

	NodeStateReady    State = "Ready"
	NodeStateMaintain State = "Maintain"
	NodeStateOffline  State = "Offline"

	// LocalDiskUnclaimed represents that the disk is not bound to any LDC,
	// and is available for claiming.
	LocalDiskUnclaimed State = "Unclaimed"
	// LocalDiskReleased represents that the disk is released from the LDC,
	LocalDiskReleased State = "Released"
	// LocalDiskClaimed represents that the disk is bound to a LDC
	LocalDiskClaimed State = "Claimed"
	// LocalDiskInUse represents that the disk is in use but not claimed by a LDC
	LocalDiskInUse State = "Inuse"
	// LocalDiskReserved represents that the disk will be used in the feature
	LocalDiskReserved State = "Reserved"

	// LocalDiskActive is the state for the disk that is connected
	LocalDiskActive State = "Active"
	// LocalDiskInactive is the state for the disk that is disconnected
	LocalDiskInactive State = "Inactive"
	// LocalDiskUnknown is the state for the disk that cannot be determined
	// at this time(whether attached or detached)
	LocalDiskUnknown State = "Unknown"

	ModuleStatusRunning  State = "Running"
	ModuleStatusNotReady State = "NotReady"

	DrbdModuleStatusEnabled  State = "Enabled"
	DrbdModuleStatusDisabled State = "Disabled"
)

//// LocalPool is storage pool struct
//type LocalPool struct {
//	// Supported pool name: HDD_POOL, SSD_POOL, NVMe_POOL 存储池
//	Name string `json:"name,omitempty"`
//}
//
//// LocalDisk is disk struct
//type LocalDisk struct {
//	// e.g. /dev/sdb 磁盘路径
//	DevPath string `json:"devPath,omitempty"`
//
//	// Supported: HDD, SSD, NVMe, RAM 磁盘类型
//	Class string `json:"type,omitempty"`
//
//	// HasRAID 是否Raid
//	HasRAID bool `json:"hasRaid,omitempty"`
//
//	// TotalCapacityBytes 总容量
//	TotalCapacityBytes int64 `json:"totalCapacityBytes,omitempty"`
//
//	// AvailableCapacityBytes 可用容量
//	AvailableCapacityBytes int64 `json:"availableCapacityBytes,omitempty"`
//
//	// Possible state: Claimed, UnClaimed, Inuse, Released, Reserved 状态
//	State State `json:"state,omitempty"`
//
//	// LocalStoragePooLName 存储池名称
//	LocalStoragePooLName string `json:"localStoragePooLName,omitempty"`
//}
//
//// LocalDiskList
//type LocalDiskList struct {
//	// nodeName 节点名称
//	NodeName string `json:"nodeName,omitempty"`
//
//	// LocalStoragePooLName 存储池名称
//	LocalStoragePooLName string `json:"localStoragePooLName,omitempty"`
//
//	// localDisks 节点磁盘列表
//	LocalDisks []*LocalDisk `json:"localDisks,omitempty"`
//}
//
//// LocalStorageNode
//type LocalStorageNode struct {
//	// name 节点名字
//	Name string `json:"name,omitempty"`
//	// ip 节点IP
//	IP string `json:"ip,omitempty"`
//	// node state 节点状态
//	NodeState string `json:"nodeState,omitempty"`
//	// driver status 驱动状态
//	DriverStatus State `json:"driverStatus,omitempty"`
//	// totalLocalDiskCount 总磁盘数
//	TotalLocalDiskCount int64 `json:"totalLocalDiskCount,omitempty"`
//	// usedLocalDiskCount 已纳管磁盘数
//	UsedLocalDiskCount int64 `json:"usedLocalDiskCount,omitempty"`
//	// freeCapacityBytes LSN可分配存储容量
//	FreeCapacityBytes int64 `json:"freeCapacityBytes,omitempty"`
//	// totalHDDCapacityBytes HDD存储总容量
//	TotalHDDCapacityBytes int64 `json:"totalHDDCapacityBytes,omitempty"`
//	// totalSSDCapacityBytes SSD存储总容量
//	TotalSSDCapacityBytes int64 `json:"totalSSDCapacityBytes,omitempty"`
//	// usedHDDCapacityBytes HDD已经使用存储量
//	UsedHDDCapacityBytes int64 `json:"usedHDDCapacityBytes,omitempty"`
//	// usedSSDCapacityBytes SSD已经使用存储量
//	UsedSSDCapacityBytes int64 `json:"usedSSDCapacityBytes,omitempty"`
//	// IsRAID 是否Raid
//	IsRAID bool `json:"isRaid,omitempty"`
//}
//
//// LocalStorageNodeList
//type LocalStorageNodeList struct {
//	// LocalStoragePooLName 存储池名称
//	LocalStoragePooLName string `json:"localStoragePooLName,omitempty"`
//	// LocalStorageNodes
//	LocalStorageNodes []*LocalStorageNode `json:"localStorageNodes,omitempty"`
//}
//
//// LocalVolume
//type LocalVolume struct {
//	// local volume name 名称
//	Name string `json:"name,omitempty"`
//
//	// local volume state 状态
//	State State `json:"state,omitempty"`
//
//	// replica number 副本数
//	ReplicaNumber int64 `json:"replicaNumber,omitempty"`
//
//	// VolumeGroup is the group name of the local volumes. It is designed for the scheduling and allocating. 磁盘组
//	VolumeGroup string `json:"volumegroup,omitempty"`
//
//	// size 容量
//	RequiredCapacityBytes int64 `json:"requiredCapacityBytes,omitempty"`
//
//	// PersistentVolumeClaimNamespace is the namespace of the associated PVC 命名空间
//	PersistentVolumeClaimNamespace string `json:"pvcNamespace,omitempty"`
//
//	// PersistentVolumeClaimName is the name of the associated PVC 绑定PVC
//	PersistentVolumeClaimName string `json:"pvcName,omitempty"`
//
//	// Convertible 转换高可用模式
//	Convertible bool `json:"convertible,omitempty"`
//
//	// createTime 创建时间
//	CreateTime time.Time `json:"createTime,omitempty"`
//}
//
//// LocalVolumeList
//type LocalVolumeList struct {
//	LocalVolumes []*LocalVolume `json:"localVolumes,omitempty"`
//}
//
//// LocalVolumeReplica
//type LocalVolumeReplica struct {
//	// replica name
//	Name string `json:"name,omitempty"`
//
//	// replica state
//	State State `json:"state,omitempty"`
//
//	// Synced is the sync state of the volume replica, which is important in HA volume 同步状态
//	Synced bool `json:"synced,omitempty"`
//
//	// NodeName is the assigned node where the volume replica is located 节点
//	NodeName string `json:"nodeName,omitempty"`
//
//	// RequiredCapacityBytes 容量
//	RequiredCapacityBytes int64 `json:"requiredCapacityBytes,omitempty"`
//
//	// StoragePath is a real path of the volume replica, like /dev/sdg.
//	StoragePath string `json:"storagePath,omitempty"`
//
//	// DevicePath is a link path of the StoragePath of the volume replica,
//	// e.g. /dev/LocalStorage_PoolHDD/pvc-fbf3ffc3-66db-4dae-9032-bda3c61b8f85
//	DevicePath string `json:"devicePath,omitempty"`
//}
//
//// LocalVolumeReplicaList
//type LocalVolumeReplicaList struct {
//	// volume name
//	VolumeName string `json:"volumeName,omitempty"`
//	// localVolumeReplicas
//	LocalVolumeReplicas []*LocalVolumeReplica `json:"localVolumeReplicas,omitempty"`
//}
//
//// LocalVolumeOperationList
//type LocalVolumeOperationList struct {
//	// node name
//	NodeName string `json:"nodeName,omitempty"`
//	// localVolumeOperations
//	LocalVolumeOperations []*LocalVolumeOperation `json:"localVolumeOperations,omitempty"`
//}
//
//// only migrate operation now
//type LocalVolumeOperation struct {
//	// volume name
//	VolumeName string `json:"volumeName,omitempty"`
//	// localVolumeMigrates
//	LocalVolumeMigrates []*LocalVolumeMigrate `json:"localVolumeMigrates,omitempty"`
//}
//
//type LocalVolumeMigrate struct {
//	// LocalVolumeMigrateName 迁移CRD名称
//	Name string `json:"name"`
//
//	// State 迁移状态
//	State State `json:"state,omitempty"`
//
//	// LocalVolumeName 迁移卷名称
//	LocalVolumeName string `json:"localVolumeName"`
//
//	// SourceNode 迁移源节点
//	SourceNode string `json:"sourceNode"`
//
//	// TargetNode 迁移目的节点
//	TargetNode string `json:"targetNode"`
//
//	// StartTime 迁移开始时间
//	StartTime time.Time `json:"startTime,omitempty"`
//
//	// EndTime 迁移结束时间
//	EndTime time.Time `json:"endTime,omitempty"`
//}
//
//type LocalStoragePool struct {
//	// Supported pool name: HDD_POOL, SSD_POOL, NVMe_POOL 存储池名称
//	Name string `json:"name,omitempty"`
//
//	// Supported class: HDD, SSD, NVMe 磁盘类型
//	Class string `json:"class"`
//
//	// HighAvailability 是否高可用
//	HighAvailability bool `json:"highAvailability"`
//
//	// TotalCapacityBytes 存储池对应存储总容量
//	TotalCapacityBytes int64 `json:"totalCapacityBytes"`
//
//	// UsedCapacityBytes 存储池已经使用存储容量
//	UsedCapacityBytes int64 `json:"usedCapacityBytes"`
//
//	// NodesNum 节点数
//	NodeNum int64 `json:"nodesNum"`
//
//	// createTime 创建时间
//	CreateTime time.Time `json:"createTime,omitempty"`
//}
//
//// LocalStoragePoolList
//type LocalStoragePoolList struct {
//	LocalStoragePools []*LocalStoragePool `json:"localStoragePools"`
//}
//
//type BaseMetric struct {
//	// 高可用卷数目
//	HighAvailableVolumeNum int64 `json:"highAvailableVolumeNum"`
//	// 非高可用卷数目
//	NonHighAvailableVolumeNum int64 `json:"nonHighAvailableVolumeNum"`
//	// 本地卷数目
//	LocalVolumeNum int64 `json:"localVolumeNum"`
//	// 总容量
//	TotalCapacityBytes int64 `json:"totalCapacityBytes"`
//	// 已使用容量
//	UsedCapacityBytes int64 `json:"usedCapacityBytes"`
//	// 已预留容量
//	ReservedCapacityBytes int64 `json:"reservedCapacityBytes"`
//	// 可使用容量
//	FreeCapacityBytes int64 `json:"freeCapacityBytes"`
//	// 总磁盘数
//	TotalDiskNum int64 `json:"totalDiskNum"`
//	// 纳管磁盘
//	ClaimedDiskNum int64 `json:"claimedDiskNum"`
//	// 健康磁盘
//	HealthyDiskNum int64 `json:"healthyDiskNum"`
//	// 错误磁盘
//	UnHealthyDiskNum int64 `json:"unHealthyDiskNum"`
//	// 总节点数
//	TotalNodeNum int64 `json:"totalNodeNum"`
//	// 纳管节点数
//	ClaimedNodeNum int64 `json:"claimedNodeNum"`
//}
//
//type StoragePool struct {
//	// 存储池名字
//	Name string `json:"name"`
//	// 总容量
//	TotalCapacityBytes int64 `json:"totalCapacityBytes"`
//	// 已使用容量
//	UsedCapacityBytes int64 `json:"usedCapacityBytes"`
//}
//
//type StoragePoolUseMetric struct {
//	StoragePoolsUse []StoragePool `json:"storagePoolsUse"`
//}
//
//type NodeStorageUse struct {
//	// 存储节点名字
//	Name string `json:"name"`
//	// 总容量
//	TotalCapacityBytes int64 `json:"totalCapacityBytes"`
//	// 已使用容量
//	UsedCapacityBytes int64 `json:"usedCapacityBytes"`
//}
//
//type NodeStorageUseMetric struct {
//	NodeStoragesUse []NodeStorageUse `json:"nodeStoragesUse"`
//}
//
//type ModuleStatus struct {
//	Name  string `json:"name"`
//	State State  `json:"state"`
//}
//
//type ModuleStatusMetric struct {
//	ModulesStatus []ModuleStatus `json:"modulesStatus"`
//}
//
//type DrbdEnableSetting struct {
//	Enabledrbd bool  `json:"enabledrbd"`
//	State      State `json:"state"`
//}
//
//func ToLocalVolumeResource(lv apisv1alpha1.LocalVolume) *LocalVolume {
//	tmplv := &LocalVolume{}
//	tmplv.Name = lv.Name
//	tmplv.RequiredCapacityBytes = lv.Spec.RequiredCapacityBytes
//	tmplv.VolumeGroup = lv.Spec.VolumeGroup
//	tmplv.State = State(lv.Status.State)
//	tmplv.PersistentVolumeClaimName = lv.Spec.PersistentVolumeClaimName
//	tmplv.PersistentVolumeClaimNamespace = lv.Spec.PersistentVolumeClaimNamespace
//	tmplv.RequiredCapacityBytes = lv.Spec.RequiredCapacityBytes
//	tmplv.ReplicaNumber = lv.Spec.ReplicaNumber
//	tmplv.CreateTime = lv.CreationTimestamp.Time
//
//	return tmplv
//}
//
//func ToLocalVolumeReplicaResource(lvr apisv1alpha1.LocalVolumeReplica) *LocalVolumeReplica {
//	r := &LocalVolumeReplica{}
//	r.Name = lvr.Name
//	return r
//}
//
//func ToLocalStorageNodeResource(lsn apisv1alpha1.LocalStorageNode) *LocalStorageNode {
//	r := &LocalStorageNode{}
//
//	r.Name = lsn.Name
//	r.DriverStatus = State(lsn.Status.State)
//
//	return r
//}

//func ToLocalVolumeOperationResource(lvm apisv1alpha1.LocalVolumeOperation) *LocalVolumeOperation {
//	r := &LocalVolumeOperation{}
//
//	r.Name = lvm.Name
//	r.LocalVolumeName = lvm.Spec.VolumeName
//	r.SourceNode = lvm.Spec.SourceNodesNames[0]
//	r.TargetNode = lvm.Spec.TargetNodesNames[0]
//	r.State = State(lvm.Status.State)
//	r.StartTime = lvm.CreationTimestamp.Time
//	// todo r.EndTime
//
//	return r
//}

//func ToLocalDiskResource(ld apisv1alpha1.LocalDisk) *LocalDisk {
//	r := &LocalDisk{}
//
//	r.DevPath = ld.Spec.DevicePath
//	r.State = State(ld.Status.State)
//	r.HasRAID = ld.Spec.HasRAID
//	r.Class = ld.Spec.DiskAttributes.Type
//	r.TotalCapacityBytes = ld.Spec.Capacity
//	// todo r.LocalStoragePooLName = ld.Spec.DiskAttributes.Type
//
//	return r
//}
