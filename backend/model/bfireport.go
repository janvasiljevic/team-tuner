// Code generated by ent, DO NOT EDIT.

package model

import (
	"encoding/json"
	"fmt"
	"jv/team-tone-tuner/model/bfireport"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/schema"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// BfiReport is the model entity for the BfiReport schema.
type BfiReport struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// The time when the record was created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The time when the record was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Conscientiousness holds the value of the "conscientiousness" field.
	Conscientiousness schema.BfiReportItem `json:"conscientiousness,omitempty"`
	// Extraversion holds the value of the "extraversion" field.
	Extraversion schema.BfiReportItem `json:"extraversion,omitempty"`
	// Agreeableness holds the value of the "agreeableness" field.
	Agreeableness schema.BfiReportItem `json:"agreeableness,omitempty"`
	// Neuroticism holds the value of the "neuroticism" field.
	Neuroticism schema.BfiReportItem `json:"neuroticism,omitempty"`
	// Openness holds the value of the "openness" field.
	Openness schema.BfiReportItem `json:"openness,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BfiReportQuery when eager-loading is set.
	Edges        BfiReportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BfiReportEdges holds the relations/edges for other nodes in the graph.
type BfiReportEdges struct {
	// Student holds the value of the student edge.
	Student *User `json:"student,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StudentOrErr returns the Student value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BfiReportEdges) StudentOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Student == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Student, nil
	}
	return nil, &NotLoadedError{edge: "student"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BfiReport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bfireport.FieldConscientiousness, bfireport.FieldExtraversion, bfireport.FieldAgreeableness, bfireport.FieldNeuroticism, bfireport.FieldOpenness:
			values[i] = new([]byte)
		case bfireport.FieldCreatedAt, bfireport.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case bfireport.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BfiReport fields.
func (br *BfiReport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bfireport.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				br.ID = *value
			}
		case bfireport.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				br.CreatedAt = value.Time
			}
		case bfireport.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				br.UpdatedAt = value.Time
			}
		case bfireport.FieldConscientiousness:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field conscientiousness", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.Conscientiousness); err != nil {
					return fmt.Errorf("unmarshal field conscientiousness: %w", err)
				}
			}
		case bfireport.FieldExtraversion:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field extraversion", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.Extraversion); err != nil {
					return fmt.Errorf("unmarshal field extraversion: %w", err)
				}
			}
		case bfireport.FieldAgreeableness:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field agreeableness", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.Agreeableness); err != nil {
					return fmt.Errorf("unmarshal field agreeableness: %w", err)
				}
			}
		case bfireport.FieldNeuroticism:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field neuroticism", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.Neuroticism); err != nil {
					return fmt.Errorf("unmarshal field neuroticism: %w", err)
				}
			}
		case bfireport.FieldOpenness:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field openness", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &br.Openness); err != nil {
					return fmt.Errorf("unmarshal field openness: %w", err)
				}
			}
		default:
			br.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BfiReport.
// This includes values selected through modifiers, order, etc.
func (br *BfiReport) Value(name string) (ent.Value, error) {
	return br.selectValues.Get(name)
}

// QueryStudent queries the "student" edge of the BfiReport entity.
func (br *BfiReport) QueryStudent() *UserQuery {
	return NewBfiReportClient(br.config).QueryStudent(br)
}

// Update returns a builder for updating this BfiReport.
// Note that you need to call BfiReport.Unwrap() before calling this method if this BfiReport
// was returned from a transaction, and the transaction was committed or rolled back.
func (br *BfiReport) Update() *BfiReportUpdateOne {
	return NewBfiReportClient(br.config).UpdateOne(br)
}

// Unwrap unwraps the BfiReport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (br *BfiReport) Unwrap() *BfiReport {
	_tx, ok := br.config.driver.(*txDriver)
	if !ok {
		panic("model: BfiReport is not a transactional entity")
	}
	br.config.driver = _tx.drv
	return br
}

// String implements the fmt.Stringer.
func (br *BfiReport) String() string {
	var builder strings.Builder
	builder.WriteString("BfiReport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", br.ID))
	builder.WriteString("created_at=")
	builder.WriteString(br.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(br.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("conscientiousness=")
	builder.WriteString(fmt.Sprintf("%v", br.Conscientiousness))
	builder.WriteString(", ")
	builder.WriteString("extraversion=")
	builder.WriteString(fmt.Sprintf("%v", br.Extraversion))
	builder.WriteString(", ")
	builder.WriteString("agreeableness=")
	builder.WriteString(fmt.Sprintf("%v", br.Agreeableness))
	builder.WriteString(", ")
	builder.WriteString("neuroticism=")
	builder.WriteString(fmt.Sprintf("%v", br.Neuroticism))
	builder.WriteString(", ")
	builder.WriteString("openness=")
	builder.WriteString(fmt.Sprintf("%v", br.Openness))
	builder.WriteByte(')')
	return builder.String()
}

// BfiReports is a parsable slice of BfiReport.
type BfiReports []*BfiReport