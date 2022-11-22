package hwameistor

import (
	"bytes"
	"context"
	"fmt"
	hwameistorapi "github.com/hwameistor/hwameistor-ui/server/api"
	utils "github.com/hwameistor/hwameistor-ui/server/util"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"math"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	groupName   = "hwameistor.io"
	versionName = "v1"
)

// LocalVolumeController
type LocalVolumeController struct {
	client.Client
	record.EventRecorder

	clientset *kubernetes.Clientset
}

// NewLocalVolumeController
func NewLocalVolumeController(client client.Client, clientset *kubernetes.Clientset, recorder record.EventRecorder) *LocalVolumeController {
	return &LocalVolumeController{
		Client:        client,
		EventRecorder: recorder,
		clientset:     clientset,
	}
}

// ListLocalVolume
func (lvController *LocalVolumeController) ListLocalVolume(page, pageSize int32) (*hwameistorapi.VolumeList, error) {
	var volList = &hwameistorapi.VolumeList{}
	vols, err := lvController.listLocalVolume()
	fmt.Println("ListLocalVolume vols = %v", vols)
	if err != nil {
		log.WithError(err).Fatal("Failed to listLocalVolume")
		return nil, err
	}
	volList.Volumes = utils.DataPatination(vols, page, pageSize)

	var pagination = &hwameistorapi.Pagination{}
	pagination.Page = page
	pagination.PageSize = pageSize
	pagination.Total = uint32(len(vols))
	if len(vols) == 0 {
		pagination.Pages = 0
	} else {
		pagination.Pages = int32(math.Ceil(float64(len(vols)) / float64(pageSize)))
	}
	volList.Page = pagination

	return volList, nil
}

// listLocalVolume
func (lvController *LocalVolumeController) listLocalVolume() ([]*hwameistorapi.Volume, error) {
	lvList := &apisv1alpha1.LocalVolumeList{}
	if err := lvController.Client.List(context.TODO(), lvList); err != nil {
		log.WithError(err).Fatal("Failed to list LocalVolumes")
		return nil, err
	}
	fmt.Println("listLocalVolume lvList = %v", lvList)

	var vols []*hwameistorapi.Volume
	for _, lv := range lvList.Items {
		var vol = &hwameistorapi.Volume{}
		vol.Name = lv.Name
		vol.ReplicaNumber = lv.Spec.ReplicaNumber
		vol.Convertible = lv.Spec.Convertible
		vol.RequiredCapacityBytes = lv.Spec.RequiredCapacityBytes
		vol.PersistentVolumeClaimNamespace = lv.Spec.PersistentVolumeClaimNamespace
		vol.PersistentVolumeClaimName = lv.Spec.PersistentVolumeClaimName
		vol.State = hwameistorapi.StateConvert(lv.Status.State)
		vol.VolumeGroup = lv.Spec.VolumeGroup
		vol.CreateTime = lv.CreationTimestamp.Time
		vols = append(vols, vol)
	}

	return vols, nil
}

// GetLocalVolume
func (lvController *LocalVolumeController) GetLocalVolume(lvname string) (*hwameistorapi.Volume, error) {
	lvs, err := lvController.listLocalVolume()
	if err != nil {
		log.WithError(err).Fatal("Failed to listLocalVolume")
		return nil, err
	}

	for _, lv := range lvs {
		if lv.Name == lvname {
			return lv, nil
		}
	}

	return nil, nil
}

// GetLocalVolumeReplicas
func (lvController *LocalVolumeController) getLocalVolumeReplicas(lvname string) ([]*apisv1alpha1.LocalVolumeReplica, error) {
	lv := &apisv1alpha1.LocalVolume{}
	if err := lvController.Client.Get(context.TODO(), client.ObjectKey{Name: lvname}, lv); err != nil {
		if !errors.IsNotFound(err) {
			log.WithError(err).Error("Failed to query diskume")
		} else {
			log.Info("Not found the diskume")
		}
		return nil, err
	}

	var lvrs []*apisv1alpha1.LocalVolumeReplica
	var replicaNames = lv.Status.Replicas
	for _, replicaname := range replicaNames {
		lvr := &apisv1alpha1.LocalVolumeReplica{}
		if err := lvController.Client.Get(context.TODO(), client.ObjectKey{Name: replicaname}, lvr); err != nil {
			if !errors.IsNotFound(err) {
				log.WithError(err).Error("Failed to query localvolumereplica")
			} else {
				log.Info("Not found the localvolumereplica")
			}
			return nil, err
		}
		lvrs = append(lvrs, lvr)
	}

	return lvrs, nil
}

