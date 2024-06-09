package contact

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"luxcare/database"
	"net/http"
)

type Contact struct {
	Name     string
	Email    string
	Phone    string
	Comments string
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hellloooooooooo")
}
func Create(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		content := req.Header.Get("Content-Type")
		if content != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		}
		defer req.Body.Close()
		var target Contact
		err := json.NewDecoder(req.Body).Decode(&target)
		if err != nil {
			panic(err)
		}
		if len(target.Phone) <= 0 {
			http.Error(w, "Phone is required", http.StatusBadRequest)
		}

		db := database.Connect()
		defer db.Close()
		query := "INSERT INTO `contacts` (`name`, `email`, `phone`, `comments`, `created_at`) VALUES (?, ?, ?, ?, NOW())"

		result, err := db.ExecContext(context.Background(), query,
			target.Name, target.Email, target.Phone, target.Comments,
		)

		if err != nil {
			fmt.Printf("impossible insert contact: %s", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

		id, err := result.LastInsertId()
		if err != nil {
			fmt.Printf("impossible to retrieve last inserted id: %s", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

		log.Printf("inserted id: %d", id)
	case http.MethodOptions:
		fmt.Fprintln(w, req.Method)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
