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

type ResponseReadTodoAll struct {
	Id       int    `json:"id"`
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}
type UpdateTodoRequest struct {
	Id       int    `json:"id" validate:"required"`
	TodoName string `json:"todo_name" validate:"required"`
	IsCheck  bool   `json:"is_check"`
}

type DeleteTodo struct {
	Id int `json:"id" validate:"required"`
}

type ReadTodoRequest struct {
	TodoName string `json:"todo_name" validate:"required"`
	// IsCheck *bool `json:"is_check"`
}
