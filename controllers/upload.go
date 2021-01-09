package controllers

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UploadController struct {
	beego.Controller // 继承
}

// beego上传文件

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

func (c *UploadController) UploadFiles() {
	title := c.GetString("title")
	logs.Info(title)
	f1, h1, err := c.GetFile("file1")
	if err != nil {
		logs.Info("get file1 err ", err)
	}
	f2, h2, err := c.GetFile("file2")
	if err != nil {
		logs.Info("get file2 err ", err)
	}
	defer f1.Close()
	defer f2.Close()
	savePath := "static/upload/"
	if !Exists(savePath) {
		os.MkdirAll(savePath, 0644)
	}
	c.SaveToFile("file1", savePath+h1.Filename)
	c.SaveToFile("file2", savePath+h2.Filename)
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
