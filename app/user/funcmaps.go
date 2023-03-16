package user

import (
	"net/http"

	qa "github.com/TykTechnologies/portal-admin"
	um "github.com/padiazg/qor-oas-playground/models/user"
)

func CurrentUser(r *http.Request, Auth *qa.Auth) *um.User {
	// if iUser := Auth.GetCurrentUser(r); iUser != nil {
	// 	return iUser.(*um.User)
	// }
	return &um.User{
		Name:  "John Doe",
		Email: "jhon.d@tyk.io",
	}
}
