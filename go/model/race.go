package model

import "time"

type Race struct {
	Season string
	Name string
	DateTime time.Time
	Laps []*Lap
}

type Lap struct {
	Number int8
	Timings []*Timing
}

type Timing struct {
	DriverId string
	Position int8
	Time float32
}