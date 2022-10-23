package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	"github.com/hwameistor/hwameistor-ui/server/response"
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

// GetNode
func (n *NodeController) Get(ctx *gin.Context) {
	// 获取path中的name
	nodeName := ctx.Params.ByName("name")

	if nodeName == "" {
		response.Fail(ctx, "数据验证错误，数据卷名称必填", nil)
		return
	}
	lsn, err := n.m.LocalStorageNodeController().GetLocalStorageNode(pkgclient.ObjectKey{Name: nodeName})
	if err != nil {
		response.Fail(ctx, "数据节点不存在", nil)
	}

	response.Success(ctx, gin.H{"node": api.ToNodeResource(*lsn)}, "成功")
}

// ListNodes
func (n *NodeController) List(ctx *gin.Context) {

	lsns, err := n.m.LocalStorageNodeController().ListLocalStorageNode()
	if err != nil {
		response.Fail(ctx, "数据节点列表不存在", nil)
	}

	var nodes []*api.Node
	for _, lsn := range lsns.Items {
		nodes = append(nodes, api.ToNodeResource(lsn))
	}

	response.Success(ctx, gin.H{"node": nodes}, "成功")
}
