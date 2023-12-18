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

	switch r.Method {
	case "GET":
		race := services.GetRaceResults(season, raceNumber)


		jsonRace, _ := json.Marshal(race)
		fmt.Fprintf(w, string(jsonRace))

		for lapId := int8(1); lapId <= *race.NbLaps; lapId++ {
			race := services.GetLap(season, raceNumber, lapId)

			fmt.Println(*race);
			lap := (*race).Laps[0]
			fmt.Println(lap)
			for _, timing := range lap.Timings {
				fmt.Println(*timing)
			}

			services.SendLap(lap)
		}
	}
}