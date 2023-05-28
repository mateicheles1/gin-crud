package data

type TodoListGetResponseDTO struct {
	Owner     string                      `json:"owner"`
	Completed bool                        `json:"completed"`
	Todos     []*TodoGetResponseInListDTO `json:"todos"`
}

type TodoGetResponseInListDTO struct {
	Id        string `json:"todoId"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type TodoGetResponseDTO struct {
	ListId    string `json:"todoListId"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
