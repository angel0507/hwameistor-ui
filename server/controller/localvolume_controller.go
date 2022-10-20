package controller

import (
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// LocalVolumeController
type LocalVolumeController struct {
	client.Client
	record.EventRecorder
	LocalVolume apisv1alpha1.LocalVolume
}

// NewLocalVolumeController
func NewLocalVolumeController(client client.Client, recorder record.EventRecorder) *LocalVolumeController {
	return &LocalVolumeController{
		Client:        client,
		EventRecorder: recorder,
	}
}

// ListLocalVolume
func (lvController *LocalVolumeController) ListLocalVolume() (*apisv1alpha1.LocalVolumeList, error) {
	return &apisv1alpha1.LocalVolumeList{}, nil
}

// GetLocalVolume
func (lvController *LocalVolumeController) GetLocalVolume(key client.ObjectKey) (*apisv1alpha1.LocalVolume, error) {
	return &apisv1alpha1.LocalVolume{}, nil
}
