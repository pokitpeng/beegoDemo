package global

var (
	Sqlite3  = "sqlite3"
	Mysql    = "mysql"
	Postgres = "postgres"

	// DB
	MysqlHost           = "127.0.0.1"
	MysqlPort           = 3306
	MysqlUser           = "root"
	MysqlPassword       = "123456"
	MysqlDbname         = "goexample"
	MysqlDBMaxIdleConns = 10
	MysqlDBMaxOpenConns = 10

	PostgresHost           = "192.168.50.97"
	PostgresPort           = 30238
	PostgresUser           = "datatom"
	PostgresPassword       = "db_password"
	PostgresDefaultDbname  = "postgres"
	PostgresDbname         = "goexample"
	PostgresDBMaxIdleConns = 10
	PostgresDBMaxOpenConns = 10

	Sqlite3Path   = "./.tmp"
	Sqlite3Dbname = "demo.db"
)
