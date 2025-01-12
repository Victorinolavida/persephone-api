package server

import (
	"fmt"
	"github.com/Victorinolavida/persephone-api/config"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/Victorinolavida/persephone-api/pkg/responses"
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
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		responses.NotFoundResponse(w)
	})

	r.Get("/api/v0/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		err := responses.WriteJsonResponse(w, "ok", http.StatusOK)
		if err != nil {
			responses.ServerErrorResponse(w, err)
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
