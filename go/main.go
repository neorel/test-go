package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/neorel/test-go/services"
)

func main() {
	services.InitRedis()

	http.HandleFunc("/", getRace)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func getRace(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	season := r.URL.Query().Get("season")
	raceNumber := r.URL.Query().Get("number")

	if r.URL.Path != "/race" || season == "" || raceNumber == "" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	raceResults := services.GetRaceResults(season, raceNumber)
	jsonRaceResults, _ := json.Marshal(raceResults)
	w.Write(jsonRaceResults)
	fmt.Println(string(jsonRaceResults))

	go services.RunLaps(season, raceNumber, raceResults.NbLaps)
}