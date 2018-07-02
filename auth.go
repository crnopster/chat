package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}

func (h authHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("auth")
	if err == http.ErrNoCookie {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	h.next.ServeHTTP(w, req)	
}
func MustAuth(handler http.Handler) http.Handler {
    return &authHandler{next: handler}
}