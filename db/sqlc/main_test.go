package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var queries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
)

func TestMain(m *testing.M) {

	conn, erroSqlOpen := sql.Open(dbDriver, dbSource)
	if erroSqlOpen != nil {
		log.Fatal(erroSqlOpen)
	}
	queries = New(conn)
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("----- ", rand.Intn(10), "--------")
	fmt.Print("----- ", rand.Intn(10), "--------")
	fmt.Print("----- ", rand.Intn(10), "--------")
	log.Print("Conn Ok!!!")
	os.Exit(m.Run())

}
