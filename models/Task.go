package models

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model //Aca le indicamos a gorm que mantenga las migraciones

	Title 		string `gorm:"type:varchar(100);not null; unique_index"`
	Description string
	Done 		bool `gorm:"default:false"`
	UserID		uint
}