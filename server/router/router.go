package api

import (
	"fmt"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/controller"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	mgrpkg "sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var (
	RetryCounts         = 5
	RetryInterval       = 100 * time.Millisecond
	metricsHost         = "0.0.0.0"
	metricsPort   int32 = 8384
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	fmt.Println("CollectRoute start ...")

	m := BuildServerMgr()

	v1 := r.Group("/api/v1")
	volumeController := controller.NewVolumeController(m)
	volumeRoutes := v1.Group("/volumes")
	volumeRoutes.GET("/volumes", volumeController.List)
	volumeRoutes.GET("/volumes/:name", volumeController.Get)

	nodeController := controller.NewNodeController(m)
	nodeRoutes := v1.Group("/nodes")
	nodeRoutes.GET("/nodes", nodeController.List)
	nodeRoutes.GET("/nodes/:name", nodeController.Get)

	fmt.Println("CollectRoute end ...")

	return r
}

func BuildServerMgr() *manager.ServerManager {
	fmt.Println("buildServerMgr start ...")

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Set default manager options
	options := mgrpkg.Options{MetricsBindAddress: fmt.Sprintf("%s:%d", metricsHost, metricsPort)}

	// Create a new manager to provide shared dependencies and start components
	mgr, err := mgrpkg.New(cfg, options)
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup Scheme for all resources
	if err := api.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	if err := apisv1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "Failed to setup scheme for ldm resources")
		os.Exit(1)
	}

	// Setup all Controllers
	if err := controller.AddToManager(mgr); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	stopCh := signals.SetupSignalHandler()
	// Start the resource controllers manager
	go func() {
		log.Info("Starting the manager of all local storage resources.")
		if err := mgr.Start(stopCh); err != nil {
			log.WithError(err).Error("Failed to run resources manager")
			os.Exit(1)
		}
	}()

	// Create a new manager to provide shared dependencies and start components
	smgr, err := manager.NewServerManager(mgr)
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}
	return smgr
}
