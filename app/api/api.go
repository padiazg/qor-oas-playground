package api

import (
	"path"

	admin "github.com/TykTechnologies/portal-admin"
	"github.com/TykTechnologies/portal-admin/oas"
	"github.com/go-oas/docs"
	"github.com/padiazg/qor-oas-playground/config/application"
	"github.com/padiazg/qor-oas-playground/config/db"
	tm "github.com/padiazg/qor-oas-playground/models/theme"
	um "github.com/padiazg/qor-oas-playground/models/user"
	"github.com/qor/roles"
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
	// API.AddResource(&um.Team{})
	// API.AddResource(&pm.Product{})
	// API.AddResource(&pm.Plan{})
	// API.AddResource(&pm.Catalogue{})

	// ar := API.AddResource(&cm.AccessRequest{})

	// ar.Meta(&admin.Meta{
	// 	Name:   "Fake",
	// 	Type:   "string",
	// 	Valuer: func(r interface{}, ctx *qor.Context) (res interface{}) { return "fake" },
	// })

	// ar.EditAttrs("Model", "Status")
	// ar.ShowAttrs("Model", "Status", "Client", "Fake")

	// API.AddResource(&cm.Credential{})

	// a := API.AddResource(&cm.Client{}, &admin.Config{Name: "App"})
	// accessRequests, _ := a.AddSubResource("AccessRequests")
	// accessRequests.AddSubResource("Credentials")

	themes := tm.Themes{
		"theme1": &tm.Theme{
			Name:    "theme1",
			Version: "1.0.0",
			Author:  "Tyk",
			Path:    "./themes/theme1",
		},
		"theme2": &tm.Theme{
			Name:    "theme2",
			Version: "1.0.0",
			Author:  "Tyk",
			Path:    "./themes/theme2",
		},
	}

	z := API.AddResource(&tm.Config{}, &admin.Config{Name: "Theme"})
	z.IndexAttrs("ID", "Name", "Status", "Author", "Version", "Path")

	z.FindManyHandler = themes.FindManyHandler
	z.FindOneHandler = themes.FindOneHandler

	z.RegisterRoute("GET",
		path.Join(z.ParamIDName(), "download"),
		tm.Download,
		&admin.RouteConfig{PermissionMode: roles.Read},
		oas.RouteConfig{
			Method:   "GET",
			Summary:  "Download a zip file",
			AddToDoc: true,
		},
	)

	z.RegisterRoute("PUT",
		path.Join(z.ParamIDName(), "upload"),
		tm.Upload,
		&admin.RouteConfig{PermissionMode: roles.Update},
		oas.RouteConfig{
			Method:   "PUT",
			Summary:  "Update a zip file",
			AddToDoc: true,
		},
	)

	z.RegisterRoute("POST",
		path.Join("upload"),
		tm.Upload,
		&admin.RouteConfig{PermissionMode: roles.Create},
		oas.RouteConfig{
			Method:   "POST",
			Summary:  "Upload a zip file",
			AddToDoc: true,
		},
	)

	application.Router.Mount(app.Config.Prefix, API.NewServeMux(app.Config.Prefix))

	API.GetRouter().PrintRoutes()

}
