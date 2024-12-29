package dbhandler

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)


var (
	// RemoteDB is the connection to the remote database
	RemoteDB *pgx.Conn
)

func Init() { 
	Remote, err := getRemote()
	if err != nil {
		fmt.Println("Error getting remote connection")
		fmt.Println(err)
		return
	}

	RemoteDB = Remote
}

// GetRemote returns a connection to the remote database
func getRemote() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error connecting to remote database")
		return nil, err
	}
	return conn, nil
}