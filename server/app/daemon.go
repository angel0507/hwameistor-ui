package app

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // for runtime profiling
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log"
	pkgmgr "sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/hwameistor/hwameistor-ui/server/api"
	"github.com/hwameistor/hwameistor-ui/server/manager"
	"github.com/hwameistor/hwameistor-ui/server/types"
	utils "github.com/hwameistor/hwameistor-ui/server/util"
)

const (
	EnvPodIP = "POD_IP"
)

func DaemonCmd() cli.Command {
	return cli.Command{
		Name: "daemon",
		Action: func(c *cli.Context) {
			if err := startManager(c); err != nil {
				logrus.Fatalf("Error starting manager: %v", err)
			}
		},
	}
}

func startManager(c *cli.Context) error {
	var (
		err error
	)

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Log.Error(err, "")
		os.Exit(1)
	}

	// Set default manager options
	options := pkgmgr.Options{}

	// Create a new manager to provide shared dependencies and start components
	mgr, err := pkgmgr.New(cfg, options)
	if err != nil {
		log.Log.Error(err, "")
		os.Exit(1)
	}

	// Create a new manager to provide shared dependencies and start components
	smgr, err := manager.NewServerManager(mgr)
	if err != nil {
		log.Log.Error(err, "")
		os.Exit(1)
	}

	currentIP, err := utils.GetRequiredEnv(EnvPodIP)
	if err != nil {
		return fmt.Errorf("BUG: failed to detect the node IP")
	}

	done := make(chan struct{})

	server := api.NewServer(smgr)
	router := http.Handler(api.NewRouter(server))

	listen := types.GetAPIServerAddressFromIP(currentIP)
	logrus.Infof("Listening on %s", listen)

	go http.ListenAndServe(listen, router)

	go func() {
		debugAddress := "127.0.0.1:6060"
		debugHandler := http.DefaultServeMux
		logrus.Infof("Debug Server listening on %s", debugAddress)
		if err := http.ListenAndServe(debugAddress, debugHandler); err != nil && err != http.ErrServerClosed {
			logrus.Errorf(fmt.Sprintf("ListenAndServe: %s", err))
		}
	}()

	utils.RegisterShutdownChannel(done)
	<-done
	return nil
}
