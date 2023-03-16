package application

import (
	"net/http"

	admin "github.com/TykTechnologies/portal-admin"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/qor/middlewares"
	"github.com/qor/wildcard_router"
)

// MicroAppInterface micro app interface
type MicroAppInterface interface {
	ConfigureApplication(*Application)
}

// Config application config
type Config struct {
	Router   *chi.Mux
	Handlers []http.Handler
	Admin    *admin.Admin
	DB       *gorm.DB
}

// Application main application
type Application struct {
	*Config
}

// New return new application instance
func New(config *Config) *Application {
	if config == nil {
		config = &Config{}
	}

	if config.Router == nil {
		config.Router = chi.NewRouter()
	}

	return &Application{Config: config}
}

// Use mount router into micro app
func (application *Application) Use(app MicroAppInterface) {
	app.ConfigureApplication(application)
}

func (application *Application) NewServeMux() http.Handler {
	if len(application.Config.Handlers) == 0 {
		return middlewares.Apply(application.Config.Router)
	}

	wildcardRouter := wildcard_router.New()
	for _, handler := range application.Config.Handlers {
		wildcardRouter.AddHandler(handler)
	}
	wildcardRouter.AddHandler(application.Config.Router)

	return middlewares.Apply(wildcardRouter)
}
