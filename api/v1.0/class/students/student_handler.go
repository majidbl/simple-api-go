package students

import (
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	"fmt"
	"net/http"

	"github.com/majidzarephysics/university/model"
)

// create func will be use to create new class with specific StudentID
func create(c *gin.Context) {
	
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("sid")
	cid := c.Param("cid")
	studentclass := model.StudentClass{
		StudentID: sid,
		ClassID:   cid,
	}
	db.Create(&studentclass)
	c.JSON(http.StatusOK, studentclass)
	//c.JSON(http.StatusOK, studentJSON)
}

// read func return all classes that created by specific StudentID
func read(c *gin.Context) {
	// Get all records
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	studenSid := c.Param("Sid")
	var allStudentClass []model.StudentClass
	var allClass []model.Class
	var class model.Class
	db.Where(&model.StudentClass{StudentID: studenSid}).Find(&allStudentClass)
	for i := 0; i < len(allStudentClass); i++ {
		db.Where(&model.Class{CID: allStudentClass[i].ClassID}).Find(&class)
		allClass = append(allClass, class)
	}
	fmt.Println(allStudentClass[0])
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "result": strresult})
	c.JSON(http.StatusOK, allClass)

}

//readByID func return specific class with classid and StudentID
func readByID(c *gin.Context) {
	// return specific user by student ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("sid")
	cid := c.Param("cid")
	var studentclass model.StudentClass
	var class model.Class
	// Get first matched record
	db.Where(&model.StudentClass{StudentID: sid, ClassID: cid}).First(&studentclass)
	db.Where(&model.Class{CID: studentclass.ClassID}).First(&class)
	c.JSON(http.StatusOK, class)
}

// deleteByID func will be use to  Delete specific class
func deleteByID(c *gin.Context) {
	// return specific user by student ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	sid := c.Param("sid")
	cid := c.Param("cid")
	var student model.Student
	var class model.Class
	// Get first matched record
	db.Where(&model.Student{SID: sid}).First(&student)
	db.Model(&student).Where(&model.Class{CID: cid}).Association("Class").Find(&class)
	db.Model(&student).Association("Class").Find(&class)
	c.JSON(http.StatusOK, gin.H{"message": "specific class deleted succesfully", "status": http.StatusOK})

}
