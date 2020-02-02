package controllers

import (
	"encoding/json"
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
}

type allFIlms []film

var films = allFIlms{}

// func HomeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome home!")
// }

func CreateOneFilm(w http.ResponseWriter, r *http.Request) {
	var newFilm film
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event id, title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newFilm)

	if len(films) == 0 {
		newFilm.ID = "0"
	} else {
		inInt, err := strconv.Atoi(films[len(films)-1].ID)
		if err != nil {
			// handle error
			// log??
		}
		newFilm.ID = strconv.Itoa(inInt + 1)
	}
	newFilm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	newFilm.CreateAt = newFilm.UptateAt

	films = append(films, newFilm)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFilm)
}

func GetOneFilm(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["id"]

	for _, singleFIlm := range films {
		if singleFIlm.ID == filmID {
			json.NewEncoder(w).Encode(singleFIlm)
		}
	}
}

func GetAllFilms(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(films)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["id"]
	var updatedFilm film
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updatedFilm)

	for i, singleFIlm := range films {
		if singleFIlm.ID == filmID {
			singleFIlm.Production = updatedFilm.Production
			singleFIlm.Vendor = updatedFilm.Vendor
			singleFIlm.ISO = updatedFilm.ISO
			singleFIlm.ForceProcessing = updatedFilm.ForceProcessing
			singleFIlm.FilmSize = updatedFilm.FilmSize
			singleFIlm.FilmType = updatedFilm.FilmType
			singleFIlm.Description = updatedFilm.Description
			singleFIlm.UptateAt = time.Now().Format("2006-01-02 15:04:05")
			films[i] = singleFIlm
			json.NewEncoder(w).Encode(singleFIlm)
		}
	}
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["id"]

	for i, singleFIlm := range films {
		if singleFIlm.ID == filmID {
			films = append(films[:i], films[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", filmID)
		}
	}
}
