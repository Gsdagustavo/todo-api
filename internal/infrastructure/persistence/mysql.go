package persistence

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func OpenDatabase() (*sql.DB, error) {
	databaseConfig := NewDatabaseConfig()

	mysqlCfg := mysql.Config{
		User:   databaseConfig.User,
		Passwd: databaseConfig.Password,
		Net:    databaseConfig.Net,
		Addr:   databaseConfig.Address(),
		DBName: databaseConfig.DBName,
	}

	log.Printf("Opening database %s", databaseConfig.DBName)

	db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db, nil
}
