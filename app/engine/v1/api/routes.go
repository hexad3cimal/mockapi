package api

import (
	"encoding/json"
	"github.com/hexad3cimal/mockapi/app/engine/v1/models"
	"github.com/hexad3cimal/mockapi/app/engine/v1/repository"
	"github.com/hexad3cimal/mockapi/app/engine/v1/responses"
	"io/ioutil"
	"net/http"
)

func ApiPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(responses.Response{ResponseType: "error", Message: err.Error(), Code: "500"})
	}

	api := models.Api{}
	err = json.Unmarshal(body, &api)

	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(responses.Response{ResponseType: "error", Message: err.Error(), Code: "500"})
	}

	response := repository.AddApi(api)
	 if response == "error"{
		 w.WriteHeader(400)
		 _ = json.NewEncoder(w).Encode(responses.Response{ResponseType: "error", Message:"error occured", Code: "400"})
	 }

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(responses.Response{ResponseType: "success", Message:"success", Code: "200"})
}
