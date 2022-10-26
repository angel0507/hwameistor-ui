package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	"net/http"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type INodeController interface {
	RestController
}

// NodeController
type NodeController struct {
	m *manager.ServerManager
}

func NewNodeController(m *manager.ServerManager) INodeController {
	return &NodeController{m}
}

// Get godoc
// @Summary 摘要 获取节点
// @Description get node
// @Tags        Node
// @Param       Name query string true "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.Node
// @Router      /nodes/nodes/:name [get]
func (n *NodeController) Get(ctx *gin.Context) {
	// 获取path中的name
	nodeName := ctx.Params.ByName("name")

	if nodeName == "" {
		ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
		return
	}
	lsn, err := n.m.LocalStorageNodeController().GetLocalStorageNode(pkgclient.ObjectKey{Name: nodeName})
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	node := api.ToNodeResource(*lsn)
	ctx.JSON(http.StatusOK, node)
}

// List godoc
// @Summary     摘要 获取节点列表
// @Description list nodes
// @Tags        Node
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object} api.NodeList
// @Router      /nodes/nodes [get]
func (n *NodeController) List(ctx *gin.Context) {

	lsns, err := n.m.LocalStorageNodeController().ListLocalStorageNode()
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	var nodes []*api.Node
	for _, lsn := range lsns.Items {
		nodes = append(nodes, api.ToNodeResource(lsn))
	}

	var nodeList api.NodeList
	nodeList.Nodes = nodes
	ctx.JSON(http.StatusOK, nodes)
}
