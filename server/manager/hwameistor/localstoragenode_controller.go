package hwameistor

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"

	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LocalStorageNodeController
type LocalStorageNodeController struct {
	client.Client
	record.EventRecorder
}

// NewLocalStorageNodeController
func NewLocalStorageNodeController(client client.Client, recorder record.EventRecorder) *LocalStorageNodeController {
	return &LocalStorageNodeController{
		Client:        client,
		EventRecorder: recorder,
	}
}

// ListLocalStorageNode
func (lsnController *LocalStorageNodeController) ListLocalStorageNode() (*apisv1alpha1.LocalStorageNodeList, error) {
	lsnList := &apisv1alpha1.LocalStorageNodeList{}
	if err := lsnController.Client.List(context.TODO(), lsnList); err != nil {
		log.WithError(err).Fatal("Failed to list LocalStorageNodes")
	}
	return lsnList, nil
}

// GetLocalStorageNode
func (lsnController *LocalStorageNodeController) GetLocalStorageNode(key client.ObjectKey) (*apisv1alpha1.LocalStorageNode, error) {
	lsn := &apisv1alpha1.LocalStorageNode{}
	if err := lsnController.Client.Get(context.TODO(), key, lsn); err != nil {
		if !errors.IsNotFound(err) {
			log.WithError(err).Error("Failed to query lsn")
		} else {
			log.Printf("GetLocalStorageNode: not found lsn")
			log.WithError(err)
		}
		return nil, err
	}
	return lsn, nil
}
