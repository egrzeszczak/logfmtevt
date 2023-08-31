package logfmtevt

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Pair struct {
	Key   string
	Value string
}

type Event struct {
	Time   time.Time
	ID     uuid.UUID
	Values []Pair
}

// Create a new event object
func New(values []Pair) Event {
	// Time will be logged in UTC
	tm := time.Now().UTC()
	// Unique ID will is UUID
	id := uuid.New()

	return Event{
		Time:   tm,
		ID:     id,
		Values: values,
	}
}

// Add key-value pairs, like evttype="security"
func (e Event) Add(values []Pair) Event {
	e.Values = append(e.Values, values...)
	return e
}

// Convert the event into a logfmt string
func (e Event) String() string {
	// Create the base for the event containing time in RFC 3339 / ISO 8601 format, and a unique identifier
	str := fmt.Sprintf("time=%s id=%s", e.Time.Format("2006-01-02T15:04:05.99999Z07:00"), e.ID.String())

	// Add dynamic values
	for _, pair := range e.Values {
		str = str + " " + fmt.Sprintf("%s=\"%v\"", pair.Key, pair.Value)
	}

	return str
}
