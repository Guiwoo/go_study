// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"semina_entgo/custom"
	"semina_entgo/ent/tester"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TesterCreate is the builder for creating a Tester dto.
type TesterCreate struct {
	config
	mutation *TesterMutation
	hooks    []Hook
}

// SetPascalCase sets the "PascalCase" field.
func (tc *TesterCreate) SetPascalCase(s string) *TesterCreate {
	tc.mutation.SetPascalCase(s)
	return tc
}

// SetLetMeCheck sets the "let_me_check" field.
func (tc *TesterCreate) SetLetMeCheck(s string) *TesterCreate {
	tc.mutation.SetLetMeCheck(s)
	return tc
}

// SetSize sets the "size" field.
func (tc *TesterCreate) SetSize(t tester.Size) *TesterCreate {
	tc.mutation.SetSize(t)
	return tc
}

// SetShape sets the "shape" field.
func (tc *TesterCreate) SetShape(c custom.Shape) *TesterCreate {
	tc.mutation.SetShape(c)
	return tc
}

// SetLevel sets the "level" field.
func (tc *TesterCreate) SetLevel(c custom.Level) *TesterCreate {
	tc.mutation.SetLevel(c)
	return tc
}

// Mutation returns the TesterMutation object of the builder.
func (tc *TesterCreate) Mutation() *TesterMutation {
	return tc.mutation
}

// Save creates the Tester in the database.
func (tc *TesterCreate) Save(ctx context.Context) (*Tester, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TesterCreate) SaveX(ctx context.Context) *Tester {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TesterCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TesterCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TesterCreate) check() error {
	if _, ok := tc.mutation.PascalCase(); !ok {
		return &ValidationError{Name: "PascalCase", err: errors.New(`ent: missing required field "Tester.PascalCase"`)}
	}
	if _, ok := tc.mutation.LetMeCheck(); !ok {
		return &ValidationError{Name: "let_me_check", err: errors.New(`ent: missing required field "Tester.let_me_check"`)}
	}
	if _, ok := tc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Tester.size"`)}
	}
	if v, ok := tc.mutation.Size(); ok {
		if err := tester.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "Tester.size": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Shape(); !ok {
		return &ValidationError{Name: "shape", err: errors.New(`ent: missing required field "Tester.shape"`)}
	}
	if v, ok := tc.mutation.Shape(); ok {
		if err := tester.ShapeValidator(v); err != nil {
			return &ValidationError{Name: "shape", err: fmt.Errorf(`ent: validator failed for field "Tester.shape": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Tester.level"`)}
	}
	if v, ok := tc.mutation.Level(); ok {
		if err := tester.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Tester.level": %w`, err)}
		}
	}
	return nil
}

func (tc *TesterCreate) sqlSave(ctx context.Context) (*Tester, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TesterCreate) createSpec() (*Tester, *sqlgraph.CreateSpec) {
	var (
		_node = &Tester{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tester.Table, sqlgraph.NewFieldSpec(tester.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.PascalCase(); ok {
		_spec.SetField(tester.FieldPascalCase, field.TypeString, value)
		_node.PascalCase = value
	}
	if value, ok := tc.mutation.LetMeCheck(); ok {
		_spec.SetField(tester.FieldLetMeCheck, field.TypeString, value)
		_node.LetMeCheck = value
	}
	if value, ok := tc.mutation.Size(); ok {
		_spec.SetField(tester.FieldSize, field.TypeEnum, value)
		_node.Size = value
	}
	if value, ok := tc.mutation.Shape(); ok {
		_spec.SetField(tester.FieldShape, field.TypeEnum, value)
		_node.Shape = value
	}
	if value, ok := tc.mutation.Level(); ok {
		_spec.SetField(tester.FieldLevel, field.TypeEnum, value)
		_node.Level = value
	}
	return _node, _spec
}

// TesterCreateBulk is the builder for creating many Tester entities in bulk.
type TesterCreateBulk struct {
	config
	err      error
	builders []*TesterCreate
}

// Save creates the Tester entities in the database.
func (tcb *TesterCreateBulk) Save(ctx context.Context) ([]*Tester, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tester, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TesterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TesterCreateBulk) SaveX(ctx context.Context) []*Tester {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TesterCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TesterCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
