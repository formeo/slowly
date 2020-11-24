package views

import (
	"encoding/json"
	"github.com/formeo/slowly/pkg/models"
	"net/http"
	"time"
)

func ReturnPostWithParamTimeout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var postTimeout models.PostTimeout
	_ = json.NewDecoder(r.Body).Decode(&postTimeout)
	time.Sleep(time.Duration(postTimeout.Timeout) * time.Millisecond)
	response := models.Response{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
