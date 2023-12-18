package services

import (
	"encoding/json"
	"github.com/neorel/test-go/models"
	"strconv"
	"strings"
	"time"
)

func ParseErgastResponse(data []byte) *models.ErgastResponse {
	var ergastResponse models.ErgastResponse
	json.Unmarshal(data, &ergastResponse)

	return &ergastResponse
}

func convertToRace(ergastResponse *models.ErgastResponse) (*models.ErgastRace, *models.Race) {
	ergastRace := (*ergastResponse).MRData.RaceTable.Races[0]

	dateTime, _ := time.Parse("2006-01-02 15:04:05Z", ergastRace.Date + " " + ergastRace.Time)

	race := models.Race{
		Season: ergastRace.Season,
		Name: ergastRace.RaceName,
		DateTime: dateTime,
	}

	return &ergastRace, &race
}

func ConvertToRaceResults(ergastResponse *models.ErgastResponse) *models.RaceResults {
	ergastRace, race := convertToRace(ergastResponse)

	nbLaps, _ := strconv.ParseInt((*ergastRace).Results[0].NbLaps, 10, 8)

	raceResults := models.RaceResults {
		Race: *race,
		NbLaps: int8(nbLaps),
	}

	return &raceResults
}

func ConvertToRaceLap(ergastResponse *models.ErgastResponse) *models.RaceLap {
	ergastRace, race := convertToRace(ergastResponse)

	ergastLap := ergastRace.Laps[0]
	lapNumber, _ := strconv.ParseInt(ergastLap.Number, 10, 8)

	raceLap := models.RaceLap{
		Race: *race,
		LapNumber: int8(lapNumber),
		Timings: []*models.Timing{},
	}

	for _, ergastTiming := range ergastLap.Timings {
		splitTime := strings.Split(ergastTiming.Time, ":")
		minutes, _ := strconv.ParseFloat(splitTime[0], 64)
		seconds, _ := strconv.ParseFloat(splitTime[1], 64)

		position, _ := strconv.ParseInt(ergastTiming.Position, 10, 8)

		timning := models.Timing{
			DriverId: ergastTiming.DriverId,
			Position: int8(position),
			Time: float32(minutes * 60 + seconds),
		}
		raceLap.Timings = append(raceLap.Timings, &timning)
	}

	return &raceLap
}