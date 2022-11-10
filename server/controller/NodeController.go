package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/manager"
)

type INodeController interface {
	//RestController
	StorageNodeGet(ctx *gin.Context)
	StorageNodeList(ctx *gin.Context)
	StorageNodeMigrateGet(ctx *gin.Context)

	StorageNodeDisksList(ctx *gin.Context)
}

// NodeController
type NodeController struct {
	m *manager.ServerManager
}

func NewNodeController(m *manager.ServerManager) INodeController {
	return &NodeController{m}
}

// StorageNodeGet godoc
// @Summary 摘要 获取指定存储节点
// @Description get StorageNode
// @Tags        Node
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.StorageNode
// @Router      /nodes/storagenodes/:storagenodename [get]
func (n *NodeController) StorageNodeGet(ctx *gin.Context) {
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

// StorageNodeList godoc
// @Summary     摘要 获取存储节点列表
// @Description list StorageNode
// @Tags        Node
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object} api.StorageNodeList
// @Router      /nodes/storagenodes [get]
func (n *NodeController) StorageNodeList(ctx *gin.Context) {

	//lsns, err := n.m.StorageNodeController().ListStorageNode()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//var nodes []*api.StorageNode
	//for _, lsn := range lsns.Items {
	//	nodes = append(nodes, api.ToNodeResource(lsn))
	//}
	//
	//var nodeList api.NodeList
	//nodeList.Nodes = nodes
	//ctx.JSON(http.StatusOK, nodes)
}

// StorageNodeMigrateGet godoc
// @Summary     摘要 获取指定节点数据卷迁移任务列表
// @Description get StorageNodeMigrate
// @Tags        Node
// @Param       NodeName query string true "nodeName"
// @Accept      json
// @Produce     json
// @Success     200 {object} api.VolumeOperationList
// @Router      /nodes/storagenodes/:storagenodename/migrates [get]
func (n *NodeController) StorageNodeMigrateGet(ctx *gin.Context) {

}

// StorageNodeDisksList godoc
// @Summary 摘要 获取指定存储节点磁盘列表
// @Description list StorageNodeDisks
// @Tags        Node
// @Param       NodeName query string true "nodeName"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.LocalDiskListByNode
// @Router      /nodes/storagenodes/:storagenodename/disks [get]
func (n *NodeController) StorageNodeDisksList(ctx *gin.Context) {
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
