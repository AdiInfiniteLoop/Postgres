package main

import (
	"github.com/AdiInfiniteLoop/xDev_postgres/db"
	"log"
)

func main() {
	err := db.DBConnection()
	if err != nil {
		log.Println("Cannot Connect", err)
	}
	Connect()

	err = db.DBClose()
	if err != nil {
		log.Println("Cannot Close", err)
	}
}
