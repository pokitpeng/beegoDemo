package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"beegoDemo/models"
)

type StudentController struct {
	beego.Controller
}

// http://127.0.0.1:8080/student
func (c *StudentController) Get() {
	// 	查询学生信息
	// var students []models.Student
	// models.DB.Find(&students)
	// c.Data["json"] = students
	// c.ServeJSON()

	// 	查询课程信息
	// var lessons []models.Lesson
	// models.DB.Find(&lessons)
	// c.Data["json"] = lessons
	// c.ServeJSON()

	// 	查询学生信息时，获取对应的课程信息
	var students []models.Student
	err := models.DB.Preload("Lessons").Find(&students).Error
	if err != nil {
		logs.Error(err)
	}
	c.Data["json"] = students
	c.ServeJSON()
}
