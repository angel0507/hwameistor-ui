package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	"net/http"
)

type IMetricsController interface {
	//RestController
	BaseMetric(ctx *gin.Context)
	StoragePoolUseMetric(ctx *gin.Context)
	NodeStorageUseMetric(ctx *gin.Context)
	ModuleStatusMetric(ctx *gin.Context)
	OperationList(ctx *gin.Context)
}

// MetricsController
type MetricsController struct {
	m *manager.ServerManager
}

func NewMetricsController(m *manager.ServerManager) IMetricsController {
	fmt.Println("NewMetricsController start")

	return &MetricsController{m}
}

// Get godoc
// @Summary     摘要 获取基础监控指标
// @Description get volume
// @Tags        Metric
// @Param       Name query string false "name"
// @Accept      application/json
// @Produce     application/json
// @Success     200 {object} api.BaseMetric  "成功"
// @Router      /metrics/basemetric [get]
func (v *MetricsController) BaseMetric(ctx *gin.Context) {

	baseCapacity, err := v.m.MetricController().GetBaseMetric()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, baseCapacity)
}

// StoragePoolMetric godoc
// @Summary 摘要 获取存储池资源监控指标
// @Description StoragePoolMetric
// @Tags        Metric
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.StoragePoolUseMetric  "成功"
// @Router      /metrics/storagepoolusemetric [get]
func (v *MetricsController) StoragePoolUseMetric(ctx *gin.Context) {

	storagePoolUseMetric, err := v.m.MetricController().GetStoragePoolUseMetric()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, storagePoolUseMetric)
}

// NodeStorageUseMetric godoc
// @Summary 摘要 获取指定存储池类型节点存储TOP5使用率监控指标
// @Description NodeStorageUseMetric
// @Tags        Metric
// @Param       StoragePoolClass query string true "storagePoolClass"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.NodeStorageUseMetric  "成功"
// @Router      /metrics/nodestorageusemetric/:StoragePoolClass [get]
func (v *MetricsController) NodeStorageUseMetric(ctx *gin.Context) {
	// 获取path中的storagePoolClass
	storagePoolClass := ctx.Query("StoragePoolClass")
	fmt.Println("NodeStorageUseMetric storagePoolClass = %", storagePoolClass)

	if storagePoolClass == "" {
		ctx.JSON(http.StatusNonAuthoritativeInfo, nil)
		return
	}
	
	nodeStorageUseMetric, err := v.m.MetricController().GetNodeStorageUseMetric(storagePoolClass)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, nodeStorageUseMetric)

}

// ModuleStatusMetric godoc
// @Summary 摘要 获取组件状态监控指标
// @Description ModuleStatusMetric
// @Tags        Metric
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.ModuleStatusMetric  "成功"
// @Router      /metrics/modulestatusmetric [get]
func (v *MetricsController) ModuleStatusMetric(ctx *gin.Context) {

	moduleStatusMetric, err := v.m.MetricController().GetModuleStatusMetric()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, moduleStatusMetric)
}

// OperationList godoc
// @Summary 摘要 获取操作记录列表
// @Description OperationList
// @Tags        Metric
// @Param       Name query string false "name"
// @Accept      json
// @Produce     json
// @Success     200 {object}  api.OperationMetric  "成功"
// @Router      /metrics/operations [get]
func (v *MetricsController) OperationList(ctx *gin.Context) {
	operationListMetric, err := v.m.MetricController().OperationListMetric()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, operationListMetric)
}
