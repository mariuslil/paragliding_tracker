package main

import (
	"math/rand"
	"reflect"
)


type Track struct {
	H_date string `json:"h_date"`
	Pilot string `json:"pilot"`
	Glider string `json:"glider"`
	Glider_id string `json:"glider_id"`
	Track_length float64 `json:"track_length"`
	Track_src_url string `json:"track_src_url"`
}

type TrackDB struct {
	tracks map[int]Track
}

func (db *TrackDB) Init (){
	db.tracks = make(map[int]Track)
}

func (db *TrackDB) fieldExist (id int, field string) string{
	track := reflect.ValueOf(db.tracks[id])
	value := reflect.Indirect(track).FieldByName(field)

	return string(value.String())
}

func (db *TrackDB) Add(t Track) int{
	id := rand.Int()

	_, ok := db.Get(id)

	for ok{
		id = rand.Int()
		_, ok = db.Get(id)
	}

	db.tracks[id] = t
	return id
}

func (db *TrackDB) Get(i int) (Track, bool) {
	t, ok := db.tracks[i]
	return t, ok
}