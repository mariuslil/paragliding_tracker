package main

import "time"

type Ticker struct {
	T_latest time.Time `json: "t_latest"`
	T_start time.Time `json: "t_start"`
	T_stop time.Time `json: "t_stop"`
	Tracks map[int]Track `json: "tracks"`
	Processing int `json: "processing"`
}
