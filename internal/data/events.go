package data

import (
	"context"
	"database/sql"
	"fmt"
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
	Location    string    `json:"location"`
	Tags        []string  `json:"genres,omitempty"`
	Cover       string    `json:"images,omitempty"`
	Version     int32     `json:"version"`
}

func ValidateEvent(v *validator.Validator, event *Event) {
	v.Check(event.Title != "", "title", "title must be provided")
	v.Check(len(event.Title) <= 500, "title", "title must not exceed 500 bytes")

	// todo : check that the end time is not in the past, start time can be in the past
	v.Check(event.EndTime.After(time.Now()), "start_time", "must not be in the past")
	v.Check(event.StartTime.Before(event.EndTime), "start_time", "must be before end time")

	v.Check(event.Cover != "", "images", "must container atleast 1 cover image")

	v.Check(validator.Unique(event.Tags), "tags", "must not contain duplicate values")
}

type EventModel struct {
	DB *sql.DB
}

func (m EventModel) Insert(event *Event) (*Event, error) {
	query := `
		INSERT INTO events(start_time,end_time,title,description,venue,location,tags,cover)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)
		RETURNING id,created_at,version
	`

	args := []interface{}{event.StartTime, event.EndTime, event.Title, event.Description, event.Venue, event.Location, pq.Array(event.Tags), event.Cover}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&event.ID, &event.CreatedAt, &event.Version)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (m EventModel) GetAll(location, title, description string, tags []string, filters Filters) ([]*Event, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT COUNT(*) OVER() 
			id,created_at,start_time,end_time,title,description,venue,location,tags,cover,version
		FROM 
			events
		WHERE 
			(to_tsvector('simple',title) @@ plainto_tsquery('simple',$1) OR $1 = '')
		AND
			(to_tsvector('english',description) @@ plainto_tsquery('english',$2) OR $2 = '')
		AND 
			(tags @> $3 OR $3 = '{}')
		AND 
			location = $4
		ORDER BY 
			%s %s,id ASC
		LIMIT 
			$5
		OFFSET 
			$6
	`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, description, pq.Array(tags), location, filters.limit(), filters.offset()}
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	var events []*Event

	var totalRecords int

	for rows.Next() {
		var event Event

		err := rows.Scan(&totalRecords, &event.ID, &event.StartTime, &event.EndTime,
			&event.Title, &event.Description, &event.Venue, &event.Location, pq.Array(&event.Tags), &event.Cover, &event.Version)
		if err != nil {
			return nil, Metadata{}, err
		}

		events = append(events, &event)
	}

	metadata := calculateMetaData(totalRecords, filters.Page, filters.PageSize)

	return events, metadata, nil
}

func (m EventModel) Get(id int64) (*Event, error) {
	var event Event

	stmt := `
		SELECT 
			id,created_at,start_time,end_time,title,description,venue,tags,cover,version
		FROM 
			movies
		WHERE 
			id = $1
		`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(
		&event.ID, &event.CreatedAt, &event.StartTime, &event.EndTime, &event.Title,
		&event.Description, &event.Venue, &event.Cover, pq.Array(&event.Tags), &event.Version,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (m EventModel) Update(id int64, event *Event) (*Event, error) {
	return nil, nil
}
