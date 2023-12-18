package models

import "time"

type Race struct {
	Season string `json:"season"`
	Name string `json:"name"`
	DateTime time.Time `json:"time"`
	NbLaps *int8 `json:"nbLaps"`
	Laps []*Lap `json:"laps"`
}

type Lap struct {
	Number int8 `json:"number"`
	Timings []*Timing `json:"timings"`
}

type Timing struct {
	DriverId string `json:"driverId"`
	Position int8 `json:"position"`
	Time float32 `json:"time"`
}