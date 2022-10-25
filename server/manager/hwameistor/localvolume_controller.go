package hwameistor

import (
	"context"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"

	log "github.com/sirupsen/logrus"
)

// LocalVolumeController
type LocalVolumeController struct {
	client.Client
	record.EventRecorder
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
	volList := &apisv1alpha1.LocalVolumeList{}
	if err := lvController.Client.List(context.TODO(), volList); err != nil {
		log.WithError(err).Fatal("Failed to list LocalVolumes")
	}
	return volList, nil
}

// GetLocalVolume
func (lvController *LocalVolumeController) GetLocalVolume(key client.ObjectKey) (*apisv1alpha1.LocalVolume, error) {
	vol := &apisv1alpha1.LocalVolume{}
	if err := lvController.Client.Get(context.TODO(), key, vol); err != nil {
		if !errors.IsNotFound(err) {
			log.WithError(err).Error("Failed to query volume")
		} else {
			log.Info("Not found the volume")
		}
		return nil, err
	}
	return vol, nil
}
