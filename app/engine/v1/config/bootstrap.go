package config

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/hexad3cimal/mockapi/app/engine/v1/api"
	"github.com/hexad3cimal/mockapi/app/engine/v1/exceptions"
	"github.com/hexad3cimal/mockapi/app/engine/v1/utils"
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
	dbConnection := utils.DbConnection()
	statement, _ := dbConnection.Prepare("CREATE TABLE IF NOT EXISTS api (id VARCHAR(50) PRIMARY KEY, url TEXT, created DATE,updated DATE, active INTEGER)")
	_,err := statement.Exec()
	if err!= nil{
		log.Error(exceptions.Exception{TypeOfException: "INITIALIZEDBERROR:", Message: err.Error(), Timestamp: time.Now()})
	}
	dbConnection.Close()
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}


func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/api", api.AddApi).Methods("POST")
	a.Router.HandleFunc("/endpoint", api.AddEndpoint).Methods("POST")
}



func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))

}
