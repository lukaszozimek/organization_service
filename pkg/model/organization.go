package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Organization struct {
	gorm.Model
	Name             string    `json:"name"`
	Active           bool      `json:"active"`
	CreationDate     time.Time `json:"creationDate"`
	ModificationDate time.Time `json:"modificationDate"`
}
