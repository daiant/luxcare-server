package main

import (
	"fmt"
	"luxcare/contact"
	"luxcare/database"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	db := database.Connect()
	_, table_check := db.Query("select Count(*) from contacts;")

	if table_check == nil {
		fmt.Println("table is there, nothing to do.")
	} else {
		fmt.Println("table not there, creating table")
		database.GenerateMigrations()
	}
	http.HandleFunc("/hello", contact.Hello)
	http.HandleFunc("/create-contact", contact.Create)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":3000", nil)
}
