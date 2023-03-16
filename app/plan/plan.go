package plan

import (
	qa "github.com/TykTechnologies/portal-admin"
	"github.com/padiazg/qor-oas-playground/config/application"
	pm "github.com/padiazg/qor-oas-playground/models/product"
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
	application.Admin.AddResource(&pm.Plan{}, &qa.Config{Menu: []string{"Catalogue"}})

	app.ConfigureFuncMaps(application)
}

func (app App) ConfigureFuncMaps(application *application.Application) {
	// Admin FuncMaps
}
