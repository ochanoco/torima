// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/ochanoco/proxy/ent/predicate"
	"github.com/ochanoco/proxy/ent/requestlog"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRequestLog = "RequestLog"
)

// RequestLogMutation represents an operation that mutates the RequestLog nodes in the graph.
type RequestLogMutation struct {
	config
	op            Op
	typ           string
	id            *int
	time          *time.Time
	headers       *string
	body          *[]byte
	flag          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*RequestLog, error)
	predicates    []predicate.RequestLog
}

var _ ent.Mutation = (*RequestLogMutation)(nil)

// requestlogOption allows management of the mutation configuration using functional options.
type requestlogOption func(*RequestLogMutation)

// newRequestLogMutation creates new mutation for the RequestLog entity.
func newRequestLogMutation(c config, op Op, opts ...requestlogOption) *RequestLogMutation {
	m := &RequestLogMutation{
		config:        c,
		op:            op,
		typ:           TypeRequestLog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRequestLogID sets the ID field of the mutation.
func withRequestLogID(id int) requestlogOption {
	return func(m *RequestLogMutation) {
		var (
			err   error
			once  sync.Once
			value *RequestLog
		)
		m.oldValue = func(ctx context.Context) (*RequestLog, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().RequestLog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRequestLog sets the old RequestLog of the mutation.
func withRequestLog(node *RequestLog) requestlogOption {
	return func(m *RequestLogMutation) {
		m.oldValue = func(context.Context) (*RequestLog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RequestLogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RequestLogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RequestLogMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RequestLogMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().RequestLog.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTime sets the "time" field.
func (m *RequestLogMutation) SetTime(t time.Time) {
	m.time = &t
}

// Time returns the value of the "time" field in the mutation.
func (m *RequestLogMutation) Time() (r time.Time, exists bool) {
	v := m.time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "time" field's value of the RequestLog entity.
// If the RequestLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestLogMutation) OldTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// ResetTime resets all changes to the "time" field.
func (m *RequestLogMutation) ResetTime() {
	m.time = nil
}

// SetHeaders sets the "headers" field.
func (m *RequestLogMutation) SetHeaders(s string) {
	m.headers = &s
}

// Headers returns the value of the "headers" field in the mutation.
func (m *RequestLogMutation) Headers() (r string, exists bool) {
	v := m.headers
	if v == nil {
		return
	}
	return *v, true
}

// OldHeaders returns the old "headers" field's value of the RequestLog entity.
// If the RequestLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestLogMutation) OldHeaders(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHeaders is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHeaders requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHeaders: %w", err)
	}
	return oldValue.Headers, nil
}

// ResetHeaders resets all changes to the "headers" field.
func (m *RequestLogMutation) ResetHeaders() {
	m.headers = nil
}

// SetBody sets the "body" field.
func (m *RequestLogMutation) SetBody(b []byte) {
	m.body = &b
}

// Body returns the value of the "body" field in the mutation.
func (m *RequestLogMutation) Body() (r []byte, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "body" field's value of the RequestLog entity.
// If the RequestLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestLogMutation) OldBody(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ClearBody clears the value of the "body" field.
func (m *RequestLogMutation) ClearBody() {
	m.body = nil
	m.clearedFields[requestlog.FieldBody] = struct{}{}
}

// BodyCleared returns if the "body" field was cleared in this mutation.
func (m *RequestLogMutation) BodyCleared() bool {
	_, ok := m.clearedFields[requestlog.FieldBody]
	return ok
}

// ResetBody resets all changes to the "body" field.
func (m *RequestLogMutation) ResetBody() {
	m.body = nil
	delete(m.clearedFields, requestlog.FieldBody)
}

// SetFlag sets the "flag" field.
func (m *RequestLogMutation) SetFlag(s string) {
	m.flag = &s
}

// Flag returns the value of the "flag" field in the mutation.
func (m *RequestLogMutation) Flag() (r string, exists bool) {
	v := m.flag
	if v == nil {
		return
	}
	return *v, true
}

// OldFlag returns the old "flag" field's value of the RequestLog entity.
// If the RequestLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RequestLogMutation) OldFlag(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFlag is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFlag requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFlag: %w", err)
	}
	return oldValue.Flag, nil
}

// ResetFlag resets all changes to the "flag" field.
func (m *RequestLogMutation) ResetFlag() {
	m.flag = nil
}

// Where appends a list predicates to the RequestLogMutation builder.
func (m *RequestLogMutation) Where(ps ...predicate.RequestLog) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RequestLogMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (RequestLog).
func (m *RequestLogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RequestLogMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.time != nil {
		fields = append(fields, requestlog.FieldTime)
	}
	if m.headers != nil {
		fields = append(fields, requestlog.FieldHeaders)
	}
	if m.body != nil {
		fields = append(fields, requestlog.FieldBody)
	}
	if m.flag != nil {
		fields = append(fields, requestlog.FieldFlag)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RequestLogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case requestlog.FieldTime:
		return m.Time()
	case requestlog.FieldHeaders:
		return m.Headers()
	case requestlog.FieldBody:
		return m.Body()
	case requestlog.FieldFlag:
		return m.Flag()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RequestLogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case requestlog.FieldTime:
		return m.OldTime(ctx)
	case requestlog.FieldHeaders:
		return m.OldHeaders(ctx)
	case requestlog.FieldBody:
		return m.OldBody(ctx)
	case requestlog.FieldFlag:
		return m.OldFlag(ctx)
	}
	return nil, fmt.Errorf("unknown RequestLog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RequestLogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case requestlog.FieldTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	case requestlog.FieldHeaders:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHeaders(v)
		return nil
	case requestlog.FieldBody:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case requestlog.FieldFlag:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFlag(v)
		return nil
	}
	return fmt.Errorf("unknown RequestLog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RequestLogMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RequestLogMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RequestLogMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown RequestLog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RequestLogMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(requestlog.FieldBody) {
		fields = append(fields, requestlog.FieldBody)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RequestLogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RequestLogMutation) ClearField(name string) error {
	switch name {
	case requestlog.FieldBody:
		m.ClearBody()
		return nil
	}
	return fmt.Errorf("unknown RequestLog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RequestLogMutation) ResetField(name string) error {
	switch name {
	case requestlog.FieldTime:
		m.ResetTime()
		return nil
	case requestlog.FieldHeaders:
		m.ResetHeaders()
		return nil
	case requestlog.FieldBody:
		m.ResetBody()
		return nil
	case requestlog.FieldFlag:
		m.ResetFlag()
		return nil
	}
	return fmt.Errorf("unknown RequestLog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RequestLogMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RequestLogMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RequestLogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RequestLogMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RequestLogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RequestLogMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RequestLogMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown RequestLog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RequestLogMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown RequestLog edge %s", name)
}
