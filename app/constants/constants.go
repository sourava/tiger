package constants

const (
	ServicePort             string = "SERVICE_PORT"
	MysqlHost               string = "MYSQL_HOST"
	MysqlPort               string = "MYSQL_PORT"
	MysqlDb                 string = "MYSQL_DATABASE"
	MysqlUser               string = "MYSQL_USER"
	MysqlPassword           string = "MYSQL_PASSWORD"
	SuperuserUsername       string = "SUPERUSER_USERNAME"
	SuperuserEmail          string = "SUPERUSER_EMAIL"
	SuperuserHashedPassword string = "SUPERUSER_HASHED_PASSWORD"
	DbConnectionString      string = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
)
