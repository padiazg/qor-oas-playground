package product

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	DisplayName string `gorm:"column:display_name"`
	Path        string `gorm:"column:path"`
	ReferenceID string `gorm:"column:reference_id"`
	Description string `gorm:"column:description"`
}
