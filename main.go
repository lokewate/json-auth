package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserInfo struct {
	Username string
	Password string
}

func authenticated(u UserInfo) bool {
	return true
}

// Expects JSON containing username and password and authenticates the user.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var u UserInfo
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err)
		return
	}
	if authenticated(u) {
		w.Write([]byte("{Status: success}"))
	} else {
		w.Write([]byte("{Status: failure}"))
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
