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

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonResponse)

	if err != nil {
		panic(err)
	}

}
