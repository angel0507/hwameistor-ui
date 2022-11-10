package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	log "github.com/sirupsen/logrus"

	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type IVolumeController interface {
	//RestController
	VolumeGet(ctx *gin.Context)
	VolumeList(ctx *gin.Context)
	VolumeReplicasGet(ctx *gin.Context)
	VolumeReplicaYamlGet(ctx *gin.Context)
	VolumeOperationList(ctx *gin.Context)
	VolumeOperationGet(ctx *gin.Context)
	VolumeOperationYamlGet(ctx *gin.Context)
}

// VolumeController
type VolumeController struct {
	m *manager.ServerManager
}

func NewVolumeController(m *manager.ServerManager) IVolumeController {
	fmt.Println("NewVolumeController start")

	return &VolumeController{m}
}

// VolumeGet godoc
// @Summary     摘要 获取指定数据卷基本信息
// @Description get Volume
// @Tags        Volume
// @Param       Name query string true "name"
// @Accept      application/json
// @Produce     application/json
// @Success     200 {object} api.Volume
// @Router      /volumes/volumes/:volumename [get]
func (v *VolumeController) VolumeGet(ctx *gin.Context) {
	// 获取path中的name
	volumeName := ctx.Params.ByName("name")

	if volumeName == "" {
		ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
		return
	}
	lv, err := v.m.VolumeController().GetLocalVolume(pkgclient.ObjectKey{Name: volumeName})
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}

	volume := api.ToVolumeResource(*lv)
	ctx.JSON(http.StatusOK, volume)
}

// VolumeList godoc
// @Summary 摘要 获取数据卷列表信息
// @Description list Volume
// @Tags        Volume
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.VolumeList   "成功"
// @Router      /volumes/volumes [get]
func (v *VolumeController) VolumeList(ctx *gin.Context) {

	lvs, err := v.m.VolumeController().ListLocalVolume()
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
	}
	log.Printf("VolumeList lvs = %v", lvs)

	var volums []*api.Volume
	for _, lv := range lvs.Items {
		volums = append(volums, api.ToVolumeResource(lv))
	}

	var volumeList api.VolumeList
	volumeList.Volumes = volums

	ctx.JSON(http.StatusOK, volumeList)
}

// VolumeReplicasGet godoc
// @Summary 摘要 获取指定数据卷的副本列表信息
// @Description list volumes
// @Tags        Volume
// @Param       VolumeName query string true "volumeName"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.VolumeReplicaList  "成功"
// @Router      /volumes/volumereplicas/:volumename [get]
func (v *VolumeController) VolumeReplicasGet(ctx *gin.Context) {

	//lvs, err := v.m.VolumeController().ListVolume()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//log.Printf("List lvs = %v", lvs)
	//
	//var volums []*api.Volume
	//for _, lv := range lvs.Items {
	//	volums = append(volums, api.ToVolumeResource(lv))
	//}
	//
	//var volumeList api.VolumeList
	//volumeList.Volumes = volums
	//
	//ctx.JSON(http.StatusOK, volumeList)
}

// VolumeReplicaYamlGet godoc
// @Summary 摘要 查看指定数据卷副本yaml信息
// @Description get VolumeReplicaYaml
// @Tags        Volume
// @Param       Name query string true "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.LocalVolumeReplica  "成功"
// @Router      /volumes/volumereplicas/yamls/:volumereplicaname [get]
func (v *VolumeController) VolumeReplicaYamlGet(ctx *gin.Context) {

	//lvs, err := v.m.VolumeController().ListVolume()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//log.Printf("List lvs = %v", lvs)
	//
	//var volums []*api.Volume
	//for _, lv := range lvs.Items {
	//	volums = append(volums, api.ToVolumeResource(lv))
	//}
	//
	//var volumeList api.VolumeList
	//volumeList.Volumes = volums
	//
	//ctx.JSON(http.StatusOK, volumeList)
}

// List godoc
// @Summary 摘要 获取数据卷操作记录列表信息
// @Description list VolumeOperations
// @Tags        Volume
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.VolumeOperationList      "成功"
// @Router      /volumes/volumeoperations [get]
func (v *VolumeController) VolumeOperationList(ctx *gin.Context) {

	//lvs, err := v.m.VolumeController().ListVolume()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//log.Printf("List lvs = %v", lvs)
	//
	//var volums []*api.Volume
	//for _, lv := range lvs.Items {
	//	volums = append(volums, api.ToVolumeResource(lv))
	//}
	//
	//var volumeList api.VolumeList
	//volumeList.Volumes = volums
	//
	//ctx.JSON(http.StatusOK, volumeList)
}

// VolumeOperationGet godoc
// @Summary 摘要 获取指定数据卷操作记录信息（目前仅包含迁移运维操作）
// @Description get VolumeOperation
// @Tags        Volume
// @Param       VolumeName query string true "volumeName"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.VolumeOperation      "成功"
// @Router      /volumes/volumeoperations/:volumename [get]
func (v *VolumeController) VolumeOperationGet(ctx *gin.Context) {

	//lvs, err := v.m.VolumeController().ListVolume()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//log.Printf("List lvs = %v", lvs)
	//
	//var volums []*api.Volume
	//for _, lv := range lvs.Items {
	//	volums = append(volums, api.ToVolumeResource(lv))
	//}
	//
	//var volumeList api.VolumeList
	//volumeList.Volumes = volums
	//
	//ctx.JSON(http.StatusOK, volumeList)
}

// VolumeOperationYamlGet godoc
// @Summary 摘要 获取数据卷操作记录yaml信息
// @Description get VolumeOperationYamlGet
// @Tags        Volume
// @Param       Name query string true "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.LocalVolumeMigrate  "成功"
// @Router      /volumes/volumeoperations/yamls/:operationname [get]
func (v *VolumeController) VolumeOperationYamlGet(ctx *gin.Context) {

	//lvs, err := v.m.VolumeController().ListVolume()
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//log.Printf("List lvs = %v", lvs)
	//
	//var volums []*api.Volume
	//for _, lv := range lvs.Items {
	//	volums = append(volums, api.ToVolumeResource(lv))
	//}
	//
	//var volumeList api.VolumeList
	//volumeList.Volumes = volums
	//
	//ctx.JSON(http.StatusOK, volumeList)
}
