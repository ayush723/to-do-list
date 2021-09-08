package list

type ToDoList struct {
	Id int64 `json:"id"`
	Description string `json:"description"`
	Status bool `json:"status"`
}

type ToDoLists []ToDoList