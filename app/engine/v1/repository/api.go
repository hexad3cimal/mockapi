package repository

import (
	"database/sql"
	"github.com/hexad3cimal/mockapi/app/engine/v1/exceptions"
	"github.com/hexad3cimal/mockapi/app/engine/v1/models"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"time"
)

func DbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./mock-api.db")

	if err != nil {
		log.Error(exceptions.Exception{TypeOfException: "DBEXCEPTION", Message: err.Error(), Timestamp: time.Now()})
	}

	return db
}

func GetEndpoints(apiId string) []models.EndPoint {

	connection := DbConnection()
	var endPoints[] models.EndPoint
	data, err := connection.Query("SELECT * FROM endpoints WHERE api_id=?", apiId)

	if err!= nil{
		log.Error(exceptions.Exception{TypeOfException: "GETENDPOINTS", Message: err.Error(), Timestamp: time.Now()})
	}

	for data.Next() {
		var endPoint models.EndPoint

		if err := data.Scan(&endPoint.Id, &endPoint.URL, &endPoint.ApiId,&endPoint.Headers,&endPoint.RequestBody,&endPoint.ResponseBody,&endPoint.Method,&endPoint.StatusCode); err != nil {
			log.Error(exceptions.Exception{TypeOfException: "GETENDPOINTSDB", Message: err.Error(), Timestamp: time.Now()})
		}
		endPoints = append(endPoints, endPoint)
	}

	connection.Close()
	return endPoints
}

func AddApi(api models.Api) string {
	connection := DbConnection()
	statement, _ := connection.Prepare("INSERT INTO api (id, url,created,updated,active) VALUES (?, ?,?,?,?)")
	_, err :=statement.Exec(api.Id,api.Url,api.CreatedAt,api.UpdatedAt,api.Active)

if err!=nil{
	log.Error(exceptions.Exception{TypeOfException: "ADDAPIDB", Message: err.Error(), Timestamp: time.Now()})

	return "error"
}

	return "success"
}