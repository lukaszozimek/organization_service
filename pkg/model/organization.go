package model

import (
	"github.com/jinzhu/gorm"
)

type Organization struct {
	gorm.Model
	Name              string             `json:"name"`
	UserOrganizations []OrganizationUser `gorm:"foreignkey:UserRefer"`
}
type OrganizationUser struct {
	gorm.Model
	Number    string `gorm:"unique_index;not null"`
	UserRefer uint   `gorm:"unique_index;not null"`
}
