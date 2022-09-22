package model

type TreeDir struct {
	Label       string     `json:"label,omitempty"`
	Children    []*TreeDir `json:"children,omitempty"`
	RPath       string     `json:"r_path,omitempty"`
	APath       string     `json:"-"`
	IsDir       bool       `json:"is_dir,omitempty"`
	IsHide      bool       `json:"is_hide,omitempty"`
	LastModTime string     `json:"last_mod_time,omitempty"`
}
