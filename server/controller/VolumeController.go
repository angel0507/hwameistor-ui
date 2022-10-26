package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	"github.com/hwameistor/hwameistor-ui/server/response"
	log "github.com/sirupsen/logrus"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type IVolumeController interface {
	RestController
}

// VolumeController
type VolumeController struct {
	m *manager.ServerManager
}

func NewVolumeController(m *manager.ServerManager) IVolumeController {
	fmt.Println("NewVolumeController start")

	return &VolumeController{m}
}

// GetVolume
func (v *VolumeController) Get(ctx *gin.Context) {
	// 获取path中的name
	volumeName := ctx.Params.ByName("name")

	if volumeName == "" {
		response.Fail(ctx, "数据验证错误，数据卷名称必填", nil)
		return
	}
	lv, err := v.m.LocalVolumeController().GetLocalVolume(pkgclient.ObjectKey{Name: volumeName})
	if err != nil {
		response.Fail(ctx, "数据卷不存在", nil)
	}

	response.Success(ctx, gin.H{"volume": api.ToVolumeResource(*lv)}, "成功")
}

// ListVolumes
func (v *VolumeController) List(ctx *gin.Context) {

	lvs, err := v.m.LocalVolumeController().ListLocalVolume()
	if err != nil {
		response.Fail(ctx, "数据卷列表不存在", nil)
	}
	log.Printf("List lvs = %v", lvs)

	var volums []*api.Volume
	for _, lv := range lvs.Items {
		volums = append(volums, api.ToVolumeResource(lv))
	}

	response.Success(ctx, gin.H{"volumes": volums}, "成功")
}
