package config

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/hexad3cimal/mockapi/app/engine/v1/api"
	"github.com/hexad3cimal/mockapi/app/engine/v1/exceptions"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	dbConnection := DbConnection()
	statement, _ := dbConnection.Prepare("CREATE TABLE IF NOT EXISTS api (id VARCHAR(50) PRIMARY KEY, url TEXT, created DATE,updated DATE, active INTEGER)")
	_,err := statement.Exec()
	if err!= nil{
		log.Error(exceptions.Exception{TypeOfException: "INITIALIZEDBERROR:", Message: err.Error(), Timestamp: time.Now()})
	}
	dbConnection.Close()
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}


func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api", api.PostApi).Methods("POST")
}


func DbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./mock-api-test.db")

	if err != nil {
		log.Error(exceptions.Exception{TypeOfException: "DBEXCEPTION", Message: err.Error(), Timestamp: time.Now()})
	}

	return db
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))

}
