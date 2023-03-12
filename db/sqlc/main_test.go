package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
)

func TestMain(m *testing.M) {

	conn, erroSqlOpen := sql.Open(dbDriver, dbSource)
	if erroSqlOpen != nil {
		log.Fatal(erroSqlOpen)
	}
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}
