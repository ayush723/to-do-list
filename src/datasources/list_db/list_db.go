package list_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	//Set env variables
	os.Setenv("mysql_users_username", "ayush2")
	os.Setenv("mysql_users_password", "Ayush@123")
	// os.Setenv("mysql_users_host", "127.0.0.1")
	//os.Setenv("mysql_users_port", "3306")
	os.Setenv("mysql_users_schema", "toDoList_db")

	//Getting the values
	username := os.Getenv("mysql_users_username")
	password := os.Getenv("mysql_users_password")
	host := os.Getenv("mysql_users_host")
	//port := os.Getenv("mysql_users_port")
	schema := os.Getenv("mysql_users_schema")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	fmt.Println(dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
