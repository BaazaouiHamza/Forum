package main

import (
	"database/sql"
	"log"

	"github.com/hamza-baazaoui/forum/api"
	db "github.com/hamza-baazaoui/forum/db/sqlc"
	"github.com/hamza-baazaoui/forum/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	db := db.New(conn)
	server := api.NewServer(db)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
