package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	RoleName   string `json:"roleName"`
	CreateBy   string `json:"createBy"`
	ModifiedBy string `json:"modifiedBy"`
}
