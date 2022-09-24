package model

type BaseResp struct {
	Code int         `json:"code"` // if code have omitempty tag, it would be ignored when it is assigned to 0
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
