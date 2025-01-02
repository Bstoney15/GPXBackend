package handlers

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "GPXBackend/dbhandler"
    "regexp"
)

type user struct {
    Username      string `json:"username"`
    Email         string `json:"email"`
    Password_hash string `json:"password_hash"`
    First_name    string `json:"first_name"`
    Last_name     string `json:"last_name"`
}

func isValidEmail(email string) bool {
	// Simple regex for validating email format
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
    var u user
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if !isValidEmail(u.Email) {
        http.Error(w, "Invalid email", 462)
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