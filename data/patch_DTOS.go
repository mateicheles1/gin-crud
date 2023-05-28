package data

type TodoListPatchDTO struct {
	Owner string `json:"owner" binding:"required"`
}

type TodoPatchDTO struct {
	Completed bool `json:"completed" binding:"required"`
}
