package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func JSONResponse(res http.ResponseWriter, data interface{}, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(res).Encode(data)
	}
}
