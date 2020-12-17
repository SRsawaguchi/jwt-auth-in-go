package endpoint

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SRsawaguchi/jwt-auth-in-go/internal/auth"
	"github.com/SRsawaguchi/jwt-auth-in-go/internal/model"
)

var (
	// SecretKey is secret key for creating JWT
	SecretKey = "CsrdR1twFMFzXiw"
)

// Endpoint represents endpoint
type Endpoint struct {
	Users map[string]*model.User
}

// LoginRequest represents request of login
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password`
}

// LoginResponse represents response of login
type LoginResponse struct {
	Token string `json:"token"`
}

// Login handles user login
func (e *Endpoint) Login(w http.ResponseWriter, r *http.Request) {
	// リクエストを読み込み
	loginReq := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	// パスワードを比較する
	user, ok := e.Users[loginReq.Name]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong username or password")
		return
	}
	if !auth.IsValidPassword(loginReq.Password, user.Password) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "wrong username or password")
		return
	}

	// JWTトークンを作成
	token, err := auth.GenerateToken(user.Name, SecretKey, time.Second*24)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}

	// レスポンスを作成
	json.NewEncoder(w).Encode(&LoginResponse{
		Token: token,
	})
}

// SigninRequest represents request of signin
type SigninRequest struct {
	Name     string `json:"name"`
	Password string `json:"password`
}

// SigninResponse represents response of signin
type SigninResponse struct {
	Name  string `json:"name"`
	Token string `json:"token`
}

// Signin creates new user
func (e *Endpoint) Signin(w http.ResponseWriter, r *http.Request) {
	// リクエストを読み込み
	signinReq := SigninRequest{}
	err := json.NewDecoder(r.Body).Decode(&signinReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	// パスワードをハッシュ化して保存(通常は別のレイヤーに書く)
	passwordHash, err := auth.GeneratePasswordHash(signinReq.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	user := &model.User{
		Name:     signinReq.Name,
		Password: passwordHash,
	}
	e.Users[user.Name] = user

	// JWTトークンを作成(有効期限は1日。ただし、いろいろと変えて試してみて。)
	token, err := auth.GenerateToken(user.Name, SecretKey, time.Hour*24)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	// レスポンスを作成
	err = json.NewEncoder(w).Encode(&SigninResponse{
		Name:  user.Name,
		Token: token,
	})
}

// HelloResponse represents response of hello
type HelloResponse struct {
	Message string `json:"message"`
}

// Hello is greeting logged in user
func (e *Endpoint) Hello(w http.ResponseWriter, r *http.Request) {
	// JWTトークンを検証
	tokenString := r.Header.Get("Token")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "there is no token")
		return
	}

	name, err := auth.ParseToken(tokenString, SecretKey)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, err.Error())
		return
	}

	// レスポンスを作成
	json.NewEncoder(w).Encode(&HelloResponse{
		Message: fmt.Sprintf("Hello, %v", name),
	})
}
