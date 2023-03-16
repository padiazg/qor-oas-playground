package user

import "github.com/jinzhu/gorm"

type Team struct {
	gorm.Model
	Name           string `gorm:"column:name;index"`
	Users          []User `gorm:"many2many:teams_users;"`
	OrganisationID uint   `gorm:"column:organisation_id"`
	// Organisation   Organisation
	// Default        bool `gorm:"column:is_default"`
}
