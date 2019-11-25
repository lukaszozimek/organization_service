package invoice

import (
	"github.com/jinzhu/gorm"
)

type Organization struct {
	gorm.Model
	Name string `json:"name"`
}
