package utility

import (
	"encoding/json"
	"net/http"
)

const (
	ResSuccess = "success"
	ResFailed  = "fail"
)

type Response struct {
	Message string      `json:"msg"`
	Result  string      `json:"res"`
	Data    interface{} `json:"data"`
}

func ResponseWithJSON(response http.ResponseWriter, code int, payload interface{}) {
	result, _ := json.Marshal(payload)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write(result)
}
