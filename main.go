package main

import(
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func test(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, " :test")
}


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



	mux.HandleFunc("/test", test)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}