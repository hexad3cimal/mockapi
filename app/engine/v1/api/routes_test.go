package api

import (
	"github.com/hexad3cimal/mockapi/app/engine/v1/models"
	"github.com/hexad3cimal/mockapi/app/engine/v1/repository"
	"testing"
)

func ApiPostTest(t *testing.T) {
var addApi = repository.AddApi
addApiMock := addApi

defer func() {  addApi = addApiMock }()

	addApiMock = func(api models.Api) string{
		return "success"

}

}
