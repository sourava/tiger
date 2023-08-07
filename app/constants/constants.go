package constants

const (
	ServicePort                      string = "SERVICE_PORT"
	MysqlHost                        string = "MYSQL_HOST"
	MysqlPort                        string = "MYSQL_PORT"
	MysqlDb                          string = "MYSQL_DATABASE"
	MysqlUser                        string = "MYSQL_USER"
	MysqlPassword                    string = "MYSQL_PASSWORD"
	SuperuserUsername                string = "SUPERUSER_USERNAME"
	SuperuserEmail                   string = "SUPERUSER_EMAIL"
	SuperuserHashedPassword          string = "SUPERUSER_HASHED_PASSWORD"
	SendgridApiKey                   string = "SENDGRID_API_KEY"
	SendgridSenderEmail              string = "SENDGRID_SENDER_EMAIL"
	SendgridSenderName               string = "SENDGRID_SENDER_NAME"
	TigerSightingNotificationRequest string = "tigerSightingNotificationSubscriber:========TigerSightingNotificationRequest========:"
	DbConnectionString               string = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
)
