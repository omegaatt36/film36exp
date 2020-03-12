package controllers

import (
	"encoding/json"
	"film36exp/db"
	"film36exp/model"
	"net/http"

	"github.com/gorilla/mux"
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

// UpdatePic modify one pick by "filmID" and "picID".
func UpdatePic(w http.ResponseWriter, r *http.Request) {
	// filmID := mux.Vars(r)["filmID"]
	// picID := mux.Vars(r)["picID"]
	// var updatedPic model.Pic
	// reqBody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	// }

	// f, err := getFilmByID(filmID)
	// if err != nil {
	// 	responseWithJSON(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// p, err := getPicByID(picID, f.Pics)
	// if err != nil {
	// 	responseWithJSON(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// json.Unmarshal(reqBody, &updatedPic)

	// p.Camera = updatedPic.Camera
	// p.Lens = updatedPic.Lens
	// p.Aperture = updatedPic.Aperture
	// p.Shutter = updatedPic.Shutter
	// p.Notes = updatedPic.Notes
	// f.UptateAt = time.Now().Format("2006-01-02 15:04:05")
	// responseWithJSON(w, http.StatusOK, p)
}

// func getPicByID(ID string, pics []*model.Pic) (p *model.Pic, err error) {
// 	for _, singlePic := range pics {
// 		if singlePic.ID == ID {
// 			return singlePic, nil
// 		}
// 	}
// 	return p, errors.New("model.Pic ID not found")
// }
