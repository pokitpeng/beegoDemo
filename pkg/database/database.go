package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	ormlogger "gorm.io/gorm/logger"
)

// DB define db object
var DB *gorm.DB

// init init DB
func InitDB() *gorm.DB {
	var err error

	// 1.check database exists?
	conninfo := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%s sslmode=disable", beego.AppConfig.String("pgDBUsername"),
		beego.AppConfig.String("pgDBPassword"), beego.AppConfig.String("pgDBHost"), beego.AppConfig.String("pgDBPort"), beego.AppConfig.String("pgDefaultDatabase"))
	db, err := sql.Open("postgres", conninfo)
	defer db.Close()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}

	execArgs := fmt.Sprintf("CREATE DATABASE %v", beego.AppConfig.String("pgDBDataBase"))
	_, err = db.Exec(execArgs)
	if err != nil {
		// database exists
		logs.Info("create database err:%v", err)
	} else {
		logs.Info("create database [%v] success", beego.AppConfig.String("pgDBDataBase"))
	}

	// 2.connect to db
	args := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable password=%v TimeZone=Asia/Shanghai", beego.AppConfig.String("pgDBHost"),
		beego.AppConfig.String("pgDBPort"), beego.AppConfig.String("pgDBUsername"), beego.AppConfig.String("pgDBDataBase"), beego.AppConfig.String("pgDBPassword"))
	newLogger := ormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlogger.Config{
			SlowThreshold: time.Second,      // 慢 SQL 阈值
			LogLevel:      ormlogger.Silent, // Log level
			Colorful:      false,            // 禁用彩色打印
		})

	DB, err = gorm.Open(postgres.Open(args), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("faild to connect database,err:" + err.Error())
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("faild to get sqlDB,err:" + err.Error())
	}
	pgDBMaxIdleConns, err := beego.AppConfig.Int("pgDBMaxIdleConns")
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	pgDBMaxOpenConns, err := beego.AppConfig.Int("pgDBMaxOpenConns")
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	pgDBConnMaxLifetime, err := beego.AppConfig.Int("pgDBConnMaxLifetime")
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	sqlDB.SetMaxIdleConns(pgDBMaxIdleConns) // max idle connection
	sqlDB.SetMaxOpenConns(pgDBMaxOpenConns) // max open connection,unlimited by default
	sqlDB.SetConnMaxLifetime(time.Duration(pgDBConnMaxLifetime) * time.Second)
	return DB
}

// GetDB get the initialized database object
func GetDB() *gorm.DB {
	return DB
}
