package handlers

import (
	"net/http"
	"fmt"
)

func Test(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, " :test")
}