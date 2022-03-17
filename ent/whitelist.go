// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent/project"
	"github.com/tracer-silver-bullet/tracer-silver-bullet/proxy/ent/whitelist"
)

// WhiteList is the model entity for the WhiteList schema.
type WhiteList struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WhiteListQuery when eager-loading is set.
	Edges              WhiteListEdges `json:"edges"`
	project_whitelists *int
}

// WhiteListEdges holds the relations/edges for other nodes in the graph.
type WhiteListEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Project `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WhiteListEdges) OwnerOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WhiteList) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case whitelist.FieldID:
			values[i] = new(sql.NullInt64)
		case whitelist.FieldURL:
			values[i] = new(sql.NullString)
		case whitelist.ForeignKeys[0]: // project_whitelists
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WhiteList", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WhiteList fields.
func (wl *WhiteList) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case whitelist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wl.ID = int(value.Int64)
		case whitelist.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				wl.URL = value.String
			}
		case whitelist.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_whitelists", value)
			} else if value.Valid {
				wl.project_whitelists = new(int)
				*wl.project_whitelists = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the WhiteList entity.
func (wl *WhiteList) QueryOwner() *ProjectQuery {
	return (&WhiteListClient{config: wl.config}).QueryOwner(wl)
}

// Update returns a builder for updating this WhiteList.
// Note that you need to call WhiteList.Unwrap() before calling this method if this WhiteList
// was returned from a transaction, and the transaction was committed or rolled back.
func (wl *WhiteList) Update() *WhiteListUpdateOne {
	return (&WhiteListClient{config: wl.config}).UpdateOne(wl)
}

// Unwrap unwraps the WhiteList entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wl *WhiteList) Unwrap() *WhiteList {
	tx, ok := wl.config.driver.(*txDriver)
	if !ok {
		panic("ent: WhiteList is not a transactional entity")
	}
	wl.config.driver = tx.drv
	return wl
}

// String implements the fmt.Stringer.
func (wl *WhiteList) String() string {
	var builder strings.Builder
	builder.WriteString("WhiteList(")
	builder.WriteString(fmt.Sprintf("id=%v", wl.ID))
	builder.WriteString(", url=")
	builder.WriteString(wl.URL)
	builder.WriteByte(')')
	return builder.String()
}

// WhiteLists is a parsable slice of WhiteList.
type WhiteLists []*WhiteList

func (wl WhiteLists) config(cfg config) {
	for _i := range wl {
		wl[_i].config = cfg
	}
}
