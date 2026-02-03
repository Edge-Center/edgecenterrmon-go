package checkgroup

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GroupID int    `json:"group_id"`
}
