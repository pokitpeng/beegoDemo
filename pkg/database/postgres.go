package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"beegoDemo/global"
)

type Postgres struct {
	Host           string
	Port           int
	User           string
	Dbname         string
	Password       string
	DBMaxIdleConns int
	DBMaxOpenConns int
}

func (p *Postgres) InitDB() *gorm.DB {
	var err error

	// 1.check database exists?
	conninfo := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=disable", p.User,
		p.Password, p.Host, p.Port, global.PostgresDefaultDbname)
	db, err := sql.Open(global.Postgres, conninfo)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("create database " + p.Dbname)
	if err != nil {
		// database exists
		logs.Info("create database err:%v", err)
	} else {
		logs.Info("create database [%v] success", p.Dbname)
	}

	// 2.connect to db
	args := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v", p.Host,
		p.Port, p.User, p.Dbname, p.Password)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		})
	PostgresDB, err = gorm.Open(postgres.Open(args), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("faild to connect database,err:" + err.Error())
	}
	sqlDB, err := PostgresDB.DB()
	if err != nil {
		panic("faild to get sqlDB,err:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(p.DBMaxIdleConns) // max idle connection
	sqlDB.SetMaxOpenConns(p.DBMaxOpenConns) // max open connection,unlimited by default
	return PostgresDB
}
