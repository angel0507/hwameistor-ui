package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/manager"
)

type IPoolController interface {
	//RestController
	StoragePoolGet(ctx *gin.Context)
	StoragePoolList(ctx *gin.Context)
	StorageNodesGetByPoolName(ctx *gin.Context)
	StorageNodeDisksGetByPoolName(ctx *gin.Context)
}

// PoolController
type PoolController struct {
	m *manager.ServerManager
}

func NewPoolController(m *manager.ServerManager) IPoolController {
	return &PoolController{m}
}

// StoragePoolGet godoc
// @Summary 摘要 获取指定存储池基本信息
// @Description get Pool
// @Tags        Pool
// @Param       Name query string true "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.StoragePool
// @Router      /pools/storagepools/:storagepoolname [get]
func (n *PoolController) StoragePoolGet(ctx *gin.Context) {
	// 获取path中的name
	//PoolName := ctx.Params.ByName("name")
	//
	//if PoolName == "" {
	//	ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
	//	return
	//}
	//lsn, err := n.m.StoragePoolController().GetStoragePool(pkgclient.ObjectKey{Name: PoolName})
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//Pool := api.ToPoolResource(*lsn)
	//ctx.JSON(http.StatusOK, Pool)
}

// StoragePoolList godoc
// @Summary     摘要 获取存储池列表信息
// @Description list StoragePools
// @Tags        Pool
// @Param       Name query string false "name"
// @Param       Page query int32 true "page"
// @Param       PageSize query int32 true "pageSize"
// @Accept      json
// @Produce     json
// @Success     200 {object} api.StoragePoolList
// @Router      /pools/storagepools [get]
func (n *PoolController) StoragePoolList(ctx *gin.Context) {

	//lsns, err := n.m.StoragePoolController().ListStoragePool()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//var Pools []*api.StoragePool
	//for _, lsn := range lsns.Items {
	//	Pools = append(Pools, api.ToPoolResource(lsn))
	//}
	//
	//var PoolList api.PoolList
	//PoolList.Pools = Pools
	//ctx.JSON(http.StatusOK, Pools)
}

// StorageNodesGetByPoolName godoc
// @Summary 摘要 获取指定存储池存储节点列表信息
// @Description get StorageNodesGetByPoolName
// @Tags        Pool
// @Param       StoragePoolName query string true "storagePoolName"
// @Param       Page query int32 true "page"
// @Param       PageSize query int32 true "pageSize"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.StorageNodeListByPool
// @Router      /pools/storagepools/:storagepoolname/nodes [get]
func (n *PoolController) StorageNodesGetByPoolName(ctx *gin.Context) {
	// 获取path中的name
	//nodeName := ctx.Params.ByName("name")
	//
	//if nodeName == "" {
	//	ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
	//	return
	//}
	//lsn, err := n.m.StorageNodeController().GetStorageNode(pkgclient.ObjectKey{Name: nodeName})
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//node := api.ToNodeResource(*lsn)
	//ctx.JSON(http.StatusOK, node)
}

// StorageNodeDisksGetByPoolName godoc
// @Summary 摘要 获取指定存储池指定存储节点磁盘列表信息
// @Description get StorageNodeDisksGetByPoolName
// @Tags        Pool
// @Param       NodeName query string true "nodeName"
// @Param       StoragePoolName query string true "storagePoolName"
// @Param       Page query int32 true "page"
// @Param       PageSize query int32 true "pageSize"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.NodeDiskListByPool
// @Router      /pools/storagepools/:storagepoolname/nodes/:nodename/disks [get]
func (n *PoolController) StorageNodeDisksGetByPoolName(ctx *gin.Context) {
	// 获取path中的name
	//nodeName := ctx.Params.ByName("name")
	//
	//if nodeName == "" {
	//	ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
	//	return
	//}
	//lsn, err := n.m.StorageDiskController().GetStorageNode(pkgclient.ObjectKey{Name: nodeName})
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//node := api.ToNodeResource(*lsn)
	//ctx.JSON(http.StatusOK, node)
}
