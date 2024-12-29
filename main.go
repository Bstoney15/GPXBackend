package main

import(
	"fmt"
	"net/http"
	"GPXBackend/handlers"
	"GPXBackend/dbhandler"
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

	fmt.Println("Loaded .env file")
	dbhandler.Init()
	fmt.Print("Initialized database connection")



	mux.HandleFunc("/test", handlers.Test)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}


	fmt.Println("Server running on port 8080")
	server.ListenAndServe()
}