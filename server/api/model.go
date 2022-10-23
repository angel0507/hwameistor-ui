package api

import (
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"github.com/rancher/go-rancher/client"
)

type Volume struct {
	client.Resource

	Name string `json:"name"`
	Size string `json:"size"`
}

type VolumeList struct {
	Volumes []Volume `json:"volumes"`
}

type Node struct {
	client.Resource

	Name string `json:"name"`
}

type NodeList struct {
	Nodes []Volume `json:"nodes"`
}

func ToVolumeResource(lv apisv1alpha1.LocalVolume) *Volume {
	r := &Volume{}
	return r
}

func ToNodeResource(lv apisv1alpha1.LocalStorageNode) *Node {
	r := &Node{}
	return r
}
