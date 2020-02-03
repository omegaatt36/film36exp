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

// pic is the smallest unit in this system.
type pic struct {
	ID       string `josn:"ID"`
	Camera   string `josn:"Camera"`
	Lens     string `josn:"Lens"`
	Aperture string `josn:"Aperture"`
	Shutter  string `josn:"Shutter"`
	Notes    string `josn:"Notes"`
}

// CreatePic create one pic and append into film by ID.
// if films have not filmID, will not create one film.
func CreatePic(w http.ResponseWriter, r *http.Request) {
	var newPic pic
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event id, title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newPic)
	filmID := mux.Vars(r)["filmID"]

	f, err := getFilmByID(filmID)

	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(f.Pics) == 0 {
		newPic.ID = "0"
	} else {
		idInt, err := strconv.Atoi(f.Pics[len(f.Pics)-1].ID)
		if err != nil {
			// handle error
			// log??
		}
		newPic.ID = strconv.Itoa(idInt + 1)
	}
	f.Pics = append(f.Pics, &newPic)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPic)

}

// UpdatePic modify one pick by "filmID" and "picID".
func UpdatePic(w http.ResponseWriter, r *http.Request) {
	filmID := mux.Vars(r)["filmID"]
	picID := mux.Vars(r)["picID"]
	var updatedPic pic
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	f, err := getFilmByID(filmID)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	p, err := getPicByID(picID, f.Pics)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Unmarshal(reqBody, &updatedPic)

	p.Camera = updatedPic.Camera
	p.Lens = updatedPic.Lens
	p.Aperture = updatedPic.Aperture
	p.Shutter = updatedPic.Shutter
	p.Notes = updatedPic.Notes
	f.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	responseWithJSON(w, http.StatusOK, p)
}

func getPicByID(ID string, pics []*pic) (p *pic, err error) {
	for _, singlePic := range pics {
		if singlePic.ID == ID {
			return singlePic, nil
		}
	}
	return p, errors.New("Pic ID not found")
}
