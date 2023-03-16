package team

import (
	qa "github.com/TykTechnologies/portal-admin"
	"github.com/padiazg/qor-oas-playground/config/application"
	um "github.com/padiazg/qor-oas-playground/models/user"
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
	application.Admin.AddResource(&um.Team{}, &qa.Config{Menu: []string{"User Management"}})

	app.ConfigureFuncMaps(application)
}

func (app App) ConfigureFuncMaps(application *application.Application) {
	// Admin FuncMaps
}
