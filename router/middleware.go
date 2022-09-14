package router

import (
	"net/http"

	"github.com/LCiZY/go-utils/utils"
)

func Get(pattern string, f http.HandlerFunc) {
	http.HandleFunc(pattern, defaultInterceptor(get(f)))
}

func Post(pattern string, f http.HandlerFunc) {
	http.HandleFunc(pattern, defaultInterceptor(post(f)))
}

func get(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			utils.MakeErrorResp(&w, "method not allowed")
			return
		}
		h(w, r)
	}
}

func post(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			utils.MakeErrorResp(&w, "method not allowed")
			return
		}
		h(w, r)
	}
}

func defaultInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetAllowCROS(&w)
		h(w, r)
	}
}
