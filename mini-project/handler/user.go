package handler

import (
	"database/sql"
	"fmt"
	"mini/utils"
	"net/http"
	"strconv"
)

type user struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

type userRequest struct {
	Id    uint
	Name  string
	Email string
	Age   uint
}

func UserHandler(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		userList(w, r, conn)
	case "POST":
		userCreate(w, r, conn)
	case "PUT":
		userUpdate(w, r, conn)
	case "DELETE":
		userDelete(w, r, conn)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func userCreate(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var request userRequest

	request.Name = r.FormValue("name")
	request.Email = r.FormValue("email")
	age64, err := strconv.ParseUint(r.FormValue("age"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	request.Age = uint(age64)

	_, err = conn.Exec(`INSERT INTO users (name,email,age) VALUES ($1,$2,$3)`, request.Name, request.Email, request.Age)

	if err != nil {
		utils.ResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, r, http.StatusCreated, request)

}

func userUpdate(w http.ResponseWriter, r *http.Request, conn *sql.DB) {

	var request userRequest

	request.Name = r.FormValue("name")
	request.Email = r.FormValue("email")
	age64, err := strconv.ParseUint(r.FormValue("age"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	request.Age = uint(age64)

	id64, err := strconv.ParseUint(r.FormValue("id"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	request.Id = uint(id64)

	res, err := conn.Exec(`UPDATE users SET name=$1 ,email=$2,age=$3 WHERE id=$4`, request.Name, request.Email, request.Age, id64)

	if err != nil {
		utils.ResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		utils.ResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	utils.ResponseWithJson(w, r, http.StatusCreated, map[string]int64{"RowsAffected": count})

}

func userDelete(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var id uint

	id64, err := strconv.ParseUint(r.FormValue("id"), 10, 32)
	if err != nil {
		utils.ResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	id = uint(id64)

	_, err = conn.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		utils.ResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseWithJson(w, r, http.StatusOK, map[string]string{"message": "ok"})

}

func userList(w http.ResponseWriter, r *http.Request, conn *sql.DB) {
	var users []user

	rows, err := conn.Query("SELECT id,name,email,age FROM users")
	if err != nil {
		utils.ResponseWithError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	for rows.Next() {
		var user user
		err := rows.Scan(&user.ID, &user.Name, &user.Name, &user.Email)
		if err != nil {
			utils.ResponseWithError(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		users = append(users, user)
	}

	utils.ResponseWithJson(w, r, http.StatusOK, users)
}
