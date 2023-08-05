// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/proxy/ent/hashchain"
	"github.com/ochanoco/proxy/ent/predicate"
	"github.com/ochanoco/proxy/ent/servicelog"
)

// HashChainUpdate is the builder for updating HashChain entities.
type HashChainUpdate struct {
	config
	hooks    []Hook
	mutation *HashChainMutation
}

// Where appends a list predicates to the HashChainUpdate builder.
func (hcu *HashChainUpdate) Where(ps ...predicate.HashChain) *HashChainUpdate {
	hcu.mutation.Where(ps...)
	return hcu
}

// SetHash sets the "hash" field.
func (hcu *HashChainUpdate) SetHash(b []byte) *HashChainUpdate {
	hcu.mutation.SetHash(b)
	return hcu
}

// SetSignature sets the "signature" field.
func (hcu *HashChainUpdate) SetSignature(b []byte) *HashChainUpdate {
	hcu.mutation.SetSignature(b)
	return hcu
}

// SetLogID sets the "log" edge to the ServiceLog entity by ID.
func (hcu *HashChainUpdate) SetLogID(id int) *HashChainUpdate {
	hcu.mutation.SetLogID(id)
	return hcu
}

// SetNillableLogID sets the "log" edge to the ServiceLog entity by ID if the given value is not nil.
func (hcu *HashChainUpdate) SetNillableLogID(id *int) *HashChainUpdate {
	if id != nil {
		hcu = hcu.SetLogID(*id)
	}
	return hcu
}

// SetLog sets the "log" edge to the ServiceLog entity.
func (hcu *HashChainUpdate) SetLog(s *ServiceLog) *HashChainUpdate {
	return hcu.SetLogID(s.ID)
}

// Mutation returns the HashChainMutation object of the builder.
func (hcu *HashChainUpdate) Mutation() *HashChainMutation {
	return hcu.mutation
}

// ClearLog clears the "log" edge to the ServiceLog entity.
func (hcu *HashChainUpdate) ClearLog() *HashChainUpdate {
	hcu.mutation.ClearLog()
	return hcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hcu *HashChainUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hcu.hooks) == 0 {
		affected, err = hcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HashChainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hcu.mutation = mutation
			affected, err = hcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hcu.hooks) - 1; i >= 0; i-- {
			if hcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hcu *HashChainUpdate) SaveX(ctx context.Context) int {
	affected, err := hcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hcu *HashChainUpdate) Exec(ctx context.Context) error {
	_, err := hcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcu *HashChainUpdate) ExecX(ctx context.Context) {
	if err := hcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hcu *HashChainUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hashchain.Table,
			Columns: hashchain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hashchain.FieldID,
			},
		},
	}
	if ps := hcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hcu.mutation.Hash(); ok {
		_spec.SetField(hashchain.FieldHash, field.TypeBytes, value)
	}
	if value, ok := hcu.mutation.Signature(); ok {
		_spec.SetField(hashchain.FieldSignature, field.TypeBytes, value)
	}
	if hcu.mutation.LogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashchain.LogTable,
			Columns: []string{hashchain.LogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicelog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hcu.mutation.LogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashchain.LogTable,
			Columns: []string{hashchain.LogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicelog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hashchain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// HashChainUpdateOne is the builder for updating a single HashChain entity.
type HashChainUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HashChainMutation
}

// SetHash sets the "hash" field.
func (hcuo *HashChainUpdateOne) SetHash(b []byte) *HashChainUpdateOne {
	hcuo.mutation.SetHash(b)
	return hcuo
}

// SetSignature sets the "signature" field.
func (hcuo *HashChainUpdateOne) SetSignature(b []byte) *HashChainUpdateOne {
	hcuo.mutation.SetSignature(b)
	return hcuo
}

// SetLogID sets the "log" edge to the ServiceLog entity by ID.
func (hcuo *HashChainUpdateOne) SetLogID(id int) *HashChainUpdateOne {
	hcuo.mutation.SetLogID(id)
	return hcuo
}

// SetNillableLogID sets the "log" edge to the ServiceLog entity by ID if the given value is not nil.
func (hcuo *HashChainUpdateOne) SetNillableLogID(id *int) *HashChainUpdateOne {
	if id != nil {
		hcuo = hcuo.SetLogID(*id)
	}
	return hcuo
}

// SetLog sets the "log" edge to the ServiceLog entity.
func (hcuo *HashChainUpdateOne) SetLog(s *ServiceLog) *HashChainUpdateOne {
	return hcuo.SetLogID(s.ID)
}

// Mutation returns the HashChainMutation object of the builder.
func (hcuo *HashChainUpdateOne) Mutation() *HashChainMutation {
	return hcuo.mutation
}

// ClearLog clears the "log" edge to the ServiceLog entity.
func (hcuo *HashChainUpdateOne) ClearLog() *HashChainUpdateOne {
	hcuo.mutation.ClearLog()
	return hcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hcuo *HashChainUpdateOne) Select(field string, fields ...string) *HashChainUpdateOne {
	hcuo.fields = append([]string{field}, fields...)
	return hcuo
}

// Save executes the query and returns the updated HashChain entity.
func (hcuo *HashChainUpdateOne) Save(ctx context.Context) (*HashChain, error) {
	var (
		err  error
		node *HashChain
	)
	if len(hcuo.hooks) == 0 {
		node, err = hcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HashChainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hcuo.mutation = mutation
			node, err = hcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hcuo.hooks) - 1; i >= 0; i-- {
			if hcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HashChain)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HashChainMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (hcuo *HashChainUpdateOne) SaveX(ctx context.Context) *HashChain {
	node, err := hcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hcuo *HashChainUpdateOne) Exec(ctx context.Context) error {
	_, err := hcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcuo *HashChainUpdateOne) ExecX(ctx context.Context) {
	if err := hcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hcuo *HashChainUpdateOne) sqlSave(ctx context.Context) (_node *HashChain, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   hashchain.Table,
			Columns: hashchain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hashchain.FieldID,
			},
		},
	}
	id, ok := hcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HashChain.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hashchain.FieldID)
		for _, f := range fields {
			if !hashchain.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hashchain.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hcuo.mutation.Hash(); ok {
		_spec.SetField(hashchain.FieldHash, field.TypeBytes, value)
	}
	if value, ok := hcuo.mutation.Signature(); ok {
		_spec.SetField(hashchain.FieldSignature, field.TypeBytes, value)
	}
	if hcuo.mutation.LogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashchain.LogTable,
			Columns: []string{hashchain.LogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicelog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hcuo.mutation.LogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hashchain.LogTable,
			Columns: []string{hashchain.LogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: servicelog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &HashChain{config: hcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hashchain.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}