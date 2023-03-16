package product

import (
	"github.com/jinzhu/gorm"
)

type Catalogue struct {
	gorm.Model
	Name             string    `gorm:"column:name"`
	VisibilityStatus string    `gorm:"column:visibility_status"`
	Products         []Product `gorm:"many2many:catalogue_products;"`
	Plans            []Plan    `gorm:"many2many:catalogue_plans;"`
}
