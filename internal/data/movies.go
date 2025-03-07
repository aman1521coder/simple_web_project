package data

import "time"

type Movie struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Runtime   Runtime     `json:"runtime,omitempty,string"`

	Genere  []string `json:"geners, omitempty"`
	Year    int32    `json:"years,omitempty"`
	Version int32    `json:"version"`
}
