package data

type TodoListResourceResponseDTO struct {
	Id        string                     `json:"id"`
	UserId    string                     `json:"userId,omitempty"`
	Owner     string                     `json:"owner"`
	Completed bool                       `json:"completed"`
	Todos     []*TodoResourceResponseDTO `json:"todos"`
}

type TodoResourceResponseDTO struct {
	Id        string `json:"id"`
	ListId    string `json:"todoListId,omitempty"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
