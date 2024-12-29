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
	endpoint := os.Getenv("AWS_RDL_ENDPOINT")
	port := os.Getenv("AWS_RDL_PORT")
	user := os.Getenv("AWS_RDL_USERNAME")
	password := os.Getenv("AWS_RDL_PASSWORD")
	dbname := os.Getenv("AWS_RDL_DBNAME")

	

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", endpoint, port, user, password, dbname)
	fmt.Println("connect String:", connectionString)
	fmt.Print("Connecting to remote database\n\n")


	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Println("Error connecting to remote database")
		return nil, err
	}

	fmt.Print("Connected to remote database\n\n")
	return conn, nil
}