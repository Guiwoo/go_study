// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"semina_entgo/custom"
	"semina_entgo/ent/predicate"
	"semina_entgo/ent/tester"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TesterUpdate is the builder for updating Tester entities.
type TesterUpdate struct {
	config
	hooks    []Hook
	mutation *TesterMutation
}

// Where appends a list predicates to the TesterUpdate builder.
func (tu *TesterUpdate) Where(ps ...predicate.Tester) *TesterUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetPascalCase sets the "PascalCase" field.
func (tu *TesterUpdate) SetPascalCase(s string) *TesterUpdate {
	tu.mutation.SetPascalCase(s)
	return tu
}

// SetNillablePascalCase sets the "PascalCase" field if the given value is not nil.
func (tu *TesterUpdate) SetNillablePascalCase(s *string) *TesterUpdate {
	if s != nil {
		tu.SetPascalCase(*s)
	}
	return tu
}

// SetLetMeCheck sets the "let_me_check" field.
func (tu *TesterUpdate) SetLetMeCheck(s string) *TesterUpdate {
	tu.mutation.SetLetMeCheck(s)
	return tu
}

// SetNillableLetMeCheck sets the "let_me_check" field if the given value is not nil.
func (tu *TesterUpdate) SetNillableLetMeCheck(s *string) *TesterUpdate {
	if s != nil {
		tu.SetLetMeCheck(*s)
	}
	return tu
}

// SetSize sets the "size" field.
func (tu *TesterUpdate) SetSize(t tester.Size) *TesterUpdate {
	tu.mutation.SetSize(t)
	return tu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (tu *TesterUpdate) SetNillableSize(t *tester.Size) *TesterUpdate {
	if t != nil {
		tu.SetSize(*t)
	}
	return tu
}

// SetShape sets the "shape" field.
func (tu *TesterUpdate) SetShape(c custom.Shape) *TesterUpdate {
	tu.mutation.SetShape(c)
	return tu
}

// SetNillableShape sets the "shape" field if the given value is not nil.
func (tu *TesterUpdate) SetNillableShape(c *custom.Shape) *TesterUpdate {
	if c != nil {
		tu.SetShape(*c)
	}
	return tu
}

// SetLevel sets the "level" field.
func (tu *TesterUpdate) SetLevel(c custom.Level) *TesterUpdate {
	tu.mutation.SetLevel(c)
	return tu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (tu *TesterUpdate) SetNillableLevel(c *custom.Level) *TesterUpdate {
	if c != nil {
		tu.SetLevel(*c)
	}
	return tu
}

// Mutation returns the TesterMutation object of the builder.
func (tu *TesterUpdate) Mutation() *TesterMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TesterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TesterUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TesterUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TesterUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TesterUpdate) check() error {
	if v, ok := tu.mutation.Size(); ok {
		if err := tester.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "Tester.size": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Shape(); ok {
		if err := tester.ShapeValidator(v); err != nil {
			return &ValidationError{Name: "shape", err: fmt.Errorf(`ent: validator failed for field "Tester.shape": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Level(); ok {
		if err := tester.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Tester.level": %w`, err)}
		}
	}
	return nil
}

func (tu *TesterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tester.Table, tester.Columns, sqlgraph.NewFieldSpec(tester.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.PascalCase(); ok {
		_spec.SetField(tester.FieldPascalCase, field.TypeString, value)
	}
	if value, ok := tu.mutation.LetMeCheck(); ok {
		_spec.SetField(tester.FieldLetMeCheck, field.TypeString, value)
	}
	if value, ok := tu.mutation.Size(); ok {
		_spec.SetField(tester.FieldSize, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Shape(); ok {
		_spec.SetField(tester.FieldShape, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Level(); ok {
		_spec.SetField(tester.FieldLevel, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tester.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TesterUpdateOne is the builder for updating a single Tester dto.
type TesterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TesterMutation
}

// SetPascalCase sets the "PascalCase" field.
func (tuo *TesterUpdateOne) SetPascalCase(s string) *TesterUpdateOne {
	tuo.mutation.SetPascalCase(s)
	return tuo
}

// SetNillablePascalCase sets the "PascalCase" field if the given value is not nil.
func (tuo *TesterUpdateOne) SetNillablePascalCase(s *string) *TesterUpdateOne {
	if s != nil {
		tuo.SetPascalCase(*s)
	}
	return tuo
}

// SetLetMeCheck sets the "let_me_check" field.
func (tuo *TesterUpdateOne) SetLetMeCheck(s string) *TesterUpdateOne {
	tuo.mutation.SetLetMeCheck(s)
	return tuo
}

// SetNillableLetMeCheck sets the "let_me_check" field if the given value is not nil.
func (tuo *TesterUpdateOne) SetNillableLetMeCheck(s *string) *TesterUpdateOne {
	if s != nil {
		tuo.SetLetMeCheck(*s)
	}
	return tuo
}

// SetSize sets the "size" field.
func (tuo *TesterUpdateOne) SetSize(t tester.Size) *TesterUpdateOne {
	tuo.mutation.SetSize(t)
	return tuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (tuo *TesterUpdateOne) SetNillableSize(t *tester.Size) *TesterUpdateOne {
	if t != nil {
		tuo.SetSize(*t)
	}
	return tuo
}

// SetShape sets the "shape" field.
func (tuo *TesterUpdateOne) SetShape(c custom.Shape) *TesterUpdateOne {
	tuo.mutation.SetShape(c)
	return tuo
}

// SetNillableShape sets the "shape" field if the given value is not nil.
func (tuo *TesterUpdateOne) SetNillableShape(c *custom.Shape) *TesterUpdateOne {
	if c != nil {
		tuo.SetShape(*c)
	}
	return tuo
}

// SetLevel sets the "level" field.
func (tuo *TesterUpdateOne) SetLevel(c custom.Level) *TesterUpdateOne {
	tuo.mutation.SetLevel(c)
	return tuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (tuo *TesterUpdateOne) SetNillableLevel(c *custom.Level) *TesterUpdateOne {
	if c != nil {
		tuo.SetLevel(*c)
	}
	return tuo
}

// Mutation returns the TesterMutation object of the builder.
func (tuo *TesterUpdateOne) Mutation() *TesterMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TesterUpdate builder.
func (tuo *TesterUpdateOne) Where(ps ...predicate.Tester) *TesterUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned dto.
// The default is selecting all fields defined in the dto schema.
func (tuo *TesterUpdateOne) Select(field string, fields ...string) *TesterUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tester dto.
func (tuo *TesterUpdateOne) Save(ctx context.Context) (*Tester, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TesterUpdateOne) SaveX(ctx context.Context) *Tester {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the dto.
func (tuo *TesterUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TesterUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TesterUpdateOne) check() error {
	if v, ok := tuo.mutation.Size(); ok {
		if err := tester.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "Tester.size": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Shape(); ok {
		if err := tester.ShapeValidator(v); err != nil {
			return &ValidationError{Name: "shape", err: fmt.Errorf(`ent: validator failed for field "Tester.shape": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Level(); ok {
		if err := tester.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Tester.level": %w`, err)}
		}
	}
	return nil
}

func (tuo *TesterUpdateOne) sqlSave(ctx context.Context) (_node *Tester, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tester.Table, tester.Columns, sqlgraph.NewFieldSpec(tester.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tester.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tester.FieldID)
		for _, f := range fields {
			if !tester.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tester.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.PascalCase(); ok {
		_spec.SetField(tester.FieldPascalCase, field.TypeString, value)
	}
	if value, ok := tuo.mutation.LetMeCheck(); ok {
		_spec.SetField(tester.FieldLetMeCheck, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Size(); ok {
		_spec.SetField(tester.FieldSize, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Shape(); ok {
		_spec.SetField(tester.FieldShape, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Level(); ok {
		_spec.SetField(tester.FieldLevel, field.TypeEnum, value)
	}
	_node = &Tester{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tester.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
