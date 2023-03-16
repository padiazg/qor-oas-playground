package migration

import (
	"github.com/padiazg/qor-oas-playground/config/db"
	"github.com/padiazg/qor-oas-playground/models/client"
	"github.com/padiazg/qor-oas-playground/models/product"
	"github.com/padiazg/qor-oas-playground/models/user"
)

func init() {
	AutoMigrate(&user.User{}, &user.Team{})
	AutoMigrate(&product.Product{}, &product.Plan{}, &product.Catalogue{})
	AutoMigrate(&client.AccessRequest{}, &client.Client{}, &client.Credential{})
}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
