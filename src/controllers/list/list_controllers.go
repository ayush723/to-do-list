package list

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayush723/to-do-list/src/domain/list"
	"github.com/ayush723/to-do-list/src/service/list_services"

	"github.com/ayush723/utils-go_bookstore/rest_errors"
	"github.com/gin-gonic/gin"
)

func getId(toDoIdParam string) (int64, rest_errors.RestErr) {

	toDoId, err := strconv.ParseInt(toDoIdParam, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("id should be a number")
	}
	return toDoId, nil
}

func Create(c *gin.Context) {
	var toDoList list.ToDoList
	if err := c.ShouldBindJSON(&toDoList); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	toDoList.Status = false
	result, saveErr := list_services.ListService.Create(toDoList)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	toDoLists, getErr := list_services.ListService.Get()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"results":toDoLists,
	})
}

func GetById(c *gin.Context) {
	toDoId, err := getId(c.Param("todo-id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	toDoList, getErr := list_services.ListService.GetById(toDoId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	c.JSON(http.StatusOK, toDoList)
}

func Update(c *gin.Context) {
	toDoId, err := getId(c.Param("todo-id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var toDoList list.ToDoList
	if err := c.ShouldBindJSON(&toDoList); err != nil {
		fmt.Println("5")
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	toDoList.Id = toDoId

	result, err := list_services.ListService.Update(toDoList)
	if err != nil {
		fmt.Println("6")

		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	toDoId, idErr := getId(c.Param("todo-id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	if err := list_services.ListService.Delete(toDoId); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "Deleted"})
}

func Search(c *gin.Context) {
	statusQuery := c.Query("status")
	status, boolErr := strconv.ParseBool(statusQuery)
	if boolErr != nil {
		c.JSON(http.StatusBadRequest, boolErr)
	}
	toDoLists, err := list_services.ListService.Search(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, toDoLists)
}
