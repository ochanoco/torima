// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/proxy/ent/communicationlog"
)

// CommunicationLogCreate is the builder for creating a CommunicationLog entity.
type CommunicationLogCreate struct {
	config
	mutation *CommunicationLogMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (clc *CommunicationLogCreate) SetType(s string) *CommunicationLogCreate {
	clc.mutation.SetType(s)
	return clc
}

// SetTime sets the "time" field.
func (clc *CommunicationLogCreate) SetTime(t time.Time) *CommunicationLogCreate {
	clc.mutation.SetTime(t)
	return clc
}

// SetHeaders sets the "headers" field.
func (clc *CommunicationLogCreate) SetHeaders(s string) *CommunicationLogCreate {
	clc.mutation.SetHeaders(s)
	return clc
}

// SetBody sets the "body" field.
func (clc *CommunicationLogCreate) SetBody(b []byte) *CommunicationLogCreate {
	clc.mutation.SetBody(b)
	return clc
}

// Mutation returns the CommunicationLogMutation object of the builder.
func (clc *CommunicationLogCreate) Mutation() *CommunicationLogMutation {
	return clc.mutation
}

// Save creates the CommunicationLog in the database.
func (clc *CommunicationLogCreate) Save(ctx context.Context) (*CommunicationLog, error) {
	var (
		err  error
		node *CommunicationLog
	)
	if len(clc.hooks) == 0 {
		if err = clc.check(); err != nil {
			return nil, err
		}
		node, err = clc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunicationLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = clc.check(); err != nil {
				return nil, err
			}
			clc.mutation = mutation
			if node, err = clc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(clc.hooks) - 1; i >= 0; i-- {
			if clc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = clc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, clc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CommunicationLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CommunicationLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (clc *CommunicationLogCreate) SaveX(ctx context.Context) *CommunicationLog {
	v, err := clc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clc *CommunicationLogCreate) Exec(ctx context.Context) error {
	_, err := clc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clc *CommunicationLogCreate) ExecX(ctx context.Context) {
	if err := clc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (clc *CommunicationLogCreate) check() error {
	if _, ok := clc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "CommunicationLog.type"`)}
	}
	if _, ok := clc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "CommunicationLog.time"`)}
	}
	if _, ok := clc.mutation.Headers(); !ok {
		return &ValidationError{Name: "headers", err: errors.New(`ent: missing required field "CommunicationLog.headers"`)}
	}
	return nil
}

func (clc *CommunicationLogCreate) sqlSave(ctx context.Context) (*CommunicationLog, error) {
	_node, _spec := clc.createSpec()
	if err := sqlgraph.CreateNode(ctx, clc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (clc *CommunicationLogCreate) createSpec() (*CommunicationLog, *sqlgraph.CreateSpec) {
	var (
		_node = &CommunicationLog{config: clc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: communicationlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: communicationlog.FieldID,
			},
		}
	)
	if value, ok := clc.mutation.GetType(); ok {
		_spec.SetField(communicationlog.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := clc.mutation.Time(); ok {
		_spec.SetField(communicationlog.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	if value, ok := clc.mutation.Headers(); ok {
		_spec.SetField(communicationlog.FieldHeaders, field.TypeString, value)
		_node.Headers = value
	}
	if value, ok := clc.mutation.Body(); ok {
		_spec.SetField(communicationlog.FieldBody, field.TypeBytes, value)
		_node.Body = value
	}
	return _node, _spec
}

// CommunicationLogCreateBulk is the builder for creating many CommunicationLog entities in bulk.
type CommunicationLogCreateBulk struct {
	config
	builders []*CommunicationLogCreate
}

// Save creates the CommunicationLog entities in the database.
func (clcb *CommunicationLogCreateBulk) Save(ctx context.Context) ([]*CommunicationLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(clcb.builders))
	nodes := make([]*CommunicationLog, len(clcb.builders))
	mutators := make([]Mutator, len(clcb.builders))
	for i := range clcb.builders {
		func(i int, root context.Context) {
			builder := clcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommunicationLogMutation)
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
					_, err = mutators[i+1].Mutate(root, clcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, clcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, clcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (clcb *CommunicationLogCreateBulk) SaveX(ctx context.Context) []*CommunicationLog {
	v, err := clcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (clcb *CommunicationLogCreateBulk) Exec(ctx context.Context) error {
	_, err := clcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clcb *CommunicationLogCreateBulk) ExecX(ctx context.Context) {
	if err := clcb.Exec(ctx); err != nil {
		panic(err)
	}
}