package httpapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xkovacikm2/kube-app-manager/internal/apps"
)

type Handler struct {
	applicationSource apps.Source
}

func NewHandler(applicationSource apps.Source) http.Handler {
	handler := &Handler{applicationSource: applicationSource}
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("GET /apps", handler.handleApps)
	return serveMux
}

func (handler *Handler) handleApps(responseWriter http.ResponseWriter, request *http.Request) {
	applicationList, err := handler.applicationSource.ListApplications(request.Context())
	if err != nil {
		log.Printf("failed to list applications: %v", err)
		writeJSON(responseWriter, http.StatusInternalServerError, map[string]string{"error": "failed to load applications"})
		return
	}

	writeJSON(responseWriter, http.StatusOK, applicationList)
}

func writeJSON(responseWriter http.ResponseWriter, statusCode int, payload any) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)

	if err := json.NewEncoder(responseWriter).Encode(payload); err != nil {
		log.Printf("failed to write json response: %v", err)
	}
}
