package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"-"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func (m Movie) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	type MovieAlias Movie

	// Anonymous struct
	aux := struct {
		MovieAlias
		Runtime string `json:"runtime,omitempty,string"`
	}{
		MovieAlias: MovieAlias(m),
		Runtime:    runtime,
	}

	return json.Marshal(aux)
}
