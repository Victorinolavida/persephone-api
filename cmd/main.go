package main

import (
	"github.com/Victorinolavida/persephone-api/internal/config"
	"github.com/Victorinolavida/persephone-api/internal/infrastructure"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/Victorinolavida/persephone-api/pkg/server"

	"net/http"
)

func main() {
	log := logger.NewLogger(nil)
	defer log.Sync()

	conf, err := config.NewConfig()
	if err != nil {
		log.Errorf(err.Error())
	}

	_, err = infrastructure.NewDB(conf.DB)
	if err != nil {
		log.Fatalf(err.Error())
	}
	r := server.NewServer()

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hola mundo"))
	})

	err = r.Start(conf)

	if err != nil {
		log.Fatalf(err.Error())
	}

}
