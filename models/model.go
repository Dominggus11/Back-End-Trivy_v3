package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint
	ProjectName string `json:"projectname"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Dockerfiles struct {
	gorm.Model
	ID        uint
	Pathfile  string `json:"pathfile"`
	PathJson  string `json:"pathjson"`
	ProjectID int
	Project   Project `gorm:"foreignKey:ProjectID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
