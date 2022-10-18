package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Id             int `gorm:"primaryKey"`
	Username       string
	Email          string
	HashedPassword string
}

type Chat struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	UserId    int
	User      User `gorm:"foreignKey:UserId`
	Content   string
	ChannelId int
}

func Init() *gorm.DB {
	dsn := "host=localhost user=microservice password=password dbname=microservice port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(User{}, Chat{})

	return db
}



