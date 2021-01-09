package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller // 继承
}

// gorm CRUD

/*插入测试数据
INSERT INTO user (username,age,email)VALUES('tom',18,'tom@qq.com');
INSERT INTO user (username,age,email)VALUES('tom2',18,'tom2@qq.com');
INSERT INTO user (username,age,email)VALUES('tom3',19,'tom3@qq.com');
INSERT INTO user (username,age,email)VALUES('tom4',19,'tom4@qq.com');
INSERT INTO user (username,age,email)VALUES('tom5',19,'tom5@qq.com');
INSERT INTO user (username,age,email)VALUES('tom6',20,'tom6@qq.com');
*/

// http://127.0.0.1:8080/user
func (c *UserController) Get() {
	// // 1.获取所有数据
	// var users []models.User
	// models.DB.Find(&users)
	// c.Data["json"] = users
	// c.ServeJSON()

	// // 	2.获取一个数据
	// var user models.User
	// models.DB.Where(&models.User{Username: "tom"}).Find(&user)
	// c.Data["json"] = user
	// c.ServeJSON()

	// // 	3.获取多个数据
	// var users []models.User
	// models.DB.Where(&models.User{Age: 18}).Find(&users)
	// c.Data["json"] = users
	// c.ServeJSON()

	// // 	4.获取不存在的多个数据
	// var users []models.User
	// models.DB.Where(&models.User{Age: 0}).Find(&users)
	// c.Data["json"] = users
	// c.ServeJSON()

	// // 	5.获取不存在的一个数据
	// var user models.User
	// models.DB.Where(&models.User{Username: "pokit"}).Find(&user)
	// c.Data["json"] = user
	// c.ServeJSON() // 返回一个属性全为零值的对象

	// // 	6.插入一条数据, Create可重复插入，不会报错
	// var tony = &models.User{Username: "tony", Age: 20, Email: "tony@qq.com"}
	// err := models.DB.Create(&tony).Error
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = err
	// 	c.ServeJSON()
	// 	return
	// }
	// c.Data["json"] = tony
	// c.ServeJSON()

}
