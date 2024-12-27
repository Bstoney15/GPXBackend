package main

import(
	"fmt"
	"net/http"
	"os"
	"GPXBackend/handlers"
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



	mux.HandleFunc("/test", handlers.Test)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}