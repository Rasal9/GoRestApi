package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("HS256")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.AuthenticateUser(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Не удалось создать токен", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
