package adapters

import (
	"encoding/json"
	"github.com/Victorinolavida/persephone-api/internal/adapters/dto"
	"github.com/Victorinolavida/persephone-api/internal/lib"
	"github.com/Victorinolavida/persephone-api/internal/services/user"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/Victorinolavida/persephone-api/pkg/server"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserController struct {
	svc       user.Service
	router    *chi.Mux
	validator *lib.Validator
}

func NewUserController(router *server.Server, svc user.Service, v *lib.Validator) UserController {
	c := UserController{
		router:    router.Mux,
		svc:       svc,
		validator: v,
	}
	c.router.Group(
		func(r chi.Router) {
			r.Post("/api/v0/user/signin", c.handleCreateNewUser)
		},
	)
	return c
}

func (c *UserController) handleCreateNewUser(w http.ResponseWriter, r *http.Request) {
	var userInfo dto.User
	log := logger.GetLogger()

	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		log.Error(err)
		return
	}
	err = c.validator.ValidateStruct(userInfo)
	if err != nil {
		log.Error(err)
		return
	}

	err = c.validator.UserPassword(userInfo)
	if err != nil {
		log.Error(err)
		return
	}

	log.Debug("user data %v", userInfo)
	w.Write([]byte("ok"))
}
