package model

import (
	"github.com/jinzhu/gorm"
)

// StudentClass struct as model for map in db
type StudentClass struct {
	gorm.Model
	StudentID string `gorm:"column:StudentID;type:VARCHAR(160);NOT NULL" json:"studentid"`
	ClassID   string `gorm:"column:ClassID;type:VARCHAR(160);NOT NULL;" json:"classid"`
}
