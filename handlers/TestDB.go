package handlers

import (
	"net/http"
	"fmt"
	"GPXBackend/dbhandler"
	"context"
	"time"
)





func Test(w http.ResponseWriter, r *http.Request){
	rows, err := dbhandler.RemoteDB.Query(context.Background(), "SELECT * FROM tests")
	if err != nil {
		fmt.Println("Error getting remote connection")
		fmt.Println(err)
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		var testdata string
		var testtime time.Time
		err := rows.Scan(&testdata, &testtime)
		if err != nil {
			fmt.Println("Error scanning row")
			fmt.Println(err)
			return
		}
		fmt.Println(testdata, testtime)
	}


	_, err = dbhandler.RemoteDB.Exec(context.Background(), "INSERT INTO tests (testdata, testtime) VALUES ($1, $2)", "test", time.Now())
	if err != nil {
		fmt.Println("Error inserting row")
		fmt.Println(err)
		return
	}

	
	return
}