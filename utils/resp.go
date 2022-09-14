package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LCiZY/go-utils/model"
)

func SetAllowCROS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Expose-Headers", "*")
}

func MakeSuccessResp(w *http.ResponseWriter, data interface{}) {
	resp := model.BaseResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	str, _ := json.Marshal(resp)
	fmt.Fprint(*w, string(str))
}

func MakeErrorResp(w *http.ResponseWriter, msg string) {
	resp := model.BaseResp{
		Code: -1,
		Msg:  msg,
	}
	str, _ := json.Marshal(resp)
	fmt.Fprint(*w, string(str))
}

func MakeResp(w *http.ResponseWriter, code int, msg string) {
	resp := model.BaseResp{
		Code: code,
		Msg:  msg,
	}
	str, _ := json.Marshal(resp)
	fmt.Fprint(*w, string(str))
}
