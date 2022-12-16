package hwameistor

import (
	"context"
	utils "github.com/hwameistor/hwameistor-ui/server/util"
	"math"

	"k8s.io/client-go/kubernetes"

	hwameistorapi "github.com/hwameistor/hwameistor-ui/server/api"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LocalStoragePoolController
type LocalStoragePoolController struct {
	client.Client
	record.EventRecorder

	clientset *kubernetes.Clientset
}

// NewLocalStoragePoolController
func NewLocalStoragePoolController(client client.Client, clientset *kubernetes.Clientset, recorder record.EventRecorder) *LocalStoragePoolController {
	return &LocalStoragePoolController{
		Client:        client,
		EventRecorder: recorder,
		clientset:     clientset,
	}
}

// StoragePoolList
func (lspController *LocalStoragePoolController) StoragePoolList(page, pageSize int32) (*hwameistorapi.StoragePoolList, error) {

	var storagePoolList = &hwameistorapi.StoragePoolList{}
	sps, err := lspController.listLocalStoragePools()
	if err != nil {
		log.WithError(err).Error("Failed to listLocalStoragePool")
		return nil, err
	}

	storagePoolList.StoragePools = utils.DataPatination(sps, page, pageSize)

	var pagination = &hwameistorapi.Pagination{}
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Total = uint32(len(sps))
	pagination.Pages = int32(math.Ceil(float64(len(sps)) / float64(pageSize)))

	storagePoolList.Page = pagination

	return storagePoolList, nil
}

// listLocalStoragePools
func (lspController *LocalStoragePoolController) listLocalStoragePools() ([]*hwameistorapi.StoragePool, error) {

	storagePoolNodesCollectionMap, err := lspController.makeStoragePoolNodesCollectionMap()
	if err != nil {
		log.WithError(err).Error("Failed to makeStoragePoolNodesCollectionMap")
		return nil, err
	}
	var sps []*hwameistorapi.StoragePool
	for poolName, poolNodeCollection := range storagePoolNodesCollectionMap {
		var sp = &hwameistorapi.StoragePool{}
		sp.Name = poolName
		sp.Class = poolNodeCollection.StoragePool.Class
		sp.CreateTime = poolNodeCollection.StoragePool.CreateTime
		sp.TotalCapacityBytes = poolNodeCollection.StoragePool.TotalCapacityBytes
		sp.AllocatedCapacityBytes = poolNodeCollection.StoragePool.AllocatedCapacityBytes
		sp.NodeNum = int64(len(poolNodeCollection.ManagedNodeNames))
		sps = append(sps, sp)
	}

	return sps, nil
}

// makeStoragePoolNodesCollectionMap
func (lspController *LocalStoragePoolController) makeStoragePoolNodesCollectionMap() (map[string]*hwameistorapi.StoragePoolNodesCollection, error) {

	lsnList := &apisv1alpha1.LocalStorageNodeList{}
	if err := lspController.Client.List(context.TODO(), lsnList); err != nil {
		log.WithError(err).Error("Failed to list LocalStorageNodes")
		return nil, err
	}

	var storagePoolNodesCollectionMap = make(map[string]*hwameistorapi.StoragePoolNodesCollection)
	for _, lsn := range lsnList.Items {
		for _, pool := range lsn.Status.Pools {
			if spnc, exists := storagePoolNodesCollectionMap[pool.Name]; exists {
				spnc.ManagedNodeNames = append(spnc.ManagedNodeNames, lsn.Name)
				spnc.StoragePool.Name = pool.Name
				spnc.StoragePool.Class = pool.Class
				spnc.StoragePool.TotalCapacityBytes += pool.TotalCapacityBytes
				spnc.StoragePool.AllocatedCapacityBytes += pool.UsedCapacityBytes
				spnc.StoragePool.CreateTime = lsn.CreationTimestamp.Time
				storagePoolNodesCollectionMap[pool.Name] = spnc
			} else {
				spncnew := &hwameistorapi.StoragePoolNodesCollection{}
				spncnew.ManagedNodeNames = append(spncnew.ManagedNodeNames, lsn.Name)
				spncnew.StoragePool.Name = pool.Name
				spncnew.StoragePool.Class = pool.Class
				spncnew.StoragePool.TotalCapacityBytes += pool.TotalCapacityBytes
				spncnew.StoragePool.AllocatedCapacityBytes += pool.UsedCapacityBytes
				spncnew.StoragePool.CreateTime = lsn.CreationTimestamp.Time
				storagePoolNodesCollectionMap[pool.Name] = spncnew
			}
		}
	}

	return storagePoolNodesCollectionMap, nil
}

// GetStoragePool
func (lspController *LocalStoragePoolController) GetStoragePool(poolName string) (*hwameistorapi.StoragePool, error) {
	sps, err := lspController.listLocalStoragePools()
	if err != nil {
		log.WithError(err).Error("Failed to listLocalStoragePools")
		return nil, err
	}

	for _, sp := range sps {
		if sp.Name == poolName {
			return sp, nil
		}
	}

	return nil, nil
}

// GetStorageNodeByPoolName
func (lspController *LocalStoragePoolController) GetStorageNodeByPoolName(poolName string, page, pageSize int32) (*hwameistorapi.StorageNodeListByPool, error) {

	snlist, err := lspController.getStorageNodeByPoolName(poolName)
	if err != nil {
		log.WithError(err).Error("Failed to getStorageNodeByPoolName")
		return nil, err
	}
	var snlistByPool = &hwameistorapi.StorageNodeListByPool{}

	snlistByPool.StorageNodes = utils.DataPatination(snlist, page, pageSize)
	snlistByPool.StoragePoolName = poolName

	var pagination = &hwameistorapi.Pagination{}
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Total = uint32(len(snlist))
	pagination.Pages = int32(math.Ceil(float64(len(snlist)) / float64(pageSize)))
	snlistByPool.Page = pagination

	return snlistByPool, nil
}

// GetStorageNodeByPoolName
func (lspController *LocalStoragePoolController) getStorageNodeByPoolName(poolName string) ([]*hwameistorapi.StorageNode, error) {
	storagePoolNodesCollectionMap, err := lspController.makeStoragePoolNodesCollectionMap()
	if err != nil {
		log.WithError(err).Error("Failed to makeStoragePoolNodesCollectionMap")
		return nil, err
	}

	var sns []*hwameistorapi.StorageNode
	lsnController := NewLocalStorageNodeController(lspController.Client, lspController.clientset, lspController.EventRecorder)
	if spnc, exists := storagePoolNodesCollectionMap[poolName]; exists {
		for _, nodeName := range spnc.ManagedNodeNames {
			sn, err := lsnController.GetStorageNode(nodeName)
			if err != nil {
				log.WithError(err).Error("Failed to GetStorageNode")
				return nil, err
			}
			sns = append(sns, sn)
		}
	}

	return sns, nil
}

// StorageNodeDisksGetByPoolName
func (lspController *LocalStoragePoolController) StorageNodeDisksGetByPoolName(queryPage hwameistorapi.QueryPage) (*hwameistorapi.NodeDiskListByPool, error) {
	storagePoolNodesCollectionMap, err := lspController.makeStoragePoolNodesCollectionMap()
	if err != nil {
		log.WithError(err).Error("Failed to makeStoragePoolNodesCollectionMap")
		return nil, err
	}

	var nodeDiskListByPool = &hwameistorapi.NodeDiskListByPool{}
	var lds []*hwameistorapi.LocalDisk
	lsnController := NewLocalStorageNodeController(lspController.Client, lspController.clientset, lspController.EventRecorder)
	if spnc, exists := storagePoolNodesCollectionMap[queryPage.PoolName]; exists {
		for _, nn := range spnc.ManagedNodeNames {
			if nn == queryPage.NodeName {
				tmplds, err := lsnController.ListStorageNodeDisks(queryPage)
				if err != nil {
					log.WithError(err).Error("Failed to ListStorageNodeDisks")
					return nil, err
				}
				for _, ld := range tmplds {
					if ld.LocalStoragePooLName == queryPage.PoolName {
						lds = append(lds, ld)
					}
				}
			}
		}
	}
	nodeDiskListByPool.StoragePoolName = queryPage.PoolName
	nodeDiskListByPool.NodeName = queryPage.NodeName

	nodeDiskListByPool.LocalDisks = utils.DataPatination(lds, queryPage.Page, queryPage.PageSize)

	var pagination = &hwameistorapi.Pagination{}
	pagination.Page = queryPage.Page
	pagination.PageSize = queryPage.PageSize
	pagination.Total = uint32(len(lds))
	pagination.Pages = int32(math.Ceil(float64(len(lds)) / float64(queryPage.PageSize)))
	nodeDiskListByPool.Page = pagination

	return nodeDiskListByPool, nil
}

// listClaimedLocalDiskByNode
func (lspController *LocalStoragePoolController) listClaimedLocalDiskByNode(nodeName string) ([]apisv1alpha1.LocalDisk, error) {
	diskList := &apisv1alpha1.LocalDiskList{}
	if err := lspController.Client.List(context.TODO(), diskList); err != nil {
		log.WithError(err).Error("Failed to list LocalDisks")
		return nil, err
	}

	var claimedLocalDisks []apisv1alpha1.LocalDisk
	for i := range diskList.Items {
		if diskList.Items[i].Spec.NodeName == nodeName {
			if diskList.Items[i].Status.State == apisv1alpha1.LocalDiskBound {
				claimedLocalDisks = append(claimedLocalDisks, diskList.Items[i])
			}
		}
	}

	return claimedLocalDisks, nil
}

// LocalDiskListByNode
//func (lspController *LocalStoragePoolController) LocalDiskListByNode(nodeName string, page, pageSize int32) (*hwameistorapi.LocalDiskListByNode, error) {
//
//	var localDiskList = &hwameistorapi.LocalDiskListByNode{}
//
//	disks, err := lspController.ListStorageNodeDisks(nodeName)
//	if err != nil {
//		log.WithError(err).Error("Failed to ListStorageNodeDisks")
//		return nil, err
//	}
//
//	var pagination = &hwameistorapi.Pagination{}
//	pagination.Page = page
//	pagination.PageSize = pageSize
//	pagination.Total = uint32(len(disks))
//	pagination.Pages = int32(math.Ceil(float64(len(disks)) / float64(pageSize)))
//	localDiskList.Page = pagination
//
//	localDiskList.LocalDisksItemsList.LocalDisks = utils.DataPatination(disks, page, pageSize)
//	localDiskList.NodeName = nodeName
//
//	return localDiskList, nil
//}

// ListStorageNodeDisks
//func (lspController *LocalStoragePoolController) ListStorageNodeDisks(nodeName string) ([]*hwameistorapi.LocalDisk, error) {

//diskList := &apisv1alpha1.LocalDiskList{}
//if err := lspController.Client.List(context.TODO(), diskList); err != nil {
//	log.WithError(err).Error("Failed to list LocalDisks")
//	return nil, err
//}
//
//var disks []*hwameistorapi.LocalDisk
//for i := range diskList.Items {
//	if diskList.Items[i].Spec.NodeName == nodeName {
//		var disk = &hwameistorapi.LocalDisk{}
//		disk.DevPath = diskList.Items[i].Spec.DevicePath
//		disk.State = lspController.convertLocalDiskState(diskList.Items[i].Status.State)
//		if diskList.Items[i].Spec.DiskAttributes.Type == hwameistorapi.DiskClassNameHDD {
//			disk.LocalStoragePooLName = hwameistorapi.PoolNameForHDD
//		} else if diskList.Items[i].Spec.DiskAttributes.Type == hwameistorapi.DiskClassNameSSD {
//			disk.LocalStoragePooLName = hwameistorapi.PoolNameForSSD
//		}
//		disk.Class = diskList.Items[i].Spec.DiskAttributes.Type
//		disk.HasRAID = diskList.Items[i].Spec.HasRAID
//		disk.TotalCapacityBytes = diskList.Items[i].Spec.Capacity
//		availableCapacityBytes := lspController.getAvailableDiskCapacity(nodeName, diskList.Items[i].Spec.DevicePath, diskList.Items[i].Spec.DiskAttributes.Type)
//		disk.AvailableCapacityBytes = availableCapacityBytes
//		disks = append(disks, disk)
//	}
//}

//	return nil, nil
//}
