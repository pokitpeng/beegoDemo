package controllers

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UploadController struct {
	beego.Controller // 继承
}

// http://127.0.0.1:8080/upload
func (c *UploadController) Upload() {
	c.TplName = "upload/index.html"
}

func (c *UploadController) UploadFile() {
	title := c.GetString("title")
	logs.Info(title)
	f, h, err := c.GetFile("file")
	if err != nil {
		logs.Info("get file err ", err)
	}
	defer f.Close()
	savePath := "static/upload/"
	if !Exists(savePath) {
		os.MkdirAll(savePath, 0644)
	}
	c.SaveToFile("file", savePath+h.Filename)
	c.Ctx.WriteString("上传成功")
}

func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
