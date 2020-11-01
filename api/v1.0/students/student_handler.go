package students

import (
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	"fmt"
	"net/http"

	"github.com/majidzarephysics/university/model"
	"github.com/majidzarephysics/university/util"
)

// create new student
func create(c *gin.Context) {
	// You also can use a struct
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	var studentJSON model.Student
	c.BindJSON(&studentJSON)
	sid, err := util.GetPersonalId("student")
	if err != nil {
		fmt.Println(err)
	}
	studentJSON.SID = sid
	db.Create(&studentJSON)
	fmt.Println(studentJSON)
	c.JSON(http.StatusOK, gin.H{"message": "hey", "result": studentJSON})
	//c.JSON(http.StatusOK, studentJSON)
}

// read return all Student
func read(c *gin.Context) {
	// Get all records
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	var allStudent []model.Student
	_ = db.Find(&allStudent)
	fmt.Println(allStudent)
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "result": strresult})
	c.JSON(http.StatusOK, allStudent)

}

// readByID return specific Student
func readByID(c *gin.Context) {
	// return specific user by Student ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("id")
	var student model.Student
	// Get first matched record
	db.Where(&model.Student{SID: sid}).First(&student)
	c.JSON(http.StatusOK, student)
}

// deleteByID Will be use for delete specific student
func deleteByID(c *gin.Context) {
	// delete specific user by Student ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("id")
	var student model.Student
	// Get first matched record
	db.Where(&model.Student{SID: sid}).First(&student)
	db.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted succesfully", "status": http.StatusOK})

}

// update will be ues to update Student
func update(c *gin.Context) {
	// return specific user by Student ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("id")
	var student model.Student
	var newStudent model.Student
	c.BindJSON(&newStudent)
	// Get first matched record
	db.Where(&model.Student{SID: sid}).First(&student)
	student.FirstName = newStudent.FirstName
	student.LastName = newStudent.LastName
	student.Password = newStudent.Password
	db.Save(&student)
	c.JSON(http.StatusOK, student)

}
