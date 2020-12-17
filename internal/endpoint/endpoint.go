package endpoint

import (
	"fmt"
	"net/http"
)

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet: Login")
}

// Signin creates new user
func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet: Signin")
}

// RefreshToken regenerates new JWT token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet: RefleshToken")
}

// NewMemo creates new memo
func NewMemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet: NewMemo")
}

// GetMemo retrieves memos of specific user
func GetMemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented yet: GetMemo")
}
