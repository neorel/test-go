package services

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/neorel/test-go/models"
)

func fetch(url string) *[]byte {
    res, err := http.Get(url)
	if err != nil {
        fmt.Println(err)
        return nil
    }

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
        fmt.Println(err)
        return nil
    }

    return &data
}

func GetRaceResults(season string, raceNumber string) *models.Race {
    data := fetch(fmt.Sprintf("https://ergast.com/api/f1/%v/%v/results.json", season, raceNumber))

    ergastResponse := ParseErgastResponse(*data)
    return ConvertToRace(ergastResponse);
}

func GetLap(season string, gpNumber string, lapNumber int8) *models.Race {
	data := fetch(fmt.Sprintf("https://ergast.com/api/f1/%v/%v/laps/%v.json", season, gpNumber, lapNumber))
	
	ergastResponse := ParseErgastResponse(*data)
	return ConvertToRace(ergastResponse)
}
