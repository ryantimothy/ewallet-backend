package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dev"
	db_name  = "wallet_db_ryan"
)

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, db_name, port)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,
	// 		LogLevel:                  logger.Info,
	// 		IgnoreRecordNotFoundError: true,
	// 		Colorful:                  false,
	// 	})

	DBConnect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//DBConnect.AutoMigrate(entity.User{}, entity.Wallet{}, entity.Transaction{})

	if err != nil {
		log.Fatal(err)
	}
	return DBConnect
}
