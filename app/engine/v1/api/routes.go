package api

import (
	"encoding/json"
	"github.com/hexad3cimal/mockapi/app/engine/v1/models"
	"github.com/hexad3cimal/mockapi/app/engine/v1/repository"
	"github.com/hexad3cimal/mockapi/app/engine/v1/responses"
	"io/ioutil"
	"net/http"
)

func AddApi(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(responses.Response{ResponseType: "error", Message: err.Error(), Code: "500"})
		return
	}

	api := models.Api{}
	jsonError := json.Unmarshal(body, &api)
	if jsonError != nil {
		w.WriteHeader(400)
		errorResponse, _ := json.Marshal(responses.Response{ResponseType: "error", Message:jsonError.Error(), Code: "400"})
		w.Write(errorResponse)
		return
	}

	response := repository.AddApi(api)
	 if response == "error"{
	 	w.WriteHeader(400)
		 errorResponse, _ := json.Marshal(responses.Response{ResponseType: "error", Message:"error occured", Code: "500"})
		 w.Write(errorResponse)
		 return
	 }

	w.WriteHeader(200)
    successResponse, _ := json.Marshal(responses.Response{ResponseType: "success", Message:"success", Code: "200"})
	if _, err := w.Write(successResponse); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("unable to list response mapping"))
	}
}

func AddEndpoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(responses.Response{ResponseType: "error", Message: err.Error(), Code: "500"})
		return
	}

	endPoint := models.EndPoint{}
	jsonError := json.Unmarshal(body, &endPoint)
	if jsonError != nil {
		w.WriteHeader(400)
		errorResponse, _ := json.Marshal(responses.Response{ResponseType: "error", Message:jsonError.Error(), Code: "400"})
		w.Write(errorResponse)
		return
	}

	api,err := repository.GetApi(endPoint.ApiId)
	 if err != nil{
	 	w.WriteHeader(400)
		 errorResponse, _ := json.Marshal(responses.Response{ResponseType: "error", Message:"error occured", Code: "500"})
		 w.Write(errorResponse)
		 return
	 }

	if api != (models.Api{}){



	}

	w.WriteHeader(200)
    successResponse, _ := json.Marshal(responses.Response{ResponseType: "success", Message:"success", Code: "200"})
	if _, err := w.Write(successResponse); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("unable to list response mapping"))
	}
}
