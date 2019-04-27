package commons

import (
	"encoding/json"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/domain/models"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, status int, data interface{}){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(&data)

	if err != nil {
		log.Println(err)
	}
}


func BindJSON(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return false
	}
	return true
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}
