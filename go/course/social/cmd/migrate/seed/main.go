package main

import (
	"github/Tshewang2022/social/internal/db"
	"github/Tshewang2022/social/internal/env"
	"github/Tshewang2022/social/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social? sslmode=disable")

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store, conn)
}
