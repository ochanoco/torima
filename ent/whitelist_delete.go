// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/proxy/ent/predicate"
	"github.com/ochanoco/proxy/ent/whitelist"
)

// WhiteListDelete is the builder for deleting a WhiteList entity.
type WhiteListDelete struct {
	config
	hooks    []Hook
	mutation *WhiteListMutation
}

// Where appends a list predicates to the WhiteListDelete builder.
func (wld *WhiteListDelete) Where(ps ...predicate.WhiteList) *WhiteListDelete {
	wld.mutation.Where(ps...)
	return wld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wld *WhiteListDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wld.hooks) == 0 {
		affected, err = wld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WhiteListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wld.mutation = mutation
			affected, err = wld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wld.hooks) - 1; i >= 0; i-- {
			if wld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wld *WhiteListDelete) ExecX(ctx context.Context) int {
	n, err := wld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wld *WhiteListDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: whitelist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: whitelist.FieldID,
			},
		},
	}
	if ps := wld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// WhiteListDeleteOne is the builder for deleting a single WhiteList entity.
type WhiteListDeleteOne struct {
	wld *WhiteListDelete
}

// Exec executes the deletion query.
func (wldo *WhiteListDeleteOne) Exec(ctx context.Context) error {
	n, err := wldo.wld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{whitelist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wldo *WhiteListDeleteOne) ExecX(ctx context.Context) {
	wldo.wld.ExecX(ctx)
}