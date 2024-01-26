package main

import (
	"database/sql"
	"github.com/hagios2/simple-bank/api"
	db "github.com/hagios2/simple-bank/db/sqlc"
	"github.com/hagios2/simple-bank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not read env variables", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("Can not connect to db", conn)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server", conn)
	}
}
