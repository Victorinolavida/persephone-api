package responses

import (
	"encoding/json"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, data interface{}, status int) error {
	d, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(d)
	return nil
}

func RenderJson(w http.ResponseWriter, data any, status int) {
	err := WriteJsonResponse(w, data, status)
	if err != nil {
		RenderServerError(w, err)
	}
}

func RenderError(w http.ResponseWriter, status int, message any) {
	res := map[string]any{
		"error": message,
	}
	log := logger.GetLogger()
	err := WriteJsonResponse(w, res, status)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func RenderServerError(w http.ResponseWriter, err error) {
	logger.GetLogger().Error(err)
	message := "something went wrong"
	RenderError(w, http.StatusInternalServerError, message)
}

func RenderNotFound(w http.ResponseWriter) {
	message := "page not found"
	RenderError(w, http.StatusNotFound, message)
}
