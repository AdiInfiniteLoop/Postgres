package main

import (
	"github.com/AdiInfiniteLoop/xDev_postgres/db"
	"log"
)

func main() {
	//defer an anonymous function
	defer func() {
		err := db.CloseDatabaseConnection()
		if err != nil {
			log.Fatal("Cannot Close", err)
		}
	}()

	err := db.OpenPostgresConnection()
	if err != nil {
		log.Fatal("Cannot Connect", err)
	}
	Connect()

}
