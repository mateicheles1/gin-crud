package data

type TodoListResourceResponseDTO struct {
	Id        string                     `json:"id"`
	UserId    string                     `json:"userId"`
	Owner     string                     `json:"owner"`
	Completed bool                       `json:"completed"`
	Todos     []*TodoResourceResponseDTO `json:"todos"`
}

type TodoResourceResponseDTO struct {
	Id        string `json:"id"`
	ListId    string `json:"listId,omitempty"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
