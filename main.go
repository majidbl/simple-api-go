package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/majidzarephysics/university/api"
	"github.com/majidzarephysics/university/model"
	"github.com/majidzarephysics/university/util"
)

func main() {
	pid, err := util.GetPersonalId("teacher")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pid)
	cid, err := util.GetClassId("IT", "Saturday", "10-12")
	fmt.Println(cid)
	db, err := model.GetDB()
	if err != nil {
		log.Fatal("database initial was failed")
	}

	db.AutoMigrate(&model.Teacher{}, &model.Student{}, &model.Class{}, &model.StudentClass{})
	router := gin.Default()
	api.ApplyRoutes(router)
	router.GET("/ping", index)
	router.Run(":4321")
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "well come to university Api",
	})
}
