package config

type AppConfig struct {
	ServicePort string
	DB          DBConfig
	Superuser   Superuser
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
