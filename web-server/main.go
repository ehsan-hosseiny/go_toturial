package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name   string
	Family string
	Age    int
}

func createUser(w http.ResponseWriter, r *http.Request, users *[]user) {

	name := r.FormValue("name")
	family := r.FormValue("family")
	// age := r.FormValue("age")

	*users = append(*users, user{
		Name:   name,
		Family: family,
		// Age:    age,
	})
	fmt.Println(*users)

}
func updateUser(w http.ResponseWriter, r *http.Request) {
	//TODO should handle update user myself

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO should handle delete user myself

}
func getUser(w http.ResponseWriter, r *http.Request, users *[]user) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*users)

}

func main() {

	// var users []user ---> in this scenario if we create user wont add and each time just show one user

	users := new([]user)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUser(w, r, users)
			fmt.Println("Method : GET")
		case "POST":
			createUser(w, r, users)
		case "DELETE":
			fmt.Println("Method : DELETE")
		case "PUT":
			fmt.Println("Method : PUT")
		}
		fmt.Fprintf(w, "wellcome in golang")
	})

	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal(err)
	}

}
