package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `form:"email"`
	Password string
	Name     string `form:"name"`
	Role     string
	Teams    []Team `gorm:"many2many:teams_users;"`
}

func (user User) DisplayName() string {
	return user.Email
}
