package dtos

type Task struct {
	ID      int    `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}
