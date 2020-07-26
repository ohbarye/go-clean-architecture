package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type SqlHandler struct {
	Conn *sql.DB
}

var (
	username = "root"
	host     = "db"
	port     = 3306
	database = "sample"
)

func dataSourceName() string {
	return fmt.Sprintf("%s:@tcp(%s:%d)/%s", username, host, port, database)
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", dataSourceName())
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := SqlHandler{Conn: conn}
	return &sqlHandler
}
