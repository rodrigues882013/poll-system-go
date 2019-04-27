package commons

import (
	"encoding/json"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"log"
	"net/http"
	"time"
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
		log.Println("%s: %s", msg, err)
	}
}

func InSlice(a models.Nominate, list []models.Nominate) bool {
	for _, b := range list {
		if b.Id == a.Id {
			return true
		}
	}
	return false
}

func IsOutOfPeriod(t time.Time, duration int64) bool {
	return time.Now().After(t.Add(time.Duration(duration * 3600000000000)))
}
