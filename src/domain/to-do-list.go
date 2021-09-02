package domain

type toDoList struct {
	Id int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
}
