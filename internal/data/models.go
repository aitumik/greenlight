package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("err : record not found")
	ErrEditConflict   = errors.New("err : edit conflict")
)

type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		GetAll(title string, genres []string, f Filters) ([]*Movie, Metadata, error)
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
	Users interface {
		Insert(user *User) error
		GetByEmail(email string) (*User, error)
		Update(user *User) error
		GetUserForToken(scope, plainToken string) (*User, error)
	}
	Tokens interface {
		New(userID int64, ttl time.Duration, scope string) (*Token, error)
		Insert(token *Token) error
		DeleteAllForUser(scope string, userID int64) error
	}
	Permissions interface {
		GetAllForUser(userID int64) (Permissions, error)
		AddForUser(userID int64, perms ...string) error
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Movies: MockMovieModel{},
	}
}
