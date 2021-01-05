package controllers

import "github.com/astaxie/beego"

type ArticleController struct {
	beego.Controller // 继承
}

type ResponseStruct struct {
	Code int
	Msg  string
	Data string
}

// curl 127.0.0.1:8080/article/12
func (c *ArticleController) GetOne() {
	id := c.GetString(":id")
	c.Data["json"] = &ResponseStruct{Code: 200, Msg: "获取成功", Data: "id=" + id}
	c.ServeJSON()
}

// curl -X DELETE 127.0.0.1:8080/article/12
func (c *ArticleController) DeleteOne() {
	id := c.GetString(":id")
	c.Data["json"] = &ResponseStruct{Code: 200, Msg: "删除成功", Data: "id=" + id}
	c.ServeJSON()
}

// curl 127.0.0.1:8080/article
func (c *ArticleController) GetAll() {
	c.Data["title"] = "文章列表"
	c.TplName = "article/index.html"
}

// curl -X POST -d "" 127.0.0.1:8080/article
func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加文章")
}

// curl -X PUT -d "" 127.0.0.1:8080/article
func (c *ArticleController) EditArticle() {
	c.Ctx.WriteString("编辑文章")
}
