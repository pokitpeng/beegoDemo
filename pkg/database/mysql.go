package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Mysql struct {
	Host           string
	Port           int
	User           string
	Dbname         string
	Password       string
	DBMaxIdleConns int
	DBMaxOpenConns int
}

func (m *Mysql) InitDB() *gorm.DB {
	var err error

	// 1.check database exists?
	conninfo := fmt.Sprintf("%v:%v@tcp(%v:%v)/?charset=utf8&autocommit=true", m.User, m.Password, m.Host, m.Port)
	db, err := sql.Open(beego.AppConfig.String("dbDriver"), conninfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("create database " + m.Dbname)
	if err != nil {
		// database exists
		logs.Info("create database err:%v", err)
	} else {
		logs.Info("create database [%v] success", m.Dbname)
	}

	// 2.connect to db
	args := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", m.User, m.Password, m.Dbname)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		})
	MysqlDB, err = gorm.Open(mysql.Open(args), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	sqlDB, err := MysqlDB.DB()
	if err != nil {
		panic("failed to get sqlDB,err:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(m.DBMaxIdleConns) // max idle connection
	sqlDB.SetMaxOpenConns(m.DBMaxOpenConns) // max open connection,unlimited by default
	return MysqlDB
}
