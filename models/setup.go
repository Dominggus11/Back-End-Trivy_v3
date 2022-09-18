package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgres://fknulqmibfezlc:1c3d06246bc3701b199929a2b4deb1604cf3a1383addb53c52a32b5879506f63@ec2-107-23-76-12.compute-1.amazonaws.com:5432/d8rp6ktl3svj5l"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connect")
	}

	database.AutoMigrate(&Dockerfiles{}, &Projects{})

	DB = database
}
