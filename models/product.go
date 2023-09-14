package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(50);" json:"identity"`
	Key      string `gorm:"column:key; type:varchar(50);" json:"key"`
	Name     string `gorm:"column:name; type:varchar(50);" json:"name"`
	Desc     string `gorm:"column:desc; type:varchar(50);" json:"desc"`
}

func (table Product) TableName() string {
	return "product"
}
