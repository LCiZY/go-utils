package model

type Note struct {
	ID          string `json:"id,omitempty"`
	Content     string `json:"content"`
	LastModTime string `json:"last_mod_time"`
	CreateTme   string `json:"create_tme"`
}
