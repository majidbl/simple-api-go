package model

import "github.com/jinzhu/gorm"

// Class struct as model for map in db
type Class struct {
	gorm.Model
	Name      string `gorm:"column:Name;type:VARCHAR(160);NOT NULL;" json:"name"`
	Time      string `gorm:"column:Time;type:VARCHAR(160);NOT NULL;" json:"time"`
	Day       string `gorm:"column:Day;type:VARCHAR(160);NOT NULL;" json:"day"`
	Capacity  int    `gorm:"column:Capacity;type:integer;NOT NULL;" json:"capacity"`
	CID       string `gorm:"column:CID;type:VARCHAR(160);NOT NULL;" json:"cid"`
	TeacherID uint   `gorm:"column:TeacherID;type:VARCHAR(160);NOT NULL;" json:"teacherid"`
}
