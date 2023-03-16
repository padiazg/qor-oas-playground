package user

import (
	"net/http"

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
	user := application.Admin.AddResource(&um.User{}, &qa.Config{Menu: []string{"User Management"}})

	user.EditAttrs(
		"Name",
		"Email",
		"Role",
		"Password",
	)

	user.NewAttrs(user.EditAttrs())

	app.ConfigureFuncMaps(application)
}

func (app App) ConfigureFuncMaps(application *application.Application) {
	// Admin FuncMaps
	application.Admin.RegisterFuncMap("CurrentUser", func(r *http.Request) *um.User { return CurrentUser(r, nil) })
}
