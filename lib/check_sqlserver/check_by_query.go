package lib

import (
	"database/sql"
	"flag"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

type Database struct {
	UserName string
	PassWord string
	Host     string
	Instance string
	Port     int
}

func Set() *Database {
	var username = flag.String("u", "root", "select database username")
	var password = flag.String("p", "root", "select database password")
	var host = flag.String("h", "localhost", "select database host")
	var instance = flag.String("i", "", "select database instance")
	var port = flag.Int("P", 1433, "select database port")
	flag.Parse()

	database := &Database{
		UserName: *username,
		PassWord: *password,
		Host:     *host,
		Instance: *instance,
		Port:     *port,
	}

	return database
}

func (database *Database) Connect() *sql.DB {
	connectionTimeout := 60
	query := url.Values{}
	query.Add("connection timeout", fmt.Sprintf("%d", connectionTimeout))

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(database.UserName, database.PassWord),
		Host:   fmt.Sprintf("%s:%d", database.Host, database.Port),
		Path:   database.Instance,
	}

	db, err := sql.Open("sqlserver", u.String())
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
