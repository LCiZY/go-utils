package model

type Note struct {
	ID          string `json:"id,omitempty"`
	RPath       string `json:"r_path,omitempty"`
	Content     string `json:"content"`
	LastModTime string `json:"last_mod_time"`
	CreateTme   string `json:"create_tme"`
}
