package services

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/neorel/test-go/models"
	"strconv"
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

func GetNbLaps(season string, gpNumber string) *int8 {
	data := fetch(fmt.Sprintf("https://ergast.com/api/f1/%v/%v/results.json", season, gpNumber))

    ergastResponse := ParseErgastResponse(*data)
    nbLaps, _ := strconv.ParseInt((*ergastResponse).MRData.RaceTable.Races[0].Results[0].NbLaps, 10, 8)
    nbLaps8 := int8(nbLaps)
    
    return &nbLaps8
}

func GetLap(season string, gpNumber string, lapNumber int8) *models.Race {
	data := fetch(fmt.Sprintf("https://ergast.com/api/f1/%v/%v/laps/%v.json", season, gpNumber, lapNumber))
	
	ergastResponse := ParseErgastResponse(*data)
	return ConvertToRace(ergastResponse)
}
