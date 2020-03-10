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

// type allFIlms []*model.Film

// CreateOneFilm create one roll film.
func CreateOneFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm model.Film
	_ = json.NewDecoder(r.Body).Decode(&newFilm)
	newFilm.ID = primitive.NewObjectID()
	newFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	newFilm.CreateAt = newFilm.UptateAt
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	// result, _ := collection.InsertOne(ctx, newFilm)
	collection.InsertOne(ctx, newFilm)
	responseWithJSON(w, http.StatusCreated, newFilm)
}

// GetOneFilm get one roll film by "ID".
func GetOneFilm(w http.ResponseWriter, r *http.Request) {
	film, err := getFilmByID(mux.Vars(r)["filmID"])
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	responseWithJSON(w, http.StatusOK, film)
}

// GetAllFilms get all films.
func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	var films []model.Film
	collection := db.Films()
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
	var updatedFilm model.Film
	json.Unmarshal(reqBody, &updatedFilm)
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	filter := bson.M{"_id": bson.M{"$eq": id}}
	updatedFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	update := bson.M{"$set": updatedFilm}
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	if _, err := collection.UpdateMany(ctx, filter, update); err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	updatedFilm.ID = id
	responseWithJSON(w, http.StatusOK, updatedFilm)
}

// DeleteFilm remove one roll film by "ID".
func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["filmID"])
	filter := bson.M{"_id": bson.M{"$eq": id}}
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		responseWithError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getFilmByID(filmID string) (film *model.Film, err error) {
	id, _ := primitive.ObjectIDFromHex(filmID)
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	err = collection.FindOne(ctx, model.Film{ID: id}).Decode(&film)
	return film, err
}
