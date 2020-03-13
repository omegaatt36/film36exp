package controllers

import (
	"context"
	"encoding/json"
	"film36exp/db"
	"film36exp/model"
	"film36exp/utility"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePic create one model.Pic and append into film by ID.
// if films have not filmID, will not create one film.
func CreatePic(w http.ResponseWriter, r *http.Request) {

	var newPic model.Pic
	_ = json.NewDecoder(r.Body).Decode(&newPic)
	newPic.ID = primitive.NewObjectID()
	newPic.UserName = r.Header.Get("userName")
	filmID := mux.Vars(r)["filmID"]
	fid, _ := primitive.ObjectIDFromHex(filmID)
	if isFilmExist(fid) == false {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "film not found", Result: utility.ResFailed})
		return
	}
	newPic.FID = fid
	db.Create(db.CollectionPic, newPic)

	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess, Data: newPic})
}

// UpdatePic modify one pick by "picID".
func UpdatePic(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "param error", Result: utility.ResFailed})
		return
	}
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["picID"])
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "id not found", Result: utility.ResFailed})
		return
	}
	var updatedPic model.Pic
	json.Unmarshal(reqBody, &updatedPic)
	updatedPic.ID = id

	if _, err := db.Update(db.CollectionPic, id, r.Header.Get("userName"), updatedPic); err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB update error", Result: utility.ResFailed})
		return
	}
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess})
}

// GetPics get pics in the film by "filmID".
func GetPics(w http.ResponseWriter, r *http.Request) {
	var pics []model.Pic
	collection := db.GetCollection(db.CollectionPic)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "pic not found", Result: utility.ResFailed})
		return
	}
	defer cursor.Close(ctx)
	filmID := mux.Vars(r)["filmID"]
	fid, _ := primitive.ObjectIDFromHex(filmID)
	for cursor.Next(ctx) {
		var pic model.Pic
		cursor.Decode(&pic)
		if pic.FID == fid && pic.UserName == r.Header.Get("userName") {
			pics = append(pics, pic)
		}
	}
	if err := cursor.Err(); err != err {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: err.Error(), Result: utility.ResFailed})
		return
	}
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess, Data: pics})
}
