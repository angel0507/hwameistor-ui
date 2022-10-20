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

func (s *Server) VolumeList(rw http.ResponseWriter, req *http.Request) (err error) {
	defer func() {
		err = errors.Wrap(err, "unable to list")
	}()

	apiContext := api.GetApiContext(req)

	resp, err := s.volumeList(apiContext)
	if err != nil {
		return err
	}

	apiContext.Write(resp)

	return nil
}

func (s *Server) volumeList(apiContext *api.ApiContext) (*client.GenericCollection, error) {
	resp := &client.GenericCollection{}

	lvs, err := s.m.LocalVolumeController().ListLocalVolume()
	if err != nil {
		return nil, err
	}

	for _, v := range lvs.Items {

		resp.Data = append(resp.Data, toVolumeResource(v))
	}
	resp.ResourceType = "volume"
	resp.CreateTypes = map[string]string{
		"volume": apiContext.UrlBuilder.Collection("volume"),
	}

	return resp, nil
}

func (s *Server) VolumeGet(rw http.ResponseWriter, req *http.Request) error {
	id := mux.Vars(req)["name"]
	return s.responseWithVolume(rw, req, id, nil)
}

func (s *Server) responseWithVolume(rw http.ResponseWriter, req *http.Request, id string, v *apisv1alpha1.LocalVolume) error {
	var err error
	apiContext := api.GetApiContext(req)

	if v == nil {
		if id == "" {
			rw.WriteHeader(http.StatusNotFound)
			return nil
		}
		v, err = s.m.LocalVolumeController().GetLocalVolume(pkgclient.ObjectKey{Name: id})
		if err != nil {
			return errors.Wrap(err, "unable to get volume")
		}
	}

	apiContext.Write(toVolumeResource(*v))
	return nil
}
