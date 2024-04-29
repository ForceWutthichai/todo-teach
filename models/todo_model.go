package models

type CreateTodoRequest struct {
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}

type ResponseReadTodo struct {
	Id       int    `json:"id"`
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}
