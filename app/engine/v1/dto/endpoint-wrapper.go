package dto

import "github.com/hexad3cimal/mockapi/app/engine/v1/models"

type EndpointWrapper struct{

	ApiId string `json:"api_id"`
	Endpoints[] models.EndPoint `json:"endpoints"`
}
