package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "numbers"
)

type Number struct {
	gorm.Model
	Number uint16
	UUID   string `gorm:"uuid"`
}

// DB set up
func SetupDB(host string, user string, password string, dbName string, port string) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)
	// dsn := "host=localhost user=postgres password=mysecretpassword dbname=numbers port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Number{})
	//db.Create(&Number{Number: 900, UUID: "03691ee4-a5d5-4e91-95ce-331061695949"})

	return db
}
