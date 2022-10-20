package controller

import (
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LocalStorageNodeController
type LocalStorageNodeController struct {
	client.Client
	record.EventRecorder
	LocalStorageNode apisv1alpha1.LocalStorageNode
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
	return &apisv1alpha1.LocalStorageNodeList{}, nil
}

// GetLocalStorageNode
func (lsnController *LocalStorageNodeController) GetLocalStorageNode(key client.ObjectKey) (*apisv1alpha1.LocalStorageNode, error) {
	return &apisv1alpha1.LocalStorageNode{}, nil
}
