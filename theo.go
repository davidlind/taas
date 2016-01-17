package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
	"strconv"
	//"fmt"
)

type Insult struct {
	ID int `json:"id"`
	Text string `json:"text"`
}

func getInsults() []Insult {
	raw, err := ioutil.ReadFile("./insults.json")
	if err != nil {
		panic(err)
	}
	var insults []Insult
	err = json.Unmarshal(raw, &insults)
	return insults
}

func main() {
	//insults := make([]Insult, 2)
	//insults[0] = Insult{ID: 0, Text: "Sorry, no."}
	//insults[1] = Insult{ID: 1, Text: "Fuck off."}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/{id}", Index)

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
	insults := getInsults()
	vars := mux.Vars(r)
	id := vars["id"]
	id_int, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(insults[id_int])
}