package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fariedrisky/go-restful-mysql/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Could not ping database: ", err)
	}

	DB = db
	log.Println("Connected to database")
}
