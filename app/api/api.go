package api

import (
	admin "github.com/TykTechnologies/portal-admin"
	"github.com/go-oas/docs"
	"github.com/padiazg/qor-oas-playground/config/application"
	"github.com/padiazg/qor-oas-playground/config/db"
	cm "github.com/padiazg/qor-oas-playground/models/client"
	pm "github.com/padiazg/qor-oas-playground/models/product"
	um "github.com/padiazg/qor-oas-playground/models/user"
)

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/api"
	}
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	API := admin.New(&admin.AdminConfig{
		DB:         db.DB,
		IsAPI:      true,
		OASVersion: "3.0.1",
		OASInfo: &docs.Info{
			Title:       "QOR OAS Playground",
			Description: "QOR OAS Playground",
			Contact:     docs.Contact{Email: "patricio@tyk.io"},
			Version:     "1.2.0",
		},
	})

	API.AddResource(&um.User{})
	API.AddResource(&um.Team{})
	API.AddResource(&pm.Product{})
	API.AddResource(&pm.Plan{})
	API.AddResource(&pm.Catalogue{})

	a := API.AddResource(&cm.Client{}, &admin.Config{Name: "app"})
	accessRequests, _ := a.AddSubResource("AccessRequests")
	accessRequests.AddSubResource("Credentials")

	application.Router.Mount(app.Config.Prefix, API.NewServeMux(app.Config.Prefix))

	API.GetRouter().PrintRoutes()
}
