package lib

import (
	"database/sql"
	"flag"
	"fmt"
	"net/url"

	_ "github.com/lib/pq"
)

type Database struct {
	UserName string
	PassWord string
	Host     string
	Port     int
	DbName   string
}

func Set() *Database {
	var username = flag.String("u", "root", "select database username")
	var password = flag.String("p", "root", "select database password")
	var host = flag.String("h", "localhost", "select database host")
	var port = flag.Int("P", 5435, "select database port")
	var dbname = flag.String("n", "", "select database name")
	flag.Parse()

	database := &Database{
		UserName: *username,
		PassWord: *password,
		Host:     *host,
		Port:     *port,
		DbName:   *dbname,
	}

	return database
}

func (database *Database) Connect() *sql.DB {
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(database.UserName, database.PassWord),
		Host:     fmt.Sprintf("%s:%d", database.Host, database.Port),
		Path:     database.DbName,
		RawQuery: "sslmode=disable",
	}

	db, err := sql.Open("postgres", u.String())
	defer func() {
		if err := recover(); err != nil {
			db.Close()
		}
	}()

	if err != nil {
		panic(err.Error())
	}

	return db
}
