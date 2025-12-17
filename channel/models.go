package channel

type Request struct {
	Channel string `json:"channel_name"`
	Token   string `json:"token"`
}

type Response struct {
	ID      int    `json:"id"`
	Channel string `json:"channel_name"`
	GroupID int    `json:"group_id"`
	Token   string `json:"token"`
}
