package persistence

import "os"

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	Net      string
}

func (conf *DatabaseConfig) Address() string {
	return conf.Host + ":" + conf.Port
}

func NewDatabaseConfig() *DatabaseConfig {
	var (
		user       = os.Getenv("DB_USER")
		passwd     = os.Getenv("DB_PASSWORD")
		host       = os.Getenv("DB_HOST")
		dbName     = os.Getenv("DB_NAME")
		serverPort = os.Getenv("SERVER_PORT")
		net        = "tcp"
	)

	return &DatabaseConfig{
		User:     user,
		Password: passwd,
		Host:     host,
		Port:     serverPort,
		DBName:   dbName,
		Net:      net,
	}
}
