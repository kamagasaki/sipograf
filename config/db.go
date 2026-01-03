package config

import (
	"crypto/tls"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// Register a custom TLS config that skips cert verification (DEV ONLY).
	err := mysql.RegisterTLSConfig("skip-verify", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	dsn := "mariadb:7F37RDMdizgkWC75q5myFH6QFGFzwFK7JTmkR5mQMK9yYkSxzAShdkDZUf2k0GzP@tcp(fs48ws4oo88o4gsk4os4wwos:3306)/default?parseTime=true&tls=skip-verify"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
