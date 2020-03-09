package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"film36exp/db"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Film struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Vendor          string             `json:"Vendor,omitempty" bson:"Vendor,omitempty"`
	Production      string             `json:"Production,omitempty" bson:"Production,omitempty"`
	ISO             string             `json:"ISO,omitempty" bson:"ISO,omitempty"`
	ForceProcessing string             `json:"ForceProcessing,omitempty" bson:"ForceProcessing,omitempty"`
	FilmSize        string             `json:"FilmSize,omitempty" bson:"FilmSize,omitempty"`
	FilmType        string             `json:"FilmType,omitempty" bson:"FilmType,omitempty"`
	CreateAt        string             `json:"-" bson:"-"`
	UptateAt        string             `json:"-" bson:"-"`
	Description     string             `json:"Description,omitempty" bson:"Description,omitempty"`
	Pics            []*pic             `json:"-" bson:"-"`
}

type allFIlms []*Film

var films = allFIlms{}

// CreateOneFilm create one roll film.
func CreateOneFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm Film
	_ = json.NewDecoder(r.Body).Decode(&newFilm)
	newFilm.ID = primitive.NewObjectID()
	newFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	newFilm.CreateAt = newFilm.UptateAt
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	result, _ := collection.InsertOne(ctx, newFilm)
	// collection.InsertOne(ctx, newFilm)
	responseWithJSON(w, http.StatusCreated, result)
}

// GetOneFilm get one roll film by "ID".
func GetOneFilm(w http.ResponseWriter, r *http.Request) {

	filmID := mux.Vars(r)["filmID"]
	id, _ := primitive.ObjectIDFromHex(filmID)
	var film Film
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	err := collection.FindOne(ctx, Film{ID: id}).Decode(&film)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, []byte(`{"message" : "`+err.Error()+`"}`))
		return
	}
	responseWithJSON(w, http.StatusOK, film)
}

// GetAllFilms get all films.
func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	var films []Film
	collection := db.Films()
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, []byte(`{"message":"`+err.Error()+`"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var film Film
		cursor.Decode(&film)
		films = append(films, film)
	}
	if err := cursor.Err(); err != err {
		responseWithJSON(w, http.StatusInternalServerError, []byte(`{"message":"`+err.Error()+`"}`))
		return
	}
	responseWithJSON(w, http.StatusOK, films)
}

// // UpdateFilm modify one roll film by "ID".
// func UpdateFilm(w http.ResponseWriter, r *http.Request) {
// 	filmID := mux.Vars(r)["filmID"]
// 	var updatedFilm film
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
// 	}
// 	f, err := getFilmByID(filmID)

// 	if err != nil {
// 		responseWithJSON(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// }

// 	json.Unmarshal(reqBody, &updatedFilm)

// 	f.Production = updatedFilm.Production
// 	f.Vendor = updatedFilm.Vendor
// 	f.ISO = updatedFilm.ISO
// 	f.ForceProcessing = updatedFilm.ForceProcessing
// 	f.FilmSize = updatedFilm.FilmSize
// 	f.FilmType = updatedFilm.FilmType
// 	f.Description = updatedFilm.Description
// 	f.UptateAt = time.Now().Format("2006-01-02 15:04:05")
// 	responseWithJSON(w, http.StatusOK, f)
// }

// // DeleteFilm remove one roll film by "ID".
// func DeleteFilm(w http.ResponseWriter, r *http.Request) {
// 	filmID := mux.Vars(r)["filmID"]

// 	for i, singleFIlm := range films {
// 		if singleFIlm.ID == filmID {
// 			films = append(films[:i], films[i+1:]...)
// 			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", filmID)
// 		}
// 	}
// }

func getFilmByID(ID string) (f *Film, err error) {
	// for _, singleFilm := range films {
	// 	if singleFilm.ID == ID {
	// 		return singleFilm, nil
	// 	}
	// }
	return f, errors.New("Film ID not found")
}