// GetVolumeReplicas
func (lvController *LocalVolumeController) GetVolumeReplicas(lvname string) (*hwameistorapi.VolumeReplicaList, error) {
	lvrs, err := lvController.getLocalVolumeReplicas(lvname)
	if err != nil {
		log.WithError(err).Fatal("Failed to getLocalVolumeReplicas")
		return nil, err
	}

	var vrList = &hwameistorapi.VolumeReplicaList{}
	var vrs []*hwameistorapi.VolumeReplica
	for _, lvr := range lvrs {
		var vr = &hwameistorapi.VolumeReplica{}
		vr.Name = lvr.Name
		vr.NodeName = lvr.Spec.NodeName
		vr.DevicePath = lvr.Status.DevicePath
		vr.RequiredCapacityBytes = lvr.Spec.RequiredCapacityBytes
		vr.StoragePath = lvr.Status.StoragePath
		vr.Synced = lvr.Status.Synced
		vr.State = hwameistorapi.StateConvert(lvr.Status.State)
		vrs = append(vrs, vr)
	}
	vrList.VolumeReplicas = vrs
	vrList.VolumeName = lvname

	return vrList, nil
}

// GetVolumeOperation
func (lvController *LocalVolumeController) GetVolumeOperation(volumeName string) (*hwameistorapi.VolumeOperationByVolume, error) {

	var volumeOperation = &hwameistorapi.VolumeOperationByVolume{}
	var volumeMigrateOperations []*hwameistorapi.VolumeMigrateOperation
	lvmList := apisv1alpha1.LocalVolumeMigrateList{}
	if err := lvController.Client.List(context.Background(), &lvmList, &client.ListOptions{}); err != nil {
		return nil, err
	}

	for _, item := range lvmList.Items {
		if item.Spec.VolumeName == volumeName {
			var volumeMigrateOperation = &hwameistorapi.VolumeMigrateOperation{}
			volumeMigrateOperation.VolumeName = item.Spec.VolumeName
			volumeMigrateOperation.Name = item.Name
			volumeMigrateOperation.SourceNode = item.Spec.SourceNode
			if len(item.Spec.TargetNodesSuggested) != 0 {
				volumeMigrateOperation.TargetNode = item.Spec.TargetNodesSuggested[0]
			}
			volumeMigrateOperation.State = hwameistorapi.StateConvert(item.Status.State)
			volumeMigrateOperation.StartTime = item.CreationTimestamp.Time
			volumeMigrateOperations = append(volumeMigrateOperations, volumeMigrateOperation)
		}
	}

	volumeOperation.VolumeMigrateOperations = volumeMigrateOperations
	volumeOperation.VolumeName = volumeName
	return volumeOperation, nil
}

// GetLocalVolumeMigrateYamlStr
func (lvController *LocalVolumeController) GetLocalVolumeMigrateYamlStr(resourceName string) (*hwameistorapi.YamlData, error) {
	lvm := &apisv1alpha1.LocalVolumeMigrate{}
	if err := lvController.Client.Get(context.TODO(), client.ObjectKey{Name: resourceName}, lvm); err != nil {
		if !errors.IsNotFound(err) {
			log.WithError(err).Error("Failed to query localvolumemigrate")
		} else {
			log.Info("Not found the localvolumemigrate")
		}
		return nil, err
	}

	resourceYamlStr, err := lvController.getLVMResourceYaml(lvm)
	if err != nil {
		log.WithError(err).Error("Failed to getLVMResourceYaml")
		return nil, err
	}
	var yamlData = &hwameistorapi.YamlData{}
	yamlData.Data = resourceYamlStr

	return yamlData, nil
}

// GetLocalVolumeReplicaYamlStr
func (lvController *LocalVolumeController) GetLocalVolumeReplicaYamlStr(resourceName string) (*hwameistorapi.YamlData, error) {
	lvr := &apisv1alpha1.LocalVolumeReplica{}
	if err := lvController.Client.Get(context.TODO(), client.ObjectKey{Name: resourceName}, lvr); err != nil {
		if !errors.IsNotFound(err) {
			log.WithError(err).Error("Failed to query localvolumereplica")
		} else {
			log.Info("Not found the localvolumereplica")
		}
		return nil, err
	}

	resourceYamlStr, err := lvController.getLVRResourceYaml(lvr)
	if err != nil {
		log.WithError(err).Error("Failed to getLVRResourceYaml")
		return nil, err
	}
	var yamlData = &hwameistorapi.YamlData{}
	yamlData.Data = resourceYamlStr

	return yamlData, nil
}

// getLVMResourceYaml
func (lvController *LocalVolumeController) getLVMResourceYaml(lvm *apisv1alpha1.LocalVolumeMigrate) (string, error) {

	buf := new(bytes.Buffer)

	lvm.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Group:   groupName,
		Version: versionName,
		Kind:    lvm.Kind,
	})
	y := printers.YAMLPrinter{}
	err := y.PrintObj(lvm, buf)
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}

// getLVRResourceYaml
func (lvController *LocalVolumeController) getLVRResourceYaml(lvr *apisv1alpha1.LocalVolumeReplica) (string, error) {

	buf := new(bytes.Buffer)

	lvr.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
		Group:   groupName,
		Version: versionName,
		Kind:    lvr.Kind,
	})
	y := printers.YAMLPrinter{}
	err := y.PrintObj(lvr, buf)
	if err != nil {
		panic(err)
	}

	return buf.String(), nil
}
