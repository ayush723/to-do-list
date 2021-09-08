package app

import (
	"github.com/ayush723/to-do-list/src/controllers/list"
	"github.com/ayush723/to-do-list/src/controllers/ping"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)
	router.POST("/todo", list.Create)
}