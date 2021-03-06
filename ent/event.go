// Code generated by entc, DO NOT EDIT.

package ent

import (
	"EventCatalog/ent/category"
	"EventCatalog/ent/event"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Price holds the value of the "Price" field.
	Price float32 `json:"Price,omitempty"`
	// Artist holds the value of the "Artist" field.
	Artist string `json:"Artist,omitempty"`
	// Date holds the value of the "Date" field.
	Date int64 `json:"Date,omitempty"`
	// Description holds the value of the "Description" field.
	Description string `json:"Description,omitempty"`
	// ImageUrl holds the value of the "ImageUrl" field.
	ImageUrl string `json:"ImageUrl,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges           EventEdges `json:"edges"`
	category_events *uuid.UUID
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Category `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) OwnerOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Event) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case event.FieldDate:
			values[i] = new(sql.NullInt64)
		case event.FieldName, event.FieldArtist, event.FieldDescription, event.FieldImageUrl:
			values[i] = new(sql.NullString)
		case event.FieldID:
			values[i] = new(uuid.UUID)
		case event.ForeignKeys[0]: // category_events
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case event.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case event.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field Price", values[i])
			} else if value.Valid {
				e.Price = float32(value.Float64)
			}
		case event.FieldArtist:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Artist", values[i])
			} else if value.Valid {
				e.Artist = value.String
			}
		case event.FieldDate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Date", values[i])
			} else if value.Valid {
				e.Date = value.Int64
			}
		case event.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case event.FieldImageUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ImageUrl", values[i])
			} else if value.Valid {
				e.ImageUrl = value.String
			}
		case event.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field category_events", values[i])
			} else if value.Valid {
				e.category_events = new(uuid.UUID)
				*e.category_events = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Event entity.
func (e *Event) QueryOwner() *CategoryQuery {
	return (&EventClient{config: e.config}).QueryOwner(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return (&EventClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", Name=")
	builder.WriteString(e.Name)
	builder.WriteString(", Price=")
	builder.WriteString(fmt.Sprintf("%v", e.Price))
	builder.WriteString(", Artist=")
	builder.WriteString(e.Artist)
	builder.WriteString(", Date=")
	builder.WriteString(fmt.Sprintf("%v", e.Date))
	builder.WriteString(", Description=")
	builder.WriteString(e.Description)
	builder.WriteString(", ImageUrl=")
	builder.WriteString(e.ImageUrl)
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

func (e Events) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
