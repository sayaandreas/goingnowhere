package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	HOST = "database"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	db := Database{}
	dsn := viper.GetString("dsn")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
