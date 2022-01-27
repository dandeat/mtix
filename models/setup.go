package models

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {

	DB, err = gorm.Open("mysql","root:@/db_mtix?charset=utf8&parseTime=true")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&Tiket{},&Cust{},&Movie{},&Payment{},&Schedule{})
}
