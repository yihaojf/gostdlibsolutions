package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

var dataFile = path.Join("..", "data", "proverbs.json")

func main() {
	proverbs, err := loadProverbs(dataFile)
	if err != nil {
		log.Fatalln(err)
	}

	h := newHandler(proverbs)
	r := newRouter(h)

	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

// newRouter returns a router to expose the following endpoints
// POST /proverbs (create proverb)
// GET /proverbs (get all proverbs)
// GET /proverbs/{id} (get a specific proverb)
// PUT /proverbs/{id} (update a specific proverb)
// DELETE /proverbs/{id} (delete a specific proverb)
func newRouter(h *handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/proverbs", h.createProverb).Methods("POST")
	r.HandleFunc("/proverbs", h.getProverbs).Methods("GET")
	r.HandleFunc("/proverbs/{id:[0-9]+}", h.getProverb).Methods("GET")
	r.HandleFunc("/proverbs/{id:[0-9]+}", h.updateProverb).Methods("PUT")
	r.HandleFunc("/proverbs/{id:[0-9]+}", h.deleteProverb).Methods("DELETE")
	return r
}

func loadProverbs(dataFile string) ([]Proverb, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	var proverbs []Proverb
	if err := json.NewDecoder(file).Decode(&proverbs); err != nil {
		return nil, err
	}
	return proverbs, nil
}
