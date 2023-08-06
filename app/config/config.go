package config

type AppConfig struct {
	ServicePort string
	DB          DBConfig
	Superuser   Superuser
	Sendgrid    Sendgrid
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Superuser struct {
	Username       string
	Email          string
	HashedPassword string
}

type Sendgrid struct {
	ApiKey      string
	SenderEmail string
	SenderName  string
}
