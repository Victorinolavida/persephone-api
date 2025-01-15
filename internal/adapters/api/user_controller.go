package adapters

import (
	"encoding/json"
	"github.com/Victorinolavida/persephone-api/internal/adapters/dto"
	"github.com/Victorinolavida/persephone-api/internal/lib"
	userModel "github.com/Victorinolavida/persephone-api/internal/models/user"
	"github.com/Victorinolavida/persephone-api/internal/services/user"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/Victorinolavida/persephone-api/pkg/responses"
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
	err = c.svc.ValidateUserData(nil, userInfo)
	if err != nil {
		responses.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := userModel.CreateUserFromDTO(userInfo)
	if err != nil {
		responses.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = c.svc.Create(nil, newUser)

	if err != nil {
		responses.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}
	responses.RenderJson(w, newUser, http.StatusCreated)
}
