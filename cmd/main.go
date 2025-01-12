package main

import (
	"github.com/Victorinolavida/persephone-api/config"
	adapters "github.com/Victorinolavida/persephone-api/internal/adapters/api"
	"github.com/Victorinolavida/persephone-api/internal/infrastructure"
	"github.com/Victorinolavida/persephone-api/internal/lib"
	"github.com/Victorinolavida/persephone-api/internal/repositories/user"
	userSvc "github.com/Victorinolavida/persephone-api/internal/services/user"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/Victorinolavida/persephone-api/pkg/server"
)

func main() {
	log := logger.NewLogger(nil)
	defer log.Sync()

	conf, err := config.NewConfig()
	if err != nil {
		log.Errorf(err.Error())
	}

	db, err := infrastructure.NewDB(conf.DB)
	if err != nil {
		log.Fatalf(err.Error())
	}
	r := server.NewServer()

	userRepository := user.NewUserRepository(db)

	userService := userSvc.NewUserService(userRepository)
	validator := lib.NewValidator()

	adapters.NewUserController(r, userService, validator)

	err = r.Start(conf)

	if err != nil {
		log.Fatalf(err.Error())
	}

}
