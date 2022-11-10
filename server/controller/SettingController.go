package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/manager"
)

type ISettingController interface {
	//RestController
	EnableDRBDSetting(ctx *gin.Context)
	DRBDSettingGet(ctx *gin.Context)
}

// SettingController
type SettingController struct {
	m *manager.ServerManager
}

func NewSettingController(m *manager.ServerManager) ISettingController {
	return &SettingController{m}
}

// EnableDRBDSetting godoc
// @Summary 摘要 高可用设置
// @Description post EnableDRBDSetting
// @Tags        Setting
// @Param       Enabledrbd query string true "enabledrbd"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.DrbdEnableSetting
// @Router      /Settings/highavailabilitysetting/:enabledrbd [post]
func (n *SettingController) EnableDRBDSetting(ctx *gin.Context) {
	// 获取path中的name
	//SettingName := ctx.Params.ByName("name")
	//
	//if SettingName == "" {
	//	ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
	//	return
	//}
	//lsn, err := n.m.LocalStorageSettingController().GetLocalStorageSetting(pkgclient.ObjectKey{Name: SettingName})
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//Setting := api.ToSettingResource(*lsn)
	//ctx.JSON(http.StatusOK, Setting)
}

// DRBDSettingGet godoc
// @Summary 摘要 获取高可用设置
// @Description get DRBDSettingGet
// @Tags        Setting
// @Param       Enabledrbd query string false "enabledrbd"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.DrbdEnableSetting
// @Router      /Settings/highavailabilitysetting [get]
func (n *SettingController) DRBDSettingGet(ctx *gin.Context) {
	// 获取path中的name
	//SettingName := ctx.Params.ByName("name")
	//
	//if SettingName == "" {
	//	ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
	//	return
	//}
	//lsn, err := n.m.LocalStorageSettingController().GetLocalStorageSetting(pkgclient.ObjectKey{Name: SettingName})
	//if err != nil {
	//	ctx.JSON(http.StatusNotFound, nil)
	//}
	//
	//Setting := api.ToSettingResource(*lsn)
	//ctx.JSON(http.StatusOK, Setting)
}
