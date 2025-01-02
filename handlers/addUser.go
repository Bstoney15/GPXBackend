package handlers

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "GPXBackend/dbhandler"
    "GPXBackend/datavalidators"
)

type user struct {
    Username      string `json:"username"`
    Email         string `json:"email"`
    Password_hash string `json:"password_hash"`
    First_name    string `json:"first_name"`
    Last_name     string `json:"last_name"`
}



func AddUser(w http.ResponseWriter, r *http.Request) {
    var u user
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if !datavalidators.IsValidEmail(u.Email) {
        http.Error(w, "Invalid email", 462)
        return
    }

    if !datavalidators.IsValidUsername(u.Username) {
        http.Error(w, "Invalid username", 463)
        return
    }
    
    if !datavalidators.IsValidName(u.First_name) || !datavalidators.IsValidName(u.Last_name) {
        http.Error(w, "Invalid name", 464)
        return
    }

    _, err2 := dbhandler.RemoteDB.Exec(context.Background(), "INSERT INTO users (username, email, password_hash, first_name, last_name) VALUES ($1, $2, $3, $4, $5)", u.Username, u.Email, u.Password_hash, u.First_name, u.Last_name)
    if err2 != nil {
        http.Error(w, err2.Error(), 461)
        fmt.Println("User add failed:", err2)
        return
    }

    fmt.Fprintln(w, "User added successfully")
}