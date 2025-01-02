package main

import (
	"GPXBackend/dbhandler"
	"GPXBackend/handlers"
	"context"
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
)




func main(){

	mux := http.NewServeMux()
	if mux == nil {
		fmt.Println("Error creating mux")
		return
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	fmt.Print("Loaded .env file\n\n")
	dbhandler.Init() // Initialize database connection
	fmt.Print("Initialized database connection\n\n")

	mux.HandleFunc("/test", handlers.Test)
	mux.HandleFunc("/Add_User", handlers.AddUser)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	defer dbhandler.RemoteDB.Close(context.Background()) // Close database connection when server is closed
	fmt.Println("Server running on port 8080")
	server.ListenAndServe() // Start server
} 