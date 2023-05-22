package main

import (
	"database/sql"
	"log"
	"os"
)

func version() (v string) {
	v = "v1.1.1"
	return
}

func main() {
	var (
		db         *sql.DB
		exoplanets *Exoplanets
	)

	log.Printf("Starting exoplanets %s\n", version())

	if os.Getenv("DB_HOST") != "" {
		db = dbConnect(dbInfo{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			dbname:   os.Getenv("DB_NAME")})
		defer db.Close()
	}

	exoplanets = &Exoplanets{DB: db}
	exoplanets.populate()

	log.Fatal(listenAndServe("8080", exoplanets))
}
