package config

import (
	"Find-Backend/features/user"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke Database, error: ", err)
	}


	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatal("Gagal untuk migrasi database, error: ", err)
	}	
}