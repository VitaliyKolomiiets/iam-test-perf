package models

import (
	"iam-test-perf/cmd/servce/krn"
)

// User store created users in Postgres DB.
type User struct {
	IDModel `gorm:"embedded"`

	KRN      *krn.KRN `gorm:"primarykey;column:krn;type:string;size:256"                  json:"krn,omitempty"`
	TenantID string   `gorm:"not null;column:tenant_id;type:string;size:256;default:null" json:"tenantID"`

	AuthID       string `gorm:"column:auth_id;type:string;size:256"                    json:"authID"`
	Email        string `gorm:"column:email;type:string;size:256"                             json:"email,omitempty" validate:"omitempty,email"`
	Enabled      bool   `gorm:"column:enabled;type:bool;default:false"                        json:"enabled"`
	FirstName    string `gorm:"column:first_name;type:string;size:256"                        json:"firstName"`
	LastName     string `gorm:"column:last_name;type:string;size:256"                         json:"lastName"`
	Username     string `gorm:"not null;column:username;type:string;size:256;default:null"    json:"username"        validate:"omitempty,name"`
	Path         string `gorm:"column:path;type:string;size:256"                              json:"path"`
	DefaultGroup string `gorm:"column:default_group;type:string;size:256"                     json:"defaultGroup,omitempty"`
	Status       string `gorm:"column:status;type:string;size:256"                            json:"status,omitempty"`

	Password        string `gorm:"column:password;type:string;size:256" json:"password,omitempty" validate:"required"`
	ConfirmPassWord string `gorm:"-" json:"confirmPassword,omitempty"                             validate:"required"`

	IdpID       string `gorm:"column:idp_id;type:string;size:256"        json:"IdpID,omitempty"`
	IdpProvider string `gorm:"column:idp_provider;type:string;size:256"  json:"IdpProvider,omitempty"`

	SendVerifyEmail bool   `gorm:"-" json:"sendVerifyEmail,omitempty"`
	RedirectURI     string `gorm:"-" json:"redirectUri,omitempty" validate:"omitempty,url"`
}
