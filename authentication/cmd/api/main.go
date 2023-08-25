package main

import (
	"authentication/cmd/api/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var dbConnectAttempts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

/*
// --- this is the lib to connect to postgress DB
go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/jackc/pgx/v4/stdlib

*/

func main() {

	log.Printf("Starting authentication service on port %s\n", webPort)

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can not connect to Postgres!")
	}

	// setup config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	// http define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	// read value from environment
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			fmt.Println("Postgres not yet ready ...")
			dbConnectAttempts++
		} else {
			fmt.Println("Connected to Postgres!")
			return connection
		}

		if dbConnectAttempts > 10 {
			fmt.Println(err)
			return nil
		}

		fmt.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
