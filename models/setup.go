package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/tugas_rest_api"))
	if  err != nil {
		panic("Gagal Menyambung")
	}

	database.AutoMigrate(&Users{})

	DB = database
}