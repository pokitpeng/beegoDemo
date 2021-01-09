package database

import (
	"strings"

	"gorm.io/gorm"

	"beegoDemo/global"
)

// DB global db object
var (
	PostgresDB *gorm.DB
	MysqlDB    *gorm.DB
	Sqlite3DB  *gorm.DB
)

type InitDBAPI interface {
	InitDB() *gorm.DB
}

func InitDB(name string) {
	switch strings.ToLower(name) {
	case global.Mysql:
		m := &Mysql{
			Host:           global.MysqlHost,
			Port:           global.MysqlPort,
			User:           global.MysqlUser,
			Dbname:         global.MysqlDbname,
			Password:       global.MysqlPassword,
			DBMaxIdleConns: global.MysqlDBMaxIdleConns,
			DBMaxOpenConns: global.MysqlDBMaxOpenConns,
		}
		m.InitDB()
	case global.Postgres:
		p := &Postgres{
			Host:           global.PostgresHost,
			Port:           global.PostgresPort,
			User:           global.PostgresUser,
			Dbname:         global.PostgresDbname,
			Password:       global.PostgresPassword,
			DBMaxIdleConns: global.PostgresDBMaxIdleConns,
			DBMaxOpenConns: global.PostgresDBMaxOpenConns,
		}
		p.InitDB()
	case global.Sqlite3:
		s := &Sqlite3{
			Path:   global.Sqlite3Path,
			Dbname: global.Sqlite3Dbname,
		}
		s.InitDB()
	}
}

// GetDB get the initialized database object
func GetDB(name string) *gorm.DB {
	switch strings.ToLower(name) {
	case global.Mysql:
		return MysqlDB
	case global.Postgres:
		return PostgresDB
	case global.Sqlite3:
		return Sqlite3DB
	default:
		return Sqlite3DB
	}
}
