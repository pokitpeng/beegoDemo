package main

import (
	"beegoDemo/models"
	"beegoDemo/pkg/database"
	_ "beegoDemo/routers"

	"github.com/astaxie/beego"
)

func init() {
	database.InitDB(beego.AppConfig.String("dbDriver"))
	models.DB = database.GetDB(beego.AppConfig.String("dbDriver"))
	models.InitTales()
}

func main() {
	beego.Run()
}
