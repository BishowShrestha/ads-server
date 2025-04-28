package db

import (
	//"database/sql"
	//_ "github.com/lib/pq"
	"ad-server/internal/ads"
	"ad-server/internal/clicks"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgresDB(url string) *gorm.DB {
	//db, err := sql.Open("postgres", url)
	//if err != nil {
	//	log.Fatalf("failed to connect to database: %v", err)
	//}
	//
	//if err := db.Ping(); err != nil {
	//	log.Fatalf("failed to ping database: %v", err)
	//}
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&clicks.ClickEvent{}, &ads.Ad{})
	return db
}
