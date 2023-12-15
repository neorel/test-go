package model

type ErgastResponse struct {
	MRData ErgastMRData `json:"MRData"`
}

type ErgastMRData struct {
	RaceTable ErgastRaceTable `json:"RaceTable"`
}

type ErgastRaceTable struct {
	Races []ErgastRace `json:"Races"`
}

type ErgastRace struct {
	Season string `json:"season"`
	RaceName string `json:"raceName"`
	Circuit ErgastCircuit `json:"Circuit"`
	Date string `json:"date"`
	Time string `json:"time"`
	Laps []ErgastLap `json:"Laps"`
}

type ErgastCircuit struct {
	Id string `json:"circuitId"`
	Name string `json:"circuitName"`
	Location ErgastLocation `json:"Location"`
}

type ErgastLocation struct {
	Lat string `json:"lat"`
	Lng string `json:"long"`
	Locality string `json:"locality"`
	Country string `json:"country"`
}

type ErgastLap struct {
	Number string `json:"number"`
	Timings []ErgastTiming `json:"Timings"`
}

type ErgastTiming struct {
	DriverId string `json:"driverId"`
	Position string `json:"position"`
	Time string `json:"time"`
}