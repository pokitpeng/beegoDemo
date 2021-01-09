package database

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego/logs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"beegoDemo/pkg/utils"
)

type Sqlite3 struct {
	Path   string
	Dbname string
}

func (s *Sqlite3) InitDB() *gorm.DB {
	var err error
	allPath := path.Join(s.Path, s.Dbname)
	// 1.check database exists?
	if !utils.Exists(allPath) {
		logs.Info("database [%v] not exist, start create", s.Dbname)
		if err = os.MkdirAll(s.Path, 755); err != nil {
			panic("create database path " + s.Path + "err")
		}
		if _, err = os.Create(allPath); err != nil {
			panic("create database " + s.Dbname + " err")
		}
	} else {
		logs.Info("database [%v] already exist", s.Dbname)
	}

	// 2.connect to db
	args := fmt.Sprintf(allPath)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		})
	Sqlite3DB, err = gorm.Open(sqlite.Open(args), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("faild to connect database,err:" + err.Error())
	}
	return Sqlite3DB
}
