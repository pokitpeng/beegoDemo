package routers

import (
	"beegoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article/:id", &controllers.ArticleController{}, "get:GetOne")
	beego.Router("/article/:id", &controllers.ArticleController{}, "delete:DeleteOne")
	beego.Router("/article", &controllers.ArticleController{}, "get:GetAll")
	beego.Router("/article", &controllers.ArticleController{}, "post:AddArticle")
	beego.Router("/article", &controllers.ArticleController{}, "put:EditArticle")

	beego.Router("/upload", &controllers.UploadController{}, "get:Upload")
	beego.Router("/upload/file", &controllers.UploadController{}, "post:UploadFile")
	beego.Router("/upload/files", &controllers.UploadController{}, "post:UploadFiles")

}
