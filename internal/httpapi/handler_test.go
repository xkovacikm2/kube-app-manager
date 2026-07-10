package httpapi

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xkovacikm2/kube-app-manager/internal/apps"
)

type fakeSource struct {
	applicationList []apps.Application
	err             error
}

func (source *fakeSource) ListApplications(context.Context) ([]apps.Application, error) {
	if source.err != nil {
		return nil, source.err
	}
	return source.applicationList, nil
}

func Test_handleApps_success(t *testing.T) {
	handler := NewHandler(&fakeSource{applicationList: []apps.Application{{
		Name:        "portal",
		Description: "Main portal",
		URL:         "https://portal.example.com",
		Icon:        "https://portal.example.com/icon.png",
	}}})

	request := httptest.NewRequest(http.MethodGet, "/apps", nil)
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	expectedBody := "[{\"name\":\"portal\",\"description\":\"Main portal\",\"url\":\"https://portal.example.com\",\"icon\":\"https://portal.example.com/icon.png\"}]\n"
	if responseRecorder.Body.String() != expectedBody {
		t.Fatalf("unexpected body: %s", responseRecorder.Body.String())
	}

	if responseRecorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("expected json content-type, got %s", responseRecorder.Header().Get("Content-Type"))
	}
}

func Test_handleApps_sourceError(t *testing.T) {
	handler := NewHandler(&fakeSource{err: errors.New("boom")})

	request := httptest.NewRequest(http.MethodGet, "/apps", nil)
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, responseRecorder.Code)
	}

	expectedBody := "{\"error\":\"failed to load applications\"}\n"
	if responseRecorder.Body.String() != expectedBody {
		t.Fatalf("unexpected body: %s", responseRecorder.Body.String())
	}
}
