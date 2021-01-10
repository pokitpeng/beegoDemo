package controllers

import (
	"github.com/astaxie/beego"

	"beegoDemo/models"
)

type BookController struct {
	beego.Controller
}

// http://127.0.0.1:8080/book
func (c *BookController) Get() {
	// 关联查询 一对一
	// var books []models.Book
	// models.DB.Preload("Author").Find(&books)
	// c.Data["json"] = books
	// c.ServeJSON()

	// 	关联查询 一对多
	var authors []models.Author
	models.DB.Preload("Books").Find(&authors)
	c.Data["json"] = authors
	c.ServeJSON()

	// 	关联查询 一对多 指定author条件
	// var authors []models.Author
	// models.DB.Preload("Books").Where(&models.Author{Id: 2}).Find(&authors)
	// c.Data["json"] = authors
	// c.ServeJSON()

	// 	关联查询 一对多 指定book条件
	// var authors []models.Author
	// models.DB.Preload("Books", "id=8").Where(&models.Author{Id: 2}).Find(&authors)
	// c.Data["json"] = authors
	// c.ServeJSON()
}
