package models

import "time"

type Race struct {
	Season string `json:"season"`
	Name string `json:"name"`
	DateTime time.Time `json:"time"`
}

type RaceResults struct {
	Race
	NbLaps int8 `json:"nbLaps"`
}

type RaceLap struct {
	Race
	LapNumber int8 `json:"lapNumber"`
	Timings []*Timing `json:"timings"`
}

type Timing struct {
	DriverId string `json:"driverId"`
	Position int8 `json:"position"`
	Time float32 `json:"time"`
}