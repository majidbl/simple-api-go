package model

import "github.com/jinzhu/gorm"

// Teacher struct as model for map in db
type Teacher struct {
	gorm.Model
	FirstName string  `gorm:"column:FirstName;type:VARCHAR(160);NOT NULL;" json:"firstname"`
	LastName  string  `gorm:"column:LastName;type:VARCHAR(160);NOT NULL;" json:"lastname"`
	Password  string  `gorm:"column:Password;type:VARCHAR(160);NOT NULL;" json:"password"`
	TID       string  `gorm:"column:TID;type:VARCHAR(160);NOT NULL;" json:"tid"`
	Classes   []Class `gorm:"foreignKey:TeacherID"`
}
