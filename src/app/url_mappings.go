package app

import (
	"github.com/ayush723/to-do-list/src/controllers/list"
	"github.com/ayush723/to-do-list/src/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)             //done
	router.POST("/todo", list.Create)          //done
	router.GET("/todo", list.Get)              //done
	router.GET("/todo/:todo-id", list.GetById) //done
	router.PATCH("/todo/:todo-id", list.Update) // done
	router.DELETE("/todo/:todo-id", list.Delete) //done
	router.GET("/todo/search", list.Search) // done
}
