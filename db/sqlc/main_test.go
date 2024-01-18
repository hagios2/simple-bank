package db

import (
	"database/sql"
	"github.com/hagios2/simple-bank/util"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Can not read env variables", err)
	}
	testDB, err = sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
