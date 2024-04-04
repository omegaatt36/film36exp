package main

import (
	"net/http"
	"os"
	"reflect"
	"testing"

	"github.com/omegaatt36/film36exp/routes"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestClientOptions(t *testing.T) {
	os.Setenv("profile", "prod")
	want := options.Client().ApplyURI("mongodb://db:27017")
	got := clientOptions()
	if !reflect.DeepEqual(got.GetURI(), want.GetURI()) {
		t.Errorf("clientOptions() got = %v, want %v", got, want)
	}
}

func TestClientOptionsNonProdProfile(t *testing.T) {
	os.Setenv("profile", "dev")
	want := options.Client().ApplyURI("mongodb://localhost:27017")

	got := clientOptions()

	if !reflect.DeepEqual(got.GetURI(), want.GetURI()) {
		t.Errorf("clientOptions() got = %v, want %v", got, want)
	}
}

func TestGetDefault(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		// trun on the debug mode.
		SetDebug(true).
		Run(routes.NewRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "github.com/omegaatt36/film36exp", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
