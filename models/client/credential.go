package client

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Credential struct {
	gorm.Model
	Credential                 string `gorm:"column:credential"`
	CredentialHash             string `gorm:"column:credential_hash"`
	OAuthClientID              string `gorm:"column:o_auth_client_id"`
	OAuthClientSecret          string `gorm:"column:o_auth_client_secret"`
	RedirectURI                string `gorm:"column:redirect_uris"`
	TokenEndpoints             string `gorm:"column:token_endpoints"`
	ResponseType               string `gorm:"column:response_type"`
	GrantType                  string `gorm:"column:grant_type"`
	Scope                      string `gorm:"column:scope"`
	DCRRegistrationAccessToken string `gorm:"column:registration_access_token;type:text"`
	DCRRegistrationClientURI   string `gorm:"column:registration_client_uri;type:text"`
	DCRResponse                string `gorm:"type:text;column:dcr_response"`
	Expires                    time.Time
	AccessRequest              AccessRequest
	ClientID                   uint `gorm:"column:client_id"`
	AccessRequestID            uint `gorm:"column:access_request_id"`
}
