package server

import (
	"database/sql"
	"fmt"
	"mini/config"
	"mini/handler"
	"net/http"
)

func StartServer(config config.Config, conn *sql.DB) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.UserHandler(w, r,conn)
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/products")
	})
	err := http.ListenAndServe(config.AppPort, nil)
	return err
}
