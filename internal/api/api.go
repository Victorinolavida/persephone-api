package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserController struct {
}

func NewUserController(c chi.Router) *UserController {
	controller := &UserController{}
	c.Get("/", func(writer http.ResponseWriter, request *http.Request) {
	})
	return controller
}
