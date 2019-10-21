package repository

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type DbMock struct{
	mock.Mock
}
func getEnPointsTest(t *testing.T){

	endPointsDb := new(DbMock)

	// setup expectations
	endPointsDb.On("query", "some-id").Return("true", nil)

	//GetEndpoints()


}