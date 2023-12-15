package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/neorel/test-go/model"
	"strconv"
	"strings"
	"time"
)

func main() {
	res, err := http.Get("https://ergast.com/api/f1/2011/5/laps/1.json")
	if err != nil {
        fmt.Println(err)
        return
    }

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
        fmt.Println(err)
        return
    }

	fmt.Println(string(data));

	race := parseErgastData(data)

	fmt.Println(race);
	lap := *race.Laps[0]
	fmt.Println(lap)
	for _, timing := range lap.Timings {
		fmt.Println(*timing)
	}
}

func parseErgastData(data []byte) model.Race {
	var ergastResponse model.ErgastResponse
	json.Unmarshal(data, &ergastResponse)

	ergastRace := ergastResponse.MRData.RaceTable.Races[0]

	dateTime, _ := time.Parse("2008-03-16 04:30:00Z", ergastRace.Date + " " + ergastRace.Time) 

	race := model.Race{
		Season: ergastRace.Season,
		Name: ergastRace.RaceName,
		DateTime: dateTime,
		Laps: []*model.Lap{},
	}

	number, _ := strconv.ParseInt(ergastRace.Laps[0].Number, 10, 8)

	lap := model.Lap {
		Number: int8(number),
		Timings: []*model.Timing{},
	}
	race.Laps = append(race.Laps, &lap)

	for _, ergastTiming := range ergastRace.Laps[0].Timings {
		splitTime := strings.Split(ergastTiming.Time, ":")
		minutes, _ := strconv.ParseFloat(splitTime[0], 64)
		seconds, _ := strconv.ParseFloat(splitTime[1], 64)

		position, _ := strconv.ParseInt(ergastTiming.Position, 10, 8)

		timning := model.Timing{
			DriverId: ergastTiming.DriverId,
			Position: int8(position),
			Time: float32(minutes * 60 + seconds),
		}
		lap.Timings = append(lap.Timings, &timning)
	}

	return race
}