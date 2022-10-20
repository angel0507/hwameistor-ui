package api

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rancher/go-rancher/api"
	"github.com/rancher/go-rancher/client"
	"github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	"time"
)

var (
	RetryCounts   = 5
	RetryInterval = 100 * time.Millisecond
)

type HandleFuncWithError func(http.ResponseWriter, *http.Request) error

func HandleError(s *client.Schemas, t HandleFuncWithError) http.Handler {
	return api.ApiHandler(s, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var err error
		for i := 0; i < RetryCounts; i++ {
			err = t(rw, req)
			if !apierrors.IsConflict(errors.Cause(err)) {
				break
			}
			logrus.Warnf("Retry API call due to conflict")
			time.Sleep(RetryInterval)
		}
		if err != nil {
			logrus.Warnf("HTTP handling error %v", err)
			apiContext := api.GetApiContext(req)
			apiContext.WriteErr(err)
		}
	}))
}

func NewRouter(s *Server) *mux.Router {
	schemas := NewSchema()
	r := mux.NewRouter().StrictSlash(true)
	f := HandleError

	r.Methods("GET").Path("/v1/volumes").Handler(f(schemas, s.VolumeList))
	r.Methods("GET").Path("/v1/volumes/{name}").Handler(f(schemas, s.VolumeGet))

	r.Methods("GET").Path("/v1/nodes").Handler(f(schemas, s.NodeList))
	r.Methods("GET").Path("/v1/nodes/{name}").Handler(f(schemas, s.NodeGet))

	return r
}
