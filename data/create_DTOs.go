package data

type TodoListCreateRequestDTO struct {
	Todos []string `json:"todos" binding:"required"`
}

type TodoCreateRequestDTO struct {
	Content string `json:"content" binding:"required"`
}
