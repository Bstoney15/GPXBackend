package main

import(
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, " :test")
}


func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/test", test)


	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()

	fmt.Print("test\n")
}