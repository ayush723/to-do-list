package list_services

import (
	"fmt"

	"github.com/ayush723/to-do-list/src/domain/list"
	"github.com/ayush723/utils-go_bookstore/rest_errors"
)

var (
	ListService listServiceInterface = &listService{}
)

type listServiceInterface interface {
	Create(list.ToDoList) (*list.ToDoList, rest_errors.RestErr)
	Get() (list.ToDoLists, rest_errors.RestErr)
	GetById(int64) (*list.ToDoList, rest_errors.RestErr)
	Update(list.ToDoList) (*list.ToDoList, rest_errors.RestErr)
	Delete(int64) rest_errors.RestErr
	Search(bool) (list.ToDoLists, rest_errors.RestErr)
}

type listService struct{}

func (s *listService) Create(toDoList list.ToDoList) (*list.ToDoList, rest_errors.RestErr) {
	if err := toDoList.Save(); err != nil {
		return nil, err
	}
	return &toDoList, nil
}

func (s *listService) GetById(id int64) (*list.ToDoList, rest_errors.RestErr) {
	dao := &list.ToDoList{}
	return dao.GetById(id)
}

func (s *listService) Get() (list.ToDoLists, rest_errors.RestErr) {
	dao := &list.ToDoList{}
	
	return dao.Get()
}

func (s *listService) Update(toDoList list.ToDoList) (*list.ToDoList, rest_errors.RestErr) {
	current, err := ListService.GetById(toDoList.Id)
	if err != nil {
		fmt.Println("7")
		return nil, err
	}

	toDoList.Description = current.Description

	err = toDoList.Update()
	if err != nil {
		fmt.Println("8")

		return nil, err
	}
	return &toDoList, nil
}

func (s *listService) Delete(toDoId int64) rest_errors.RestErr {
	toDoList := &list.ToDoList{Id: toDoId}
	return toDoList.Delete()
}

func (s *listService) Search(status bool) (list.ToDoLists, rest_errors.RestErr) {
	dao := &list.ToDoList{}
	return dao.FindByStatus(status)
}
