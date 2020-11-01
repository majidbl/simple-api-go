package model

import (
	"github.com/jinzhu/gorm"
	//"github.com/lib/pq"
)

// Student struct as model for map in db
type Student struct {
	gorm.Model
	FirstName string `gorm:"column:FirstName;type:VARCHAR(160);NOT NULL" json:"firstname"`
	LastName  string `gorm:"column:LastName;type:VARCHAR(160);NOT NULL;" json:"lastname"`
	Password  string `gorm:"column:Password;type:VARCHAR(160);NOT NULL;" json:"password"`
	SID       string `gorm:"column:SID;type:VARCHAR(160);NOT NULL;" json:"sid"`
}

/*curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"firstname":"testname","lastname":"testfamily","password":"xyz"}' \
  http://localhost:4321/api/v1.0/student*/
