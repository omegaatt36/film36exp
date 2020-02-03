package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type film struct {
	ID              string
	Vendor          string `json:"Vendor"`
	Production      string `json:"Production"`
	ISO             string `json:"ISO"`
	ForceProcessing string `json:"ForceProcessing"`
	FilmSize        string `json:"FilmSize"`
	FilmType        string `json:"FilmType"`
	CreateAt        string
	UptateAt        string
	Description     string `json:"Description"`
	Pics            []*pic
}

type allFIlms []*film

var films = allFIlms{}

// CreateOneFilm create one roll film.
func CreateOneFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm film
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, "Kindly enter data with the event id, title and description only in order to update")
		return
	}

	json.Unmarshal(reqBody, &newFilm)

	if len(films) == 0 {
		newFilm.ID = "0"
	} else {
		inInt, err := strconv.Atoi(films[len(films)-1].ID)
		if err != nil {
			responseWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		newFilm.ID = strconv.Itoa(inInt + 1)
	}
	newFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	newFilm.CreateAt = newFilm.UptateAt
	films = append(films, &newFilm)

	responseWithJSON(w, http.StatusCreated, newFilm)
}

// GetOneFilm get one roll film by "ID".
func GetOneFilm(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["filmID"]

	f, err := getFilmByID(filmID)

	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, err.Error())
	} else {
		responseWithJSON(w, http.StatusOK, f)
	}
}

// GetAllFilms get all films.
func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, films)
}

// UpdateFilm modify one roll film by "ID".
func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["filmID"]
	var updatedFilm film
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	f, err := getFilmByID(filmID)

	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Unmarshal(reqBody, &updatedFilm)

	f.Production = updatedFilm.Production
	f.Vendor = updatedFilm.Vendor
	f.ISO = updatedFilm.ISO
	f.ForceProcessing = updatedFilm.ForceProcessing
	f.FilmSize = updatedFilm.FilmSize
	f.FilmType = updatedFilm.FilmType
	f.Description = updatedFilm.Description
	f.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	responseWithJSON(w, http.StatusOK, f)
}

// DeleteFilm remove one roll film by "ID".
func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["filmID"]

	for i, singleFIlm := range films {
		if singleFIlm.ID == filmID {
			films = append(films[:i], films[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", filmID)
		}
	}
}

func getFilmByID(ID string) (f *film, err error) {
	for _, singleFilm := range films {
		if singleFilm.ID == ID {
			return singleFilm, nil
		}
	}
	return f, errors.New("Film ID not found")
}
