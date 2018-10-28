package main

type Ticker struct {
	T_latest int `json: "t_latest"`
	T_start int `json: "t_start"`
	T_stop int `json: "t_stop"`
	Tracks int `json: "tracks"`
	Processing int `json: "processing"`
}
