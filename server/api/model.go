package api

import (
	"github.com/hwameistor/hwameistor-ui/server/manager"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"github.com/rancher/go-rancher/client"
)

type Volume struct {
	client.Resource

	Name string `json:"name"`
	Size string `json:"size"`
}

type Node struct {
	client.Resource

	Name string `json:"name"`
	Size string `json:"size"`
}

type Server struct {
	m *manager.ServerManager
}

func NewServer(m *manager.ServerManager) *Server {
	s := &Server{
		m: m,
	}
	return s
}

func NewSchema() *client.Schemas {
	schemas := &client.Schemas{}

	schemas.AddType("apiVersion", client.Resource{})
	schemas.AddType("schema", client.Schema{})
	schemas.AddType("error", client.ServerApiError{})

	return schemas
}

func toVolumeResource(lv apisv1alpha1.LocalVolume) *Volume {
	r := &Volume{}
	return r
}

func toNodeResource(lv apisv1alpha1.LocalStorageNode) *Node {
	r := &Node{}
	return r
}
