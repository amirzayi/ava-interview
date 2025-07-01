package api

import (
	"net/http"

	"github.com/amirzayi/ava-interview/service"
)

type Router struct {
	Service *service.Service
}

func (r *Router) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /user", r.ListUsers)
	mux.HandleFunc("POST /user", r.CreateUser)
	mux.HandleFunc("GET /user/{id}", r.GetUserByID)
	mux.HandleFunc("DELETE /user/{id}", r.DeleteUserByID)
}
