package services

import (
	"fmt"
	"time"
)

func RunLaps(season string, raceNumber string, nbLaps int8) {
	fmt.Println("Start race "+season+"/"+raceNumber)

	for lapIdx := int8(1); lapIdx <= nbLaps; lapIdx++ {
		time.Sleep(5 * time.Second)
		raceLap := GetLap(season, raceNumber, lapIdx)
		SendRaceLap(raceLap)
		fmt.Println("lap: ", lapIdx)
	}
}