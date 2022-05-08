package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("err : record not found")
	ErrEditConflict = errors.New("err : edit conflict")
)

type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		GetAll(title string,genres []string,f Filters) ([]*Movie,error)
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{
			DB: db,
		},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
