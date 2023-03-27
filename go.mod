module github.com/padiazg/qor-oas-playground

go 1.19

require (
	github.com/TykTechnologies/portal-admin v0.0.0-20230316012911-4f8aae1b1ae9
	github.com/go-chi/chi v1.5.4
	github.com/go-oas/docs v1.1.0
	github.com/jinzhu/gorm v1.9.15
	github.com/qor/middlewares v0.0.0-20170822143614-781378b69454
	github.com/qor/qor v1.2.0
	github.com/qor/wildcard_router v0.0.0-20171031035524-56710e5bb5a4
)

require (
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/gin-gonic/gin v1.9.0 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/gosimple/slug v1.9.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.23 // indirect
	github.com/qor/assetfs v0.0.0-20170713023933-ff57fdc13a14 // indirect
	github.com/qor/responder v0.0.0-20171031032654-b6def473574f // indirect
	github.com/qor/roles v0.0.0-20171127035124-d6375609fe3e // indirect
	github.com/qor/session v0.0.0-20170907035918-8206b0adab70 // indirect
	github.com/qor/validations v0.0.0-20171228122639-f364bca61b46 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/theplant/cldr v0.0.0-20190423050709-9f76f7ce4ee8 // indirect
	golang.org/x/net v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/TykTechnologies/portal-admin => ../../TykTechnologies/portal-admin

// replace github.com/go-oas/docs => ../../go-oas/docs
replace github.com/go-oas/docs => ../../padiazg/docs
