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

// CreateOneFilm create one roll film.
func CreateOneFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm model.Film
	_ = json.NewDecoder(r.Body).Decode(&newFilm)
	newFilm.ID = primitive.NewObjectID()
	newFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	newFilm.CreateAt = newFilm.UptateAt
	db.Create(db.CollectionFilm, newFilm)
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess, Data: newFilm})
}

// GetOneFilm get one roll film by "ID".
func GetOneFilm(w http.ResponseWriter, r *http.Request) {
	var film model.Film
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "param error", Result: utility.ResFailed})
		return
	}
	singleResult := db.FindOne(db.CollectionFilm, model.Film{ID: id})
	if singleResult.Err() != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "film not found", Result: utility.ResFailed})
		return
	}
	err = singleResult.Decode(&film)
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB decode error", Result: utility.ResFailed})
		return
	}
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess, Data: film})
}

// GetAllFilms get all films.
func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	var films []model.Film
	collection := db.GetCollection(db.CollectionFilm)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB error", Result: utility.ResFailed})
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var film model.Film
		cursor.Decode(&film)
		films = append(films, film)
	}
	if err := cursor.Err(); err != err {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB error", Result: utility.ResFailed})
		return
	}
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess, Data: films})
}

// UpdateFilm modify one roll film by "filmID".
func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "param error", Result: utility.ResFailed})
		return
	}
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "id error", Result: utility.ResFailed})
		return
	}
	var updatedFilm model.Film
	json.Unmarshal(reqBody, &updatedFilm)
	updatedFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	updatedFilm.ID = id

	if _, err := db.Update(db.CollectionFilm, id, updatedFilm); err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB update error", Result: utility.ResFailed})
		return
	}

	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess})
}

// DeleteFilm remove one roll film by "filmID".
func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	if err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "id error", Result: utility.ResFailed})
		return
	}
	if _, err := db.Delete(db.CollectionFilm, id); err != nil {
		utility.ResponseWithJSON(w, http.StatusInternalServerError, utility.Response{Message: "DB delete error", Result: utility.ResFailed})
		return
	}
	utility.ResponseWithJSON(w, http.StatusOK, utility.Response{Result: utility.ResSuccess})
}

func isFilmExist(id primitive.ObjectID) bool {
	if singleResult := db.FindOne(db.CollectionFilm, model.Film{ID: id}); singleResult.Err() != nil {
		return false
	}
	return true
}
