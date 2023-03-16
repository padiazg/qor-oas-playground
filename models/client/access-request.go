package client

import (
	"github.com/jinzhu/gorm"
	pm "github.com/padiazg/qor-oas-playground/models/product"
	um "github.com/padiazg/qor-oas-playground/models/user"
)

type AccessRequest struct {
	gorm.Model
	User                 um.User
	Credentials          []Credential
	Products             []pm.Product `gorm:"many2many:product_requests;"`
	Catalogue            pm.Catalogue
	Client               Client
	Status               string  `gorm:"column:status"`
	AuthType             string  `gorm:"column:auth_type"`
	Plan                 pm.Plan `gorm:"foreignKey:PlanID"`
	UserID               uint    `gorm:"column:user_id"`
	ClientID             uint    `gorm:"column:client_id"`
	PlanID               uint    `gorm:"column:plan_id"`
	DCRTemplateID        uint    `gorm:"column:dcr_template_id"`
	DCREnabled           bool    `gorm:"column:dcr_request"`
	ProvisionImmediately bool    `gorm:"column:provision_immediately"`
	CatalogueID          int     `gorm:"catalogue_id"`
	ProviderID           uint    `gorm:"provider_id"`
}
