package data

type TodoListResourceResponseDTO struct {
	Id    string                     `json:"id"`
	Owner string                     `json:"owner"`
	Todos []*TodoResourceResponseDTO `json:"todos"`
}

type TodoResourceResponseDTO struct {
	Id        string `json:"id"`
	ListId    string `json:"listId,omitempty"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}