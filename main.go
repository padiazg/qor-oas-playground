package main

import (
	"fmt"
	"net/http"

	admin "github.com/TykTechnologies/portal-admin"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	accessrequest "github.com/padiazg/qor-oas-playground/app/access-request"
	adminapp "github.com/padiazg/qor-oas-playground/app/admin"
	"github.com/padiazg/qor-oas-playground/app/api"
	"github.com/padiazg/qor-oas-playground/app/catalogue"
	"github.com/padiazg/qor-oas-playground/app/client"
	"github.com/padiazg/qor-oas-playground/app/credential"
	"github.com/padiazg/qor-oas-playground/app/plan"
	"github.com/padiazg/qor-oas-playground/app/product"
	"github.com/padiazg/qor-oas-playground/app/static"
	"github.com/padiazg/qor-oas-playground/app/team"
	"github.com/padiazg/qor-oas-playground/app/user"
	"github.com/padiazg/qor-oas-playground/config/application"
	"github.com/padiazg/qor-oas-playground/config/db"
	_ "github.com/padiazg/qor-oas-playground/config/db/migration"
	"github.com/qor/qor/utils"
)

const (
	port = 7100
)

func main() {
	var (
		Router = chi.NewRouter()
		Admin  = admin.New(&admin.AdminConfig{
			SiteName: "QOR OAS Playground",
			// Auth:     auth.AdminAuth{},
			DB:    db.DB,
			IsAPI: false,
		})
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// for demo, don't use this for your production site
			w.Header().Add("Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)

	Application.Use(api.New(&api.Config{}))
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(user.New(&user.Config{}))
	Application.Use(team.New(&team.Config{}))

	Application.Use(catalogue.New(&catalogue.Config{}))
	Application.Use(product.New(&product.Config{}))
	Application.Use(plan.New(&plan.Config{}))

	Application.Use(client.New(&client.Config{}))
	Application.Use(accessrequest.New(&accessrequest.Config{}))
	Application.Use(credential.New(&credential.Config{}))

	Application.Use(static.New(&static.Config{
		Prefixs: []string{"/oas"},
		Handler: utils.FileServer(http.Dir("./public")),
	}))

	fmt.Printf("Listening on: %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), Application.NewServeMux()); err != nil {
		panic(err)
	}
}
