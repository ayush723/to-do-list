package list

import (
	"errors"

	"github.com/ayush723/to-do-list/src/datasources/list_db"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

const (
	queryInsertTask = "INSERT INTO list(description) VALUES(?);"
	queryGetTask = "SELECT * FROM toDoList_db;"
	queryFindByStatusTask= "SELECT * FROM toDoList_db WHERE status=?;"
	queryUpdateTask = "UPDATE toDoList_db SET description=?, status=? WHERE id = ?;"
	queryDeleteTask = "DELETE FROM toDoList_db WHERE id=?"
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
func (toDoList *ToDoList) Get()([]ToDoList, rest_errors.RestErr){
	stmt, err := list_db.Client.Prepare(queryGetTask)
	if err != nil {
		// logger.Error("error when trying to prepare save user statement",err)
		return nil, rest_errors.NewInternalServerError("error when trying to get task", errors.New("database error"))
	}
	defer stmt.Close()
	rows, getErr := stmt.Query()
	if getErr != nil{
		return nil, rest_errors.NewInternalServerError("error when trying to get task", errors.New("database error"))
	}
	defer rows.Close()
	results := make([]ToDoList, 0)
	for rows.Next() {
		var toDoList ToDoList
		if getErr := rows.Scan(&toDoList.Id, &toDoList.Description, &toDoList.Status); getErr != nil{
			return nil, rest_errors.NewInternalServerError("error when trying to get task", errors.New("database error"))
		}
		results = append(results, toDoList)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewInternalServerError("error when trying to get task", errors.New("database error"))
	}

	return results, nil
}
func (toDoList *ToDoList) FindByStatus(status string)([]ToDoList, rest_errors.RestErr){
	stmt, err := list_db.Client.Prepare(queryFindByStatusTask)
	if err != nil{
		return nil, rest_errors.NewInternalServerError("error when trying to find user by status", errors.New("database error"))
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to find user by status", errors.New("database error"))
	}
	defer rows.Close()

	results := make ([]ToDoList, 0)

	for rows.Next(){
		var toDoList ToDoList
		if err := rows.Scan(&toDoList.Id, &toDoList.Description, &toDoList.Status); err!=nil{
		return nil, rest_errors.NewInternalServerError("error when trying to find user by status", errors.New("database error"))

		}
		results = append(results, toDoList)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewInternalServerError("error when trying to get task", errors.New("database error"))
	}
	return results, nil
}
func (toDoList *ToDoList) Update()rest_errors.RestErr{
	stmt, err := list_db.Client.Prepare(queryUpdateTask)
	if err != nil{
		return rest_errors.NewInternalServerError("error when trying to update task", errors.New("database error"))
	}
	defer stmt.Close()
	_, err = stmt.Exec(toDoList.Id, toDoList.Description, toDoList.Status)
	if err != nil{
		return rest_errors.NewInternalServerError("error when trying to update task", errors.New("database error"))
	}
	return nil
}

func (toDoList *ToDoList) Delete()rest_errors.RestErr{
	stmt, err := list_db.Client.Prepare(queryDeleteTask)
	if err != nil{
		return rest_errors.NewInternalServerError("error when trying to delete task", errors.New("database error"))
	}
	defer stmt.Close()
	_, err = stmt.Exec(toDoList.Id)
	if err != nil{
		return rest_errors.NewInternalServerError("error when trying to delete task", errors.New("database error"))
	}
	return nil
}