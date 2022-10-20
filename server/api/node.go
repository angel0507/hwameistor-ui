package api

import (
	"github.com/gorilla/mux"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"github.com/pkg/errors"
	"github.com/rancher/go-rancher/api"
	"github.com/rancher/go-rancher/client"
	"net/http"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func (s *Server) NodeList(rw http.ResponseWriter, req *http.Request) (err error) {
	defer func() {
		err = errors.Wrap(err, "unable to list")
	}()

	apiContext := api.GetApiContext(req)

	resp, err := s.nodeList(apiContext)
	if err != nil {
		return err
	}

	apiContext.Write(resp)

	return nil
}

func (s *Server) nodeList(apiContext *api.ApiContext) (*client.GenericCollection, error) {
	resp := &client.GenericCollection{}

	lsns, err := s.m.LocalStorageNodeController().ListLocalStorageNode()
	if err != nil {
		return nil, err
	}

	for _, lsn := range lsns.Items {

		resp.Data = append(resp.Data, toNodeResource(lsn))
	}
	resp.ResourceType = "node"
	resp.CreateTypes = map[string]string{
		"node": apiContext.UrlBuilder.Collection("node"),
	}

	return resp, nil
}

func (s *Server) NodeGet(rw http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)["name"]
	return s.responseWithNode(rw, req, id, nil)
}

func (s *Server) responseWithNode(rw http.ResponseWriter, req *http.Request, id string, v *apisv1alpha1.LocalStorageNode) error {
	var err error
	apiContext := api.GetApiContext(req)

	if v == nil {
		if id == "" {
			rw.WriteHeader(http.StatusNotFound)
			return nil
		}
		v, err = s.m.LocalStorageNodeController().GetLocalStorageNode(pkgclient.ObjectKey{Name: id})
		if err != nil {
			return errors.Wrap(err, "unable to get node")
		}
	}

	apiContext.Write(toNodeResource(*v))
	return nil
}
