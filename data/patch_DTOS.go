package data

type TodoListPatchDTO struct {
	Completed bool `json:"completed"`
}

type TodoPatchDTO struct {
	Completed bool `json:"completed" binding:"required"`
}
