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

func ConvertToRace(ergastResponse *models.ErgastResponse) *models.Race {
	ergastRace := (*ergastResponse).MRData.RaceTable.Races[0]

	dateTime, _ := time.Parse("2008-03-16 04:30:00Z", ergastRace.Date + " " + ergastRace.Time)
	nbLaps8 := int8(0)

	if(len(ergastRace.Results) > 0) {
		nbLaps, _ := strconv.ParseInt(ergastRace.Results[0].NbLaps, 10, 8)
		nbLaps8 = int8(nbLaps)
	}

	race := models.Race{
		Season: ergastRace.Season,
		Name: ergastRace.RaceName,
		DateTime: dateTime,
		NbLaps: &nbLaps8,
		Laps: []*models.Lap{},
	}

	if (ergastRace.Laps == nil || len(ergastRace.Laps) == 0) {
		return &race
	}

	number, _ := strconv.ParseInt(ergastRace.Laps[0].Number, 10, 8)

	lap := models.Lap {
		Number: int8(number),
		Timings: []*models.Timing{},
	}
	race.Laps = append(race.Laps, &lap)

	for _, ergastTiming := range ergastRace.Laps[0].Timings {
		splitTime := strings.Split(ergastTiming.Time, ":")
		minutes, _ := strconv.ParseFloat(splitTime[0], 64)
		seconds, _ := strconv.ParseFloat(splitTime[1], 64)

		position, _ := strconv.ParseInt(ergastTiming.Position, 10, 8)

		timning := models.Timing{
			DriverId: ergastTiming.DriverId,
			Position: int8(position),
			Time: float32(minutes * 60 + seconds),
		}
		lap.Timings = append(lap.Timings, &timning)
	}

	return &race
}