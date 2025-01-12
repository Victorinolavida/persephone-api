package server

import (
	"fmt"
	"github.com/Victorinolavida/persephone-api/internal/config"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	*chi.Mux
}

func NewServer() *Server {
	log := logger.GetLogger()
	//
	r := chi.NewRouter()
	Logger := loggerMiddleware(log)
	r.Use(Logger)
	r.Get("/healthcheck", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("ok"))
		// write status
		if err != nil {
			log.Errorf(err.Error())
		}
	})
	return &Server{
		r,
	}

}

func (s *Server) Start(config config.Config) error {
	port := fmt.Sprintf(":%d", config.Server.Port)
	log := logger.GetLogger()
	log.Infof("starting server on %s", port)
	err := http.ListenAndServe(port, s)
	if err != nil {
		return err
	}

	return nil
}

func loggerMiddleware(log *zap.SugaredLogger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info(
				"method: ", r.Method,
				"	",
				"path: ", r.URL.Path,
			)
			next.ServeHTTP(w, r)
		})
	}
}
