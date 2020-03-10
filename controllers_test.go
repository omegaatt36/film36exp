package main

import (
	"film36exp/routes"
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetDefault(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		// trun on the debug mode.
		SetDebug(true).
		Run(routes.NewRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "film36exp", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func BenchmarkGetDefault(b *testing.B) {
	r := gofight.New()

	for i := 0; i < b.N; i++ {
		r.GET("/").
			// trun on the debug mode.
			SetDebug(true).
			Run(routes.NewRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			})
	}
}

// func TestCreateOneFilm(t *testing.T) {
// 	r := gofight.New()
// 	testVendor := "Kodak"
// 	testProduction := "Protra"
// 	testISO := "400"
// 	testFilmSize := "135"
// 	testFilmType := "negtive"
// 	r.POST("/Createfilm").
// 		// trun on the debug mode.
// 		SetDebug(true).
// 		SetJSON(gofight.D{
// 			"Vendor":     testVendor,
// 			"Production": testProduction,
// 			"ISO":        testISO,
// 			"FilmSize":   testFilmSize,
// 			"FilmType":   testFilmType,
// 		}).
// 		Run(routes.NewRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
// 			data := []byte(r.Body.String())

// 			vender, _ := jsonparser.GetString(data, "Vendor")
// 			production, _ := jsonparser.GetString(data, "Production")
// 			iso, _ := jsonparser.GetString(data, "ISO")
// 			filmSize, _ := jsonparser.GetString(data, "FilmSize")
// 			filmType, _ := jsonparser.GetString(data, "FilmType")

// 			assert.Equal(t, testVendor, vender)
// 			assert.Equal(t, testProduction, production)
// 			assert.Equal(t, testISO, iso)
// 			assert.Equal(t, testFilmSize, filmSize)
// 			assert.Equal(t, testFilmType, filmType)
// 			assert.Equal(t, http.StatusCreated, r.Code)
// 		})
// }
