package manager

import (
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	mgrpkg "sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/hwameistor/hwameistor-ui/server/controller"
	utils "github.com/hwameistor/hwameistor-ui/server/util"
	log "github.com/sirupsen/logrus"
)

type ServerManager struct {
	nodeName string

	namespace string

	apiClient client.Client

	lsnController *controller.LocalStorageNodeController

	lvController *controller.LocalVolumeController

	mgr mgrpkg.Manager

	logger *log.Entry
}

// New replacedisk manager
func NewServerManager(mgr mgrpkg.Manager) (*ServerManager, error) {
	var recorder record.EventRecorder
	return &ServerManager{
		nodeName:      utils.GetNodeName(),
		namespace:     utils.GetNamespace(),
		apiClient:     mgr.GetClient(),
		lsnController: controller.NewLocalStorageNodeController(mgr.GetClient(), recorder),
		lvController:  controller.NewLocalVolumeController(mgr.GetClient(), recorder),
		mgr:           mgr,
		logger:        log.WithField("Module", "ServerManager"),
	}, nil
}

func (m *ServerManager) LocalStorageNodeController() *controller.LocalStorageNodeController {
	var recorder record.EventRecorder
	if m.lsnController == nil {
		m.lsnController = controller.NewLocalStorageNodeController(m.mgr.GetClient(), recorder)
	}
	return m.lsnController
}

func (m *ServerManager) LocalVolumeController() *controller.LocalVolumeController {
	var recorder record.EventRecorder
	if m.lvController == nil {
		m.lvController = controller.NewLocalVolumeController(m.mgr.GetClient(), recorder)
	}
	return m.lvController
}
