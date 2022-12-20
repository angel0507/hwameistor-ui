package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	"net/http"
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
// @Param       body body api.DrbdEnableSettingReqBody true "body"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.DrbdEnableSettingRspBody
// @Failure     500 {object}  api.RspFailBody "失败"
// @Router      /settings/highavailabilitysetting/drbd [post]
func (n *SettingController) EnableDRBDSetting(ctx *gin.Context) {
	//// 获取path中的name
	//enabledrbd := ctx.Param("enabledrbd")
	var desrb api.DrbdEnableSettingReqBody
	err := ctx.ShouldBind(&desrb)
	if err != nil {
		fmt.Errorf("Unmarshal err = %v", err)
		ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
		return
	}
	enabledrbd := desrb.Enable

	if enabledrbd == true {
		setting, err := n.m.SettingController().EnableHighAvailability()
		if err != nil {
			var failRsp api.RspFailBody
			failRsp.ErrCode = 500
			failRsp.Desc = "EnableDRBDSetting Failed" + err.Error()
			ctx.JSON(http.StatusInternalServerError, failRsp)
			return
		}
		ctx.JSON(http.StatusOK, setting)
	}
}

// DRBDSettingGet godoc
// @Summary 摘要 获取高可用设置
// @Description get DRBDSettingGet
// @Tags        Setting
// @Param       enabledrbd path string false "enabledrbd"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.DrbdEnableSetting
// @Router      /settings/highavailabilitysetting [get]
func (n *SettingController) DRBDSettingGet(ctx *gin.Context) {

	setting, err := n.m.SettingController().GetDRBDSetting()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, setting)
}
