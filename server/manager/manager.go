package manager

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	mgrpkg "sigs.k8s.io/controller-runtime/pkg/manager"

	hwameistorctr "github.com/hwameistor/hwameistor-ui/server/manager/hwameistor"
	utils "github.com/hwameistor/hwameistor-ui/server/util"
	log "github.com/sirupsen/logrus"
)

type ServerManager struct {
	nodeName string

	namespace string

	apiClient client.Client

	clientset *kubernetes.Clientset

	lsnController *hwameistorctr.LocalStorageNodeController

	lvController *hwameistorctr.LocalVolumeController

	mController *hwameistorctr.MetricController

	mgr mgrpkg.Manager

	logger *log.Entry
}

// NewServerManager
func NewServerManager(mgr mgrpkg.Manager, clientset *kubernetes.Clientset) (*ServerManager, error) {
	var recorder record.EventRecorder
	return &ServerManager{
		nodeName:      utils.GetNodeName(),
		namespace:     utils.GetNamespace(),
		apiClient:     mgr.GetClient(),
		clientset:     clientset,
		lsnController: hwameistorctr.NewLocalStorageNodeController(mgr.GetClient(), recorder),
		lvController:  hwameistorctr.NewLocalVolumeController(mgr.GetClient(), recorder),
		mController:   hwameistorctr.NewMetricController(mgr.GetClient(), clientset, recorder),
		mgr:           mgr,
		logger:        log.WithField("Module", "ServerManager"),
	}, nil
}

func (m *ServerManager) StorageNodeController() *hwameistorctr.LocalStorageNodeController {
	var recorder record.EventRecorder
	if m.lsnController == nil {
		m.lsnController = hwameistorctr.NewLocalStorageNodeController(m.mgr.GetClient(), recorder)
	}
	return m.lsnController
}

func (m *ServerManager) VolumeController() *hwameistorctr.LocalVolumeController {
	var recorder record.EventRecorder
	if m.lvController == nil {
		m.lvController = hwameistorctr.NewLocalVolumeController(m.mgr.GetClient(), recorder)
	}
	return m.lvController
}

func (m *ServerManager) MetricController() *hwameistorctr.MetricController {

	var recorder record.EventRecorder
	if m.mController == nil {
		m.mController = hwameistorctr.NewMetricController(m.mgr.GetClient(), m.clientset, recorder)
	}
	return m.mController
}
