package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/httpResponse"
	"net/http"
	"time"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	saveErr := admin.Create()
	if saveErr != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	fmt.Println(admin)
	httpResponse.ResponseWithJson(w, http.StatusCreated, map[string]string{"status": "admin added"})
}

var admin model.Admin

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		httpResponse.ResponseWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()
	email := admin.Email
	var admin2 model.Admin
	loginErr := admin2.Check(email)
	if loginErr != nil {
		switch loginErr {
		case sql.ErrNoRows:
			httpResponse.ResponseWithError(w, http.StatusUnauthorized, "invalid login")
		default:
			httpResponse.ResponseWithError(w, http.StatusBadRequest, "error in database")
		}
		return
	}
	if admin.Password != admin2.Password {
		httpResponse.ResponseWithError(w, http.StatusUnauthorized, "invalid login")
		return
	}
	//create a cookie
	cookie := http.Cookie{
		Name: "admin-cookie",
		// Value: email +admin.Password,
		Value:   "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	//set cookie and send back to client
	http.SetCookie(w, &cookie)
	httpResponse.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "successful"})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "admin-cookie",
		Expires: time.Now(),
	})
	fmt.Println("logout successful")
	httpResponse.ResponseWithJson(w, http.StatusOK, map[string]string{"message": "logout successful"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("admin-cookie")
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			httpResponse.ResponseWithError(w, http.StatusSeeOther, "cookie not set")
		default:
			httpResponse.ResponseWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return false
	}
	fmt.Println(admin.Email, "email")
	fmt.Println(admin.Password, "password")
	if cookie.Value != "my-value" {
		httpResponse.ResponseWithError(w, http.StatusSeeOther, "invalid cookie")
		return false
	}
	return true
}
