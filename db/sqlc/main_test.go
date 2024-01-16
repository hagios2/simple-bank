package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB
var err error

func TestMain(m *testing.M) {
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
