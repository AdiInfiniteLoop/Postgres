package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
	//_ "github.com/lib/pq" //_ means we only need the side effects of the package, eg init(),
	//Here, we only need to set up the driver
)

var DB *sql.DB

func OpenPostgresConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	//DB, err = sql.Open("postgres", "postgres://adiinfiniteloop:mysecretpassword@localhost:5432/mydatabase")

	conn, err := pgx.Connect(ctx, "postgres://adiinfiniteloop:mysecretpassword@localhost:5432/newdatabase")

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal("Database Connection Error", err)
		}
	}(conn, ctx)

	if err != nil {
		return err
	}
	fmt.Println("Connected to DB Successfully!")
	return nil
}

func CloseDatabaseConnection() error {
	return DB.Close()
}
