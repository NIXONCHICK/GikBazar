package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	var filePath = "login.html"
	t, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Println("Error while parsing file")
		return
	}
	err = t.ExecuteTemplate(w, filePath, nil)
	if err != nil {
		fmt.Println("Error name: ", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	db, err := sql.Open("mysql", "root:root@/goWeb")
	if err != nil {
		log.Println(err)
	}
	hashedPassword := sha256.Sum256([]byte(password))
	hashedString := hex.EncodeToString(hashedPassword[:])
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	if username != "" && password != "" {
		_, err = db.Exec(query, username, hashedString)
		fmt.Fprintf(w, `
            <html>
                <head>
                    <script>
                        window.location.href = '/index';
                    </script>
                </head>
                <body></body>
            </html>
        `)
		return
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	var filePath = "index.html"
	t, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Println("Error while parsing file")
		return
	}
	err = t.ExecuteTemplate(w, filePath, nil)
	if err != nil {
		fmt.Println("Error name: ", err)
		return
	}
}

func addCard(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/index":
		index(w, r)
	case "/add/itemcard":
		addCard(w, r)
	default:
		fmt.Println("Defunct page")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Запуск HTTP сервера на порту 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
