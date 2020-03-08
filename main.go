package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Resp struct {
	D string `json:"date"`
}

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	pattern := `\d{4}-\d{2}-\d{2}T\d{2}:\d{2}`
	re := regexp.MustCompile(pattern)
	d:=re.FindString(string(body))
	dt, err := time.Parse("2006-01-02T15:04",d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	dt = dt.Add(time.Hour*12)
	d = dt.Format("2006-01-02T15:04")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Resp{D:d})

}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", AllMoviesEndPoint).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
