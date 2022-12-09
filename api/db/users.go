package db

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	jwt "maxinteg-admin-go/helpers/jwt"
)

type User struct {
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("An Unexpected Error Occurred. Please try again later.")
		return
	}

	userEmail := u.Email
	userPassword := u.Password

	client, err := GetFirebase().Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	doc := client.Collection("users").Where("user_email", "==", userEmail).Documents(ctx)
	document, err := doc.Next()
	documentData := document.Data()

	documentData["user_id"] = document.Ref.ID

	passwordValid := bcrypt.CompareHashAndPassword([]byte(documentData["user_password"].(string)), []byte(userPassword))
	if passwordValid != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid password")
		return
	}

	delete(documentData, "user_password")

	token, err := jwt.CreateToken(document.Ref.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("An Unexpected Error Occurred. Please try again later.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Set-Cookie", "token="+token+"; Max-Age=3600; HttpOnly; SameSite=None; Secure; Path=/")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(documentData)

	defer client.Close()
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Set-Cookie", "token=; HttpOnly; SameSite=None; Max-Age=0; secure;")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Logged out")
}

func GetUserByToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := w.Header().Get("User")

	client, err := GetFirebase().Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	doc := client.Collection("users").Doc(userID)
	document, err := doc.Get(ctx)
	documentData := document.Data()

	delete(documentData, "user_password")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(documentData)

}
