package controllers

import (
	"context"
	"encoding/json"
	"film36exp/db"
	"film36exp/model"
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
	filmID := mux.Vars(r)["filmID"]
	fid, _ := primitive.ObjectIDFromHex(filmID)
	if isFilmExist(fid) == false {
		responseWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "film not found."})
		return
	}
	newPic.FID = fid
	db.Create(db.CollectionPic, newPic)

	responseWithJSON(w, http.StatusCreated, newPic)
}

// UpdatePic modify one pick by "picID".
func UpdatePic(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["picID"])
	if err != nil {
		responseWithJSON(w, http.StatusOK, map[string]string{"message": "id error"})
		return
	}
	var updatedPic model.Pic
	json.Unmarshal(reqBody, &updatedPic)
	updatedPic.ID = id

	if _, err := db.Update(db.CollectionPic, id, updatedPic); err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetPics get pics in the film by "filmID".
func GetPics(w http.ResponseWriter, r *http.Request) {
	var pics []model.Pic
	collection := db.GetCollection(db.CollectionPic)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close(ctx)
	filmID := mux.Vars(r)["filmID"]
	fid, _ := primitive.ObjectIDFromHex(filmID)
	for cursor.Next(ctx) {
		var pic model.Pic
		cursor.Decode(&pic)
		if pic.FID == fid {
			pics = append(pics, pic)
		}
	}
	if err := cursor.Err(); err != err {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	responseWithJSON(w, http.StatusOK, pics)
}
