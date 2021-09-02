package domain

import (
	"errors"

	"github.com/ayush723/to-do-list/src/datasource/list_db"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

const (
	queryInsertTask = "INSERT INTO list(description) VALUES(?);"
	queryGetTask = ""
)



func (toDoList *ToDoList) Save() rest_errors.RestErr{
	stmt, err := list_db.Client.Prepare(queryInsertTask)
	if err != nil {
		// logger.Error("error when trying to prepare save user statement",err)
		return rest_errors.NewInternalServerError("error when trying to save new task", errors.New("database error"))
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(toDoList.Description)
	if saveErr != nil {
		// logger.Error("error when trying to save user",saveErr)
		return rest_errors.NewInternalServerError("error when trying to save new task", errors.New("database error"))
	}
	taskId, err := insertResult.LastInsertId()
	if err != nil {
		// logger.Error("error when trying to get last insert id after creating a new user",err)
		return rest_errors.NewInternalServerError("error when trying to save user", errors.New("database error"))
	}

	toDoList.Id = taskId

	return nil
}
func (toDoList *ToDoList) Get(){}
func (toDoList *ToDoList) FindByStatus(){}
func (toDoList *ToDoList) Update(){}
func (toDoList *ToDoList) Delete(){}