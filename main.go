package main

import (
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//Initialize a struct of type Config, inject user, pass, name from .env file
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPass,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	//Init the DB connection given the configuration
	sqlStorage := InitMYSQLStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	//Init the actual schema and tables
	store := NewStore(db)
	//Initialize and start the API
	api := newAPIServer(":3000", store)
	api.Serve() // init router & services
}
