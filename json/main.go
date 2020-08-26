package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country, omitempty"`
}

type User struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Address  Address `json:"adress"`
}

var users = []User{
	{
		Name:     "Bab",
		Password: "secret",
		Email:    "bab@mail.com",
		Address: Address{
			Street:  "Philadelphia",
			City:    "Bordeaux",
			Country: "France",
		},
	},
	{
		Name:     "Bib",
		Password: "secret",
		Email:    "bib@mail.com",
		Address: Address{
			Street:  "Lauvirat",
			City:    "Coutras",
			Country: "",
		},
	},
}

type PasswordJsonBody struct {
	UserIndex         int    `json:"user_index"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat "`
}

func userslist(w http.ResponseWriter, r *http.Request) {

	b, err := json.MarshalIndent(users, "", "  ")

	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	var p PasswordJsonBody
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%v", p)
}

func main() {
	http.HandleFunc("/users", userslist)
	http.HandleFunc("/update_password", updatePassword)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
