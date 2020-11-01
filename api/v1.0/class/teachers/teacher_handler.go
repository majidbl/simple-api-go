package teachers

import (
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	"fmt"
	"net/http"

	"github.com/majidzarephysics/university/model"
	"github.com/majidzarephysics/university/util"
)

// create func will be use to create new class with specific TeacherID
func create(c *gin.Context) {
	// You also can use a struct
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	teacherID := c.Param("tid")
	var teacher model.Teacher
	var classJSON model.Class
	db.Where(&model.Teacher{TID: teacherID}).First(&teacher)
	fmt.Println(teacher)
	//Get class as json that posted with user
	c.BindJSON(&classJSON)
	cid, err := util.GetClassId(classJSON.Name, classJSON.Day, classJSON.Time)
	if err != nil {
		fmt.Println(err)
	}

	classJSON.CID = cid
	classJSON.TeacherID = teacher.ID
	fmt.Println(classJSON)
	//db.Model(&teacher).Association("Class").Append([]model.Class{classJSON})
	db.Create(&classJSON)
	c.JSON(http.StatusOK, classJSON)
	//c.JSON(http.StatusOK, studentJSON)
}

// read func return all classes that created by specific teacherID
func read(c *gin.Context) {
	// Get all records
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	teacherID := c.Param("tid")
	var allClass []model.Class
	var teacher model.Teacher
	db.Where(&model.Teacher{TID: teacherID}).First(&teacher)
	fmt.Println(teacher)
	db.Where(&model.Class{TeacherID: teacher.ID}).Find(&allClass)
	fmt.Println(allClass)
	//c.JSON(http.StatusOK, gin.H{"message": "hey", "result": strresult})
	c.JSON(http.StatusOK, allClass)

}

//readByID func return specific class with classid and TeacherID
func readByID(c *gin.Context) {
	// return specific user by Teacher ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	tid := c.Param("tid")
	cid := c.Param("cid")
	var teacher model.Teacher
	var class model.Class
	// Get first matched record
	db.Where(&model.Teacher{TID: tid}).First(&teacher)
	db.Where(&model.Class{CID: cid, TeacherID: teacher.ID}).Find(&class)
	c.JSON(http.StatusOK, class)
}

// deleteByID func will be use to  Delete specific class
func deleteByID(c *gin.Context) {
	// return specific user by Teacher ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	tid := c.Param("tid")
	cid := c.Param("cid")
	var teacher model.Teacher
	var class model.Class
	// Get first matched record
	db.Where(&model.Teacher{TID: tid}).First(&teacher)
	db.Where(&model.Class{CID: cid, TeacherID: teacher.ID}).Delete(&class)
	c.JSON(http.StatusOK, gin.H{"message": "specific class deleted succesfully", "status": http.StatusOK})

}

// update func will be called when teacher want update specific class
func update(c *gin.Context) {
	// return specific user by Student ID
	db, errdb := model.GetDB()
	if errdb != nil {
		fmt.Println(errdb)
	}
	tid := c.Param("tid")
	var teacher model.Teacher
	var newClass model.Class
	var class model.Class
	c.BindJSON(&newClass)
	// Get first matched record
	db.Where(&model.Teacher{TID: tid}).First(&teacher)
	db.Where(&model.Class{CID: newClass.CID, TeacherID: teacher.ID}).First(&class)
	class.Time = newClass.Time
	class.Day = newClass.Day
	class.Name = newClass.Name
	db.Save(&class)
	c.JSON(http.StatusOK, class)

}
