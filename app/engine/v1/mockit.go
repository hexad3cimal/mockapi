package main

import (
	"github.com/hexad3cimal/mockapi/app/engine/v1/repository"
)

func main() {

	dbConnection := repository.DbConnection()

	statement, _ := dbConnection.Prepare("CREATE TABLE IF NOT EXISTS api (id VARCHAR(50) PRIMARY KEY, url TEXT, created DATE,updated DATE, active INTEGER)")
	statement.Exec()

	//_, err := dbConnection.Exec(
	//	"CREATE TABLE `api` (`id` VARCHAR(50) PRIMARY KEY, `url` VARCHAR(255) NULL")
	//
	//if err != nil {
	//	fmt.Errorf(err.Error())
	//}
	//
	//stmt, err2 := dbConnection.Prepare("INSERT INTO api(id, url) values(?,?)")
	//
	//if err2 != nil {
	//	fmt.Errorf(err2.Error())
	//}
	//
	//if stmt == nil {
	//	fmt.Errorf("sfsdfsdfds")
	//}
	//
	//res, err := stmt.Exec("astaxie", "1111")
	//
	//if res != nil {
	//	fmt.Errorf(err.Error())
	//}

	dbConnection.Close()

}

