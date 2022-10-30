package model

import (
	"github.com/raaaaaaaay86/gorm-custom-datatype/model/datatype"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name  string
	Owner datatype.Owner `gorm:"type:VARCHAR(255)"`
}
