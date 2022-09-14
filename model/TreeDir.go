package model

type TreeDir struct {
	Label       string
	Children    []*TreeDir
	RPath       string
	APath       string
	IsDir       bool
	IsHide      bool
	LastModTime string
}
