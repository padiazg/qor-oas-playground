package admin

import (
	admin "github.com/TykTechnologies/portal-admin"
	"github.com/padiazg/qor-oas-playground/config/application"
)

type Config struct {
	Prefix string
}

type App struct {
	*Config
}

func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/admin"
	}

	return &App{Config: config}
}

func (app App) ConfigureApplication(application *application.Application) {
	// assign new Admin insantance
	application.Admin = admin.New(&admin.AdminConfig{
		DB:    application.DB,
		IsAPI: false,
	})

	application.Router.Mount(app.Config.Prefix, application.Admin.NewServeMux(app.Config.Prefix))
}
