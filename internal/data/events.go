package data

import (
	"context"
	"database/sql"
	"greenlight/internal/validator"
	"time"

	"github.com/lib/pq"
)

type Event struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Venue       string    `json:"venue"`
	Tags        []string  `json:"genres,omitempty"`
	Cover       string    `json:"images,omitempty"`
	Version     int32     `json:"version"`
}

func ValidateEvent(v *validator.Validator, event *Event) {
	v.Check(event.Title != "", "title", "title must be provided")
	v.Check(len(event.Title) <= 500, "title", "title must not exceed 500 bytes")

	// todo : check that the end time is not in the past, start time can be in the past
	//v.Check(event.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(event.Cover != "", "images", "must container atleast 1 cover image")

	v.Check(validator.Unique(event.Tags), "tags", "must not contain duplicate values")
}

type EventModel struct {
	DB *sql.DB
}

func (m EventModel) Insert(event *Event) (*Event, error) {
	query := `
		INSERT INTO events(start_time,end_time,title,description,tags,cover)
		VALUES($1,$2,$3,$4,$5,$6)
		RETURNING id,created_at,version
	`

	args := []interface{}{event.StartTime, event.EndTime, event.Title, event.Description, pq.Array(event.Tags), event.Cover}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&event.ID, &event.CreatedAt, &event.Version)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (m EventModel) GetAll(title string, tags []string, filters Filters) ([]*Event, Metadata, error) {

	// query := fmt.Sprintf(`
	// 	SELECT count(*) OVER(),id,created_at,start_time,end_time,title,description,venue,tags,cover,version
	// 	FROM events
	// 	WHERE (to_tsvector('simple',title) @@ plainto_tsquery('simple',$1) OR $1 = '')
	// 	AND (genres @> $2 OR $2 = '{}')
	// 	ORDER BY %s %s,id ASC
	// 	LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())
	panic("implement me")

}

func (m EventModel) Get(id int64) (*Event, error) {
	return nil, nil
}

func (m EventModel) Update(id int64, event *Event) (*Event, error) {
	return nil, nil
}
