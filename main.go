package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var db *gorm.DB

func initDB() {
	var err error
	dsn := "host=db user=user password=password dbname=authdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&User{})
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword := hashPassword(creds.Password)
	user := User{Username: creds.Username, Password: hashedPassword}
	result := db.Create(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user User
	db.Where("username = ?", creds.Username).First(&user)

	if user.ID == 0 || !checkPasswordHash(creds.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func getUsetHedler(w http.ResponseWriter, r *http.Request) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем имя пользователя из параметров URL
	vars := mux.Vars(r)
	username := vars["username"]

	// Находим пользователя по имени
	var user User
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Удаляем пользователя
	db.Delete(&user)

	w.WriteHeader(http.StatusOK)
}

func hashPassword(password string) string {
	return password
}

func checkPasswordHash(password, hash string) bool {
	return password == hash
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/signup", signupHandler).Methods("POST")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/users", getUsetHedler).Methods("Get")
	r.HandleFunc("/users/{username}", deleteUserHandler).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
