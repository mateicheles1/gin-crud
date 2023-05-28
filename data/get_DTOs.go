package data

type TodoListGetResponseDTO struct {
	Owner string
	Todos []*TodoGetResponseInListDTO
}

type TodoGetResponseInListDTO struct {
	Id        string
	Content   string
	Completed bool
}

type TodoGetResponseDTO struct {
	ListId    string
	Content   string
	Completed bool
}
