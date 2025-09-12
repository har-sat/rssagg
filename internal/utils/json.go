package utils 

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Println("Failed to marshal json response: ", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, errMsg string) {
	//server level err
	if code > 499 {
		log.Printf("Responding with %v error: %v", code, errMsg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errResponse{Error: errMsg})
}
