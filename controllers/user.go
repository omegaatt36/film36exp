package controllers

import (
	"encoding/json"
	"net/http"

	"film36exp/db"
	"film36exp/model"

	"golang.org/x/crypto/bcrypt"
)

func isUserExist(userName string) bool {
	if singleResult := db.FindOne(db.CollectionUser, model.User{UserName: userName}); singleResult.Err() != nil {
		return false
	}
	return true
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		responseWithJSON(w, http.StatusBadRequest, Responce{Message: "bad param", Result: ResFailed})
		return
	}
	if isUserExist(user.UserName) {
		responseWithJSON(w, http.StatusBadRequest, Responce{Message: "user exist", Result: ResFailed})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, Responce{Message: "password parse error", Result: ResFailed})
		return
	}
	user.Password = string(hash)
	_, err = db.Create(db.CollectionUser, user)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, Responce{Message: "DB insert error", Result: ResFailed})
		return
	}
	responseWithJSON(w, http.StatusOK, Responce{Result: ResSuccess})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		responseWithJSON(w, http.StatusBadRequest, Responce{Message: "bad param", Result: ResFailed})
		return
	}
	singleResult := db.FindOne(db.CollectionUser, model.User{UserName: user.UserName})
	if singleResult.Err() != nil {
		responseWithJSON(w, http.StatusBadRequest, Responce{Message: "user not exist", Result: ResFailed})
		return

	}
	var userInDB model.User
	err = singleResult.Decode(&userInDB)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, Responce{Message: "DB decode error", Result: ResFailed})
		return
	}
	eq := bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(user.Password))
	if eq != nil {
		responseWithJSON(w, http.StatusInternalServerError, Responce{Message: "password error", Result: ResFailed})
		return
	}
	responseWithJSON(w, http.StatusOK, Responce{Result: ResSuccess})
}
