package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hexad3cimal/mockapi/app/engine/v1/models"
	"github.com/hexad3cimal/mockapi/app/engine/v1/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Route models.Route
var Router = mux.NewRouter()

func Post(w http.ResponseWriter, r *http.Request) {
	log.Info("****Post*****")

	var errString string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errString = "unable to parse request body"
		w.WriteHeader(500)
		w.Write([]byte(errString))
	}

	route := models.Route{}
	err = json.Unmarshal(body, &route)
	if err != nil {
		errString = "unable to unmarshal body"
		w.WriteHeader(500)
		w.Write([]byte(errString))
	}

	utils.Create(route);
	Router.HandleFunc(route.BaseUrl, utils.RouteHandler).Methods(route.Methods...)

}
