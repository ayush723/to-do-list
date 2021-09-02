package list

import (
	"net/http"

	"github.com/ayush723/to-do-list/src/domain/list"
	"github.com/ayush723/to-do-list/src/services/list_services"

	"github.com/ayush723/utils-go_bookstore/rest_errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context){
	var toDoList list.ToDoList
	if err := c.ShouldBindJSON(&toDoList); err != nil{
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, saveErr := list_services.ListService.Create(toDoList)
	if saveErr != nil{
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context){

}


func Update(c *gin.Context){

}


func Delete(c *gin.Context){

}

func Search(c *gin.Context){
	
}