// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/proxy/ent/requestlog"
)

// RequestLogCreate is the builder for creating a RequestLog entity.
type RequestLogCreate struct {
	config
	mutation *RequestLogMutation
	hooks    []Hook
}

// SetTime sets the "time" field.
func (rlc *RequestLogCreate) SetTime(t time.Time) *RequestLogCreate {
	rlc.mutation.SetTime(t)
	return rlc
}

// SetHeaders sets the "headers" field.
func (rlc *RequestLogCreate) SetHeaders(s string) *RequestLogCreate {
	rlc.mutation.SetHeaders(s)
	return rlc
}

// SetBody sets the "body" field.
func (rlc *RequestLogCreate) SetBody(b []byte) *RequestLogCreate {
	rlc.mutation.SetBody(b)
	return rlc
}

// SetFlag sets the "flag" field.
func (rlc *RequestLogCreate) SetFlag(s string) *RequestLogCreate {
	rlc.mutation.SetFlag(s)
	return rlc
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (rlc *RequestLogCreate) SetNillableFlag(s *string) *RequestLogCreate {
	if s != nil {
		rlc.SetFlag(*s)
	}
	return rlc
}

// Mutation returns the RequestLogMutation object of the builder.
func (rlc *RequestLogCreate) Mutation() *RequestLogMutation {
	return rlc.mutation
}

// Save creates the RequestLog in the database.
func (rlc *RequestLogCreate) Save(ctx context.Context) (*RequestLog, error) {
	var (
		err  error
		node *RequestLog
	)
	rlc.defaults()
	if len(rlc.hooks) == 0 {
		if err = rlc.check(); err != nil {
			return nil, err
		}
		node, err = rlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequestLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rlc.check(); err != nil {
				return nil, err
			}
			rlc.mutation = mutation
			if node, err = rlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rlc.hooks) - 1; i >= 0; i-- {
			if rlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rlc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rlc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RequestLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RequestLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rlc *RequestLogCreate) SaveX(ctx context.Context) *RequestLog {
	v, err := rlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rlc *RequestLogCreate) Exec(ctx context.Context) error {
	_, err := rlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlc *RequestLogCreate) ExecX(ctx context.Context) {
	if err := rlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rlc *RequestLogCreate) defaults() {
	if _, ok := rlc.mutation.Flag(); !ok {
		v := requestlog.DefaultFlag
		rlc.mutation.SetFlag(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rlc *RequestLogCreate) check() error {
	if _, ok := rlc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "RequestLog.time"`)}
	}
	if _, ok := rlc.mutation.Headers(); !ok {
		return &ValidationError{Name: "headers", err: errors.New(`ent: missing required field "RequestLog.headers"`)}
	}
	if _, ok := rlc.mutation.Flag(); !ok {
		return &ValidationError{Name: "flag", err: errors.New(`ent: missing required field "RequestLog.flag"`)}
	}
	return nil
}

func (rlc *RequestLogCreate) sqlSave(ctx context.Context) (*RequestLog, error) {
	_node, _spec := rlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rlc *RequestLogCreate) createSpec() (*RequestLog, *sqlgraph.CreateSpec) {
	var (
		_node = &RequestLog{config: rlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: requestlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: requestlog.FieldID,
			},
		}
	)
	if value, ok := rlc.mutation.Time(); ok {
		_spec.SetField(requestlog.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	if value, ok := rlc.mutation.Headers(); ok {
		_spec.SetField(requestlog.FieldHeaders, field.TypeString, value)
		_node.Headers = value
	}
	if value, ok := rlc.mutation.Body(); ok {
		_spec.SetField(requestlog.FieldBody, field.TypeBytes, value)
		_node.Body = value
	}
	if value, ok := rlc.mutation.Flag(); ok {
		_spec.SetField(requestlog.FieldFlag, field.TypeString, value)
		_node.Flag = value
	}
	return _node, _spec
}

// RequestLogCreateBulk is the builder for creating many RequestLog entities in bulk.
type RequestLogCreateBulk struct {
	config
	builders []*RequestLogCreate
}

// Save creates the RequestLog entities in the database.
func (rlcb *RequestLogCreateBulk) Save(ctx context.Context) ([]*RequestLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rlcb.builders))
	nodes := make([]*RequestLog, len(rlcb.builders))
	mutators := make([]Mutator, len(rlcb.builders))
	for i := range rlcb.builders {
		func(i int, root context.Context) {
			builder := rlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RequestLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rlcb *RequestLogCreateBulk) SaveX(ctx context.Context) []*RequestLog {
	v, err := rlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rlcb *RequestLogCreateBulk) Exec(ctx context.Context) error {
	_, err := rlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rlcb *RequestLogCreateBulk) ExecX(ctx context.Context) {
	if err := rlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
