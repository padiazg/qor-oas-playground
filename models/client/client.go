package client

import (
	"github.com/jinzhu/gorm"
	pm "github.com/padiazg/qor-oas-playground/models/product"
)

type Client struct {
	gorm.Model
	Name           string       `gorm:"column:name"`
	Description    string       `gorm:"column:description"`
	RedirectURLs   string       `gorm:"column:redirect_urls"`
	Products       []pm.Product `gorm:"many2many:client_products;"`
	Credentials    []Credential
	AccessRequests []AccessRequest
	UserID         uint `gorm:"column:user_id"` // Requires a reverse lookup
}
