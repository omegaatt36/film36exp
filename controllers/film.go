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

// CreateOneFilm create one roll film.
func CreateOneFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm model.Film
	_ = json.NewDecoder(r.Body).Decode(&newFilm)
	newFilm.ID = primitive.NewObjectID()
	newFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	newFilm.CreateAt = newFilm.UptateAt
	db.Create(db.CollectionFilm, newFilm)
	responseWithJSON(w, http.StatusCreated, newFilm)
}

// GetOneFilm get one roll film by "ID".
func GetOneFilm(w http.ResponseWriter, r *http.Request) {
	var film model.Film
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	if err != nil {
		responseWithJSON(w, http.StatusOK, map[string]string{"message": "id error"})
		return
	}
	if err := db.FindOne(db.CollectionFilm, id).Decode(&film); err != nil {
		// responseWithError(w, http.StatusInternalServerError, err)
		responseWithJSON(w, http.StatusOK, map[string]string{"message": err.Error()})
		return
	}
	responseWithJSON(w, http.StatusOK, film)
}

// GetAllFilms get all films.
func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	var films []model.Film
	collection := db.GetCollection(db.CollectionFilm)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var film model.Film
		cursor.Decode(&film)
		films = append(films, film)
	}
	if err := cursor.Err(); err != err {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	responseWithJSON(w, http.StatusOK, films)
}

// UpdateFilm modify one roll film by "ID".
func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	if err != nil {
		responseWithJSON(w, http.StatusOK, map[string]string{"message": "id error"})
		return
	}
	var updatedFilm model.Film
	json.Unmarshal(reqBody, &updatedFilm)
	updatedFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	updatedFilm.ID = id

	if _, err := db.Update(db.CollectionFilm, id, updatedFilm); err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	responseWithJSON(w, http.StatusOK, updatedFilm)
}

// DeleteFilm remove one roll film by "ID".
func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	if err != nil {
		responseWithJSON(w, http.StatusOK, map[string]string{"message": "id error"})
		return
	}
	if _, err := db.Delete(db.CollectionFilm, id); err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getFilmByID(filmID string) (film *model.Film, err error) {
	id, _ := primitive.ObjectIDFromHex(filmID)
	collection := db.GetCollection(db.CollectionFilm)
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	err = collection.FindOne(ctx, model.Film{ID: id}).Decode(&film)
	return film, err
}
