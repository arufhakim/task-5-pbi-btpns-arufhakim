package response

type Photo struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}
