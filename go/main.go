package main

import (
	"fmt"
	"github.com/neorel/test-go/services"
)

func main() {
	nbLaps := services.GetNbLaps("2011", "5")

	for lapId := int8(1); lapId <= *nbLaps; lapId++ {
		race := services.GetLap("2011", "5", lapId)

		fmt.Println(*race);
		lap := *race.Laps[0]
		fmt.Println(lap)
		for _, timing := range lap.Timings {
			fmt.Println(*timing)
		}
	}
}