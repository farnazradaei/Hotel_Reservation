package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"Hotel_reservation/internal/user"
)

type UserRequest struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

func Signup(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "method ot allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UserRequest 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json" , http.StatusBadRequest)
		return
	}

	if err:= validatsignup(req); err != nil {
		http.Error(w ,err.Error(), http.StatusBadRequest)
		return
	}

	err := user.Signup(req.Email, req.Password)
	if err != nil {
		http.Error(w, "signup failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("signup successful"))
}


func validatsignup (r UserRequest) error{
	if r.Email == "" {
		return errors.New("email is required")
	}

	if len(r.Email)< 5 {
		return  errors.New("email is too short")
	}

	emailRegex := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	if !emailRegex.MatchString(r.Email){
		return errors.New("invalid email format")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	if len(r.Password) < 8 {
		return  errors.New("password must be at least 8 chars")
	}
	return nil
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}

    token, err := user.Login(req.Email, req.Password)
    if err != nil {
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }

    resp := map[string]string{"token": token}
	w.Header().Set("content-type", "application/json")
    json.NewEncoder(w).Encode(resp)

}
