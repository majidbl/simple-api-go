package teachers

import (
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	"fmt"
	"net/http"

	"github.com/majidzarephysics/university/model"
	"github.com/majidzarephysics/university/util"
)

// create new Teacher
func create(c *gin.Context) {
	// You also can use a struct
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	var teacherJSON model.Teacher
	c.BindJSON(&teacherJSON)
	pid, err := util.GetPersonalId("teacher")
	if err != nil {
		fmt.Println(err)
	}
	teacherJSON.TID = pid
	db.Create(&teacherJSON)
	fmt.Println(teacherJSON)
	c.JSON(http.StatusOK, teacherJSON)
	//c.JSON(http.StatusOK, studentJSON)
}

// read all records and return as JSON
func read(c *gin.Context) {
	// Get all records
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	var allTeacher []model.Teacher
	_ = db.Find(&allTeacher)
	fmt.Println(allTeacher)
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "result": strresult})
	c.JSON(http.StatusOK, allTeacher)

}

// readByID return specific Teacher
func readByID(c *gin.Context) {
	// return specific user by Teacher ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("id")
	var teacher model.Teacher
	// Get first matched record
	db.Where(&model.Student{SID: sid}).First(&teacher)
	c.JSON(http.StatusOK, teacher)
}

// deleteByID delet specific teacher
func deleteByID(c *gin.Context) {
	// return specific user by Teacher ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	tid := c.Param("id")
	var teacher model.Teacher
	// Get first matched record
	db.Where(&model.Teacher{TID: tid}).First(&teacher)
	db.Delete(&teacher)
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted succesfully", "status": http.StatusOK})

}

// update will be ues to update teacher
func update(c *gin.Context) {
	// return specific user by Teacher ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	tid := c.Param("id")
	var teacher model.Teacher
	var newTeacher model.Teacher
	c.BindJSON(&newTeacher)
	// Get first matched record
	db.Where(&model.Teacher{TID: tid}).First(&teacher)
	teacher.FirstName = newTeacher.FirstName
	teacher.LastName = newTeacher.LastName
	teacher.Password = newTeacher.Password
	db.Save(&teacher)
	c.JSON(http.StatusOK, teacher)

}
