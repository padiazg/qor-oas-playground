package accessrequest

import (
	qa "github.com/TykTechnologies/portal-admin"
	"github.com/padiazg/qor-oas-playground/config/application"
	cm "github.com/padiazg/qor-oas-playground/models/client"
)

type Config struct{}

type App struct {
	*Config
}

func New(config *Config) *App {
	return &App{Config: config}
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	application.Admin.AddResource(&cm.AccessRequest{}, &qa.Config{Menu: []string{"Access requests"}})

	app.ConfigureFuncMaps(application)
}

func (app App) ConfigureFuncMaps(application *application.Application) {
	// Admin FuncMaps
}
