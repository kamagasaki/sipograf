package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := "mariadb:7F37RDMdizgkWC75q5myFH6QFGFzwFK7JTmkR5mQMK9yYkSxzAShdkDZUf2k0GzP@tcp(87.239.129.130:3306)/default?parseTime=true&tls=skip-verify"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
