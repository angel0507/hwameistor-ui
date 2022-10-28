package api

import (
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
)

// volume
type Volume struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

// volumeList
type VolumeList struct {
	Volumes []*Volume `json:"volumes"`
}

// node
type Node struct {
	Name string `json:"name"`
}

// nodeList
type NodeList struct {
	Nodes []*Node `json:"nodes"`
}

func ToVolumeResource(lv apisv1alpha1.LocalVolume) *Volume {
	r := &Volume{}
	r.Name = lv.Name
	r.Size = lv.Spec.RequiredCapacityBytes
	return r
}

func ToNodeResource(lsn apisv1alpha1.LocalStorageNode) *Node {
	r := &Node{}
	r.Name = lsn.Name
	return r
}
