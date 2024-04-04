package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/omegaatt36/film36exp/auth"
	"github.com/omegaatt36/film36exp/db"
	"github.com/omegaatt36/film36exp/model"
	"github.com/omegaatt36/film36exp/utility"

	"golang.org/x/crypto/bcrypt"
)

func isUserExist(userName string) bool {
	if singleResult := db.FindOne(db.CollectionUser, model.User{UserName: userName}); singleResult.Err() != nil {
		return false
	}
	return true
}

// Register name & cryp(pwd) into DB
func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		utility.ResponseWithJSON(w, http.StatusBadRequest, utility.Response{Message: "bad param", Result: utility.ResFailed})
		return
	}
	if isUserExist(user.UserName) {
		utility.ResponseWithJSON(w, http.StatusBadRequest, utility.Response{Message: "user exist", Result: utility.ResFailed})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "password parse error", Result: utility.ResFailed})
		return
	}
	user.Password = string(hash)
	_, err = db.Create(db.CollectionUser, user)
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB insert error", Result: utility.ResFailed})
		return
	}
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess})
}

// Login verify name & pwd and return token
func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserName == "" || user.Password == "" {
		utility.ResponseWithJSON(w, http.StatusBadRequest, utility.Response{Message: "bad param", Result: utility.ResFailed})
		return
	}
	singleResult := db.FindOne(db.CollectionUser, model.User{UserName: user.UserName})
	if singleResult.Err() != nil {
		utility.ResponseWithJSON(w, http.StatusBadRequest, utility.Response{Message: "user not exist", Result: utility.ResFailed})
		return

	}
	var userInDB model.User
	err = singleResult.Decode(&userInDB)
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB decode error", Result: utility.ResFailed})
		return
	}
	eq := bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(user.Password))
	if eq != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "password error", Result: utility.ResFailed})
		return
	}
	token, _ := auth.GenerateToken(&user)
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess, Data: model.JwtToken{Token: token}})
}
