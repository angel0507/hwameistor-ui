package hwameistor

import (
	"context"
	"fmt"
	hwameistorapi "github.com/hwameistor/hwameistor-ui/server/api"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

// VolumeGroupController
type VolumeGroupController struct {
	client.Client
	record.EventRecorder

	clientset *kubernetes.Clientset
}

// NewVolumeGroupController
func NewVolumeGroupController(client client.Client, clientset *kubernetes.Clientset, recorder record.EventRecorder) *VolumeGroupController {
	return &VolumeGroupController{
		Client:        client,
		EventRecorder: recorder,
		clientset:     clientset,
	}
}

// ListVolumesByVolumeGroup
func (vgController *VolumeGroupController) ListVolumesByVolumeGroup(vgName string) (hwameistorapi.VolumeGroup, error) {
	var vgvis = []hwameistorapi.VolumeGroupVolumeInfo{}

	var vg = hwameistorapi.VolumeGroup{}
	lvg := &apisv1alpha1.LocalVolumeGroup{}
	if err := vgController.Client.Get(context.TODO(), client.ObjectKey{Name: vgName}, lvg); err != nil {
		if !errors.IsNotFound(err) {
			log.WithError(err).Error("Failed to query LocalVolumeGroup")
		} else {
			log.Info("Not found the LocalVolumeGroup")
		}
		return vg, err
	}

	fmt.Println("ListVolumesByVolumeGroup lvg.Name = %v, lvg.Spec.Volumes = %v", lvg.Name, lvg.Spec.Volumes)
	vg.Name = lvg.Name

	for _, volumeinfo := range lvg.Spec.Volumes {
		var vgvi = hwameistorapi.VolumeGroupVolumeInfo{}

		vgvi.VolumeName = volumeinfo.LocalVolumeName
		if vgvi.VolumeName == "" {
			vgvi.VolumeName = volumeinfo.LocalVolumeName
		}
		fmt.Println("ListVolumesByVolumeGroup vgvi.VolumeName = %v", vgvi.VolumeName)
		lv := &apisv1alpha1.LocalVolume{}
		if err := vgController.Client.Get(context.TODO(), client.ObjectKey{Name: vgvi.VolumeName}, lv); err != nil {
			if !errors.IsNotFound(err) {
				log.WithError(err).Error("Failed to query localvolume")
			} else {
				log.Info("Not found the localvolume")
			}
			return vg, err
		}
		vgvi.State = hwameistorapi.StateConvert(lv.Status.State)

		for _, replicas := range lv.Spec.Config.Replicas {
			vgvi.NodeNames = append(vgvi.NodeNames, replicas.Hostname)
		}
		vgvis = append(vgvis, vgvi)
	}

	fmt.Println("ListVolumesByVolumeGroup len(vgvis) = %v, vgvis[0].NodeNames = %v", len(vgvis), vgvis[0].NodeNames)
	if len(vgvis) != 0 {
		vg.NodeNames = vgvis[0].NodeNames
	}
	vg.VolumeGroupVolumeInfos = vgvis
	vg.Name = vgName

	return vg, nil
}

func (vgController *VolumeGroupController) ListVolumeGroup() (*hwameistorapi.VolumeGroupList, error) {

	var vglist = &hwameistorapi.VolumeGroupList{}
	lvList := &apisv1alpha1.LocalVolumeList{}
	if err := vgController.Client.List(context.TODO(), lvList); err != nil {
		log.WithError(err).Error("Failed to list LocalVolumes")
		return nil, err
	}

	var vgnames []string
	for _, lv := range lvList.Items {
		var vgsnamestr string = strings.Join(vgnames, " ")
		if lv.Spec.VolumeGroup != "" && !strings.Contains(vgsnamestr, lv.Spec.VolumeGroup) {
			vgnames = append(vgnames, lv.Spec.VolumeGroup)
		}
	}
	vglist.VolumeGroupNames = vgnames

	return vglist, nil
}
