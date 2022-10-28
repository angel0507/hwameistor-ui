package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	log "github.com/sirupsen/logrus"
	"net/http"
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

// Get godoc
// @Summary     摘要 获取数据卷
// @Description get volume
// @Tags        Volume
// @Param       Name query string true "name"
// @Accept      application/json
// @Produce     application/json
// @Success     200 {object} api.VolumeList
// @Router      /volumes/volumes/:name [get]
func (v *VolumeController) Get(ctx *gin.Context) {
	// 获取path中的name
	volumeName := ctx.Params.ByName("name")

	if volumeName == "" {
		ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
		return
	}
	lv, err := v.m.LocalVolumeController().GetLocalVolume(pkgclient.ObjectKey{Name: volumeName})
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	volume := api.ToVolumeResource(*lv)
	ctx.JSON(http.StatusOK, volume)
}

// List godoc
// @Summary 摘要 获取数据卷列表
// @Description list volumes
// @Tags        Volume
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.Volume      "成功"
// @Router      /volumes/volumes [get]
func (v *VolumeController) List(ctx *gin.Context) {

	lvs, err := v.m.LocalVolumeController().ListLocalVolume()
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}
	log.Printf("List lvs = %v", lvs)

	var volums []*api.Volume
	for _, lv := range lvs.Items {
		volums = append(volums, api.ToVolumeResource(lv))
	}

	var volumeList api.VolumeList
	volumeList.Volumes = volums

	ctx.JSON(http.StatusOK, volumeList)
}
