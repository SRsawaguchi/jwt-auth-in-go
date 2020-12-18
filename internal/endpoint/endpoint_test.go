package endpoint

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SRsawaguchi/jwt-auth-in-go/internal/auth"
	"github.com/SRsawaguchi/jwt-auth-in-go/internal/model"
)

var mux *http.ServeMux
var endpoint *Endpoint

func TestMain(m *testing.M) {
	setUp()
	m.Run()
}

func setUp() {
	endpoint = NewEndpoint()
	mux = http.NewServeMux()
	mux.HandleFunc("/signin", endpoint.Signin)
	mux.HandleFunc("/login", endpoint.Login)
	mux.HandleFunc("/hello", endpoint.Hello)
}

func signinForTest(name, password string) (*model.User, error) {
	hash, err := auth.GeneratePasswordHash(password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		ID:       1,
		Name:     name,
		Password: hash,
	}

	endpoint.Users[user.Name] = user
	return user, nil
}

func TestSignin(t *testing.T) {
	data := `{ "name": "Kade", "password": "qwerty" }`
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/signin", strings.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != http.StatusOK {
		t.Errorf("Response code is not 200 but %v.", writer.Code)
	}

	user, ok := endpoint.Users["Kade"]
	if ok {
		t.Error("Could not save user: 'Kade'")
	}

	resp := &SigninResponse{}
	err := json.NewDecoder(writer.Body).Decode(&resp)
	if err != nil {
		t.Error(err.Error())
	}

	if resp.Name != user.Name {
		t.Errorf("Invalid Name: expects '%v' but got '%v'", user.Name, resp.Name)
	}
}

type loginTestCase struct {
	Title      string
	Req        LoginRequest
	StatusCode int
	IsError    bool
}

func TestLogin(t *testing.T) {
	name := "Kade"
	password := "ntRBMHISTKwltVs"
	_, err := signinForTest(name, password)
	if err != nil {
		t.Error(err.Error())
	}

	t.Run("Login success", func(t *testing.T) {
		testCases := []loginTestCase{
			{
				Title:      "Correct name and password",
				Req:        LoginRequest{Name: name, Password: password},
				StatusCode: http.StatusOK,
				IsError:    false,
			},
		}
		for _, tc := range testCases {
			testLogin(t, &tc)
		}
	})

	t.Run("Login failed", func(t *testing.T) {
		testCases := []loginTestCase{
			{
				Title:      "Invalid name",
				Req:        LoginRequest{Name: "invalid", Password: password},
				StatusCode: http.StatusUnauthorized,
				IsError:    true,
			},
			{
				Title:      "Invalid password",
				Req:        LoginRequest{Name: name, Password: "Invalid"},
				StatusCode: http.StatusUnauthorized,
				IsError:    true,
			},
			{
				Title:      "Empty data",
				Req:        LoginRequest{Name: "", Password: ""},
				StatusCode: http.StatusUnauthorized,
				IsError:    true,
			},
		}
		for _, tc := range testCases {
			testLogin(t, &tc)
		}
	})
}

func testLogin(t *testing.T, tc *loginTestCase) {
	t.Helper()

	data, _ := json.Marshal(tc.Req)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/login", bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	mux.ServeHTTP(writer, request)

	if writer.Code != tc.StatusCode {
		t.Errorf("Response code is not %v but %v.", tc.StatusCode, writer.Code)
	}

	if tc.IsError {
		return
	}

	resp := &LoginResponse{}
	err := json.NewDecoder(writer.Body).Decode(resp)
	if err != nil {
		t.Error(err.Error())
	}

	if resp.Token == "" {
		t.Error("Token is empty.")
	}

	name, err := auth.ParseToken(resp.Token, SecretKey)
	if err != nil {
		t.Error(err.Error())
	}
	if name != tc.Req.Name {
		t.Error("Invalid name in token.")
	}
}
func TestHello(t *testing.T) {
	name := "Kade"
	password := "ntRBMHISTKwltVs"
	_, err := signinForTest(name, password)
	if err != nil {
		t.Error(err.Error())
	}

	t.Run("Logged in", func(t *testing.T) {

	})
	t.Run("Not logged in", func(t *testing.T) {})
}
