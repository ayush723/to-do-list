package list_services

import (
	"github.com/ayush723/to-do-list/src/domain/list"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

var(
	ListService listServiceInterface = &listService{}
)

type listServiceInterface interface{
	Create(list.ToDoList)(*list.ToDoList, rest_errors.RestErr)

}

type listService struct{}

func (s *listService)Create(toDoList list.ToDoList)(*list.ToDoList, rest_errors.RestErr){
	if err := toDoList.Save(); err != nil{
		return nil, err
	}
	return &toDoList, nil
}

func (s *listService) Get()(list.ToDoLists, rest_errors.RestErr){
	dao := &list.ToDoList{}
	return dao.Get()
}

func (s *listService) Update(toDoList list.ToDoList)(list.ToDoList, rest_errors.RestErr){
	
}