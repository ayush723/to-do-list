package list

type ToDoList struct {
	Id int64 `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
}
