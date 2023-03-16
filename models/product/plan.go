package product

import "github.com/jinzhu/gorm"

type Plan struct {
	gorm.Model
	Name        string      `gorm:"column:name"`
	DisplayName string      `gorm:"column:display_name"`
	ReferenceID string      `gorm:"column:reference_id"`
	Description string      `gorm:"column:description"`
	AuthType    string      `gorm:"column:auth_type"`
	JWTScope    string      `gorm:"column:jwt_scope"`
	Catalogues  []Catalogue `gorm:"many2many:catalogue_plans;"`
	// AccessRequests            []AccessRequest
	Rate                      float64 `gorm:"column:rate"`
	Per                       float64 `gorm:"column:per"`
	QuotaMax                  int64   `gorm:"column:quota_max"`
	QuotaRenewalRate          int64   `gorm:"column:quota_renewal_rate"`
	AutoApproveAccessRequests bool    `gorm:"column:auto_approve_access_requests"`
	ProviderID                uint    `gorm:"column:provider_id"`
}
