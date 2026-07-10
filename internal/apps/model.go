package apps

import "context"

// Application is the API representation returned by GET /apps.
type Application struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Icon        string `json:"icon"`
}

// Source returns the list of registered applications.
type Source interface {
	ListApplications(context.Context) ([]Application, error)
}
