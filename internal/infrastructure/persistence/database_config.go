package persistence

import "os"

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	DBPort   string
	DBName   string
	Net      string
}

func (conf *DatabaseConfig) Address() string {
	return conf.Host + ":" + conf.DBPort
}

func NewDatabaseConfig() *DatabaseConfig {
	var (
		user   = os.Getenv("DB_USER")
		passwd = os.Getenv("DB_PASSWORD")
		host   = os.Getenv("DB_HOST")
		dbName = os.Getenv("DB_NAME")
		dbPort = os.Getenv("DB_PORT")
		net    = "tcp"
	)

	return &DatabaseConfig{
		User:     user,
		Password: passwd,
		Host:     host,
		DBPort:   dbPort,
		DBName:   dbName,
		Net:      net,
	}
}
