package list_db

import (
	"database/sql"
	"fmt"
	"log"
)

var(
	Client *sql.DB
)

func init(){
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "ayush", "envy", "127.0.0.1", "toDoList_db")
	var err error
	Client, err = sql.Open("mysql",dataSourceName)
	if err!= nil{
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}