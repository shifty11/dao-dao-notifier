// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/dao-dao-notifier/ent/predicate"
	"github.com/shifty11/dao-dao-notifier/ent/userwithzeroid"
)

// UserWithZeroIdUpdate is the builder for updating UserWithZeroId entities.
type UserWithZeroIdUpdate struct {
	config
	hooks    []Hook
	mutation *UserWithZeroIdMutation
}

// Where appends a list predicates to the UserWithZeroIdUpdate builder.
func (uwziu *UserWithZeroIdUpdate) Where(ps ...predicate.UserWithZeroId) *UserWithZeroIdUpdate {
	uwziu.mutation.Where(ps...)
	return uwziu
}

// SetUpdateTime sets the "update_time" field.
func (uwziu *UserWithZeroIdUpdate) SetUpdateTime(t time.Time) *UserWithZeroIdUpdate {
	uwziu.mutation.SetUpdateTime(t)
	return uwziu
}

// SetChatOrChannelID sets the "chat_or_channel_id" field.
func (uwziu *UserWithZeroIdUpdate) SetChatOrChannelID(i int64) *UserWithZeroIdUpdate {
	uwziu.mutation.ResetChatOrChannelID()
	uwziu.mutation.SetChatOrChannelID(i)
	return uwziu
}

// AddChatOrChannelID adds i to the "chat_or_channel_id" field.
func (uwziu *UserWithZeroIdUpdate) AddChatOrChannelID(i int64) *UserWithZeroIdUpdate {
	uwziu.mutation.AddChatOrChannelID(i)
	return uwziu
}

// SetChatOrChannelName sets the "chat_or_channel_name" field.
func (uwziu *UserWithZeroIdUpdate) SetChatOrChannelName(s string) *UserWithZeroIdUpdate {
	uwziu.mutation.SetChatOrChannelName(s)
	return uwziu
}

// SetIsGroup sets the "is_group" field.
func (uwziu *UserWithZeroIdUpdate) SetIsGroup(b bool) *UserWithZeroIdUpdate {
	uwziu.mutation.SetIsGroup(b)
	return uwziu
}

// SetChainID sets the "chain_id" field.
func (uwziu *UserWithZeroIdUpdate) SetChainID(s string) *UserWithZeroIdUpdate {
	uwziu.mutation.SetChainID(s)
	return uwziu
}

// Mutation returns the UserWithZeroIdMutation object of the builder.
func (uwziu *UserWithZeroIdUpdate) Mutation() *UserWithZeroIdMutation {
	return uwziu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uwziu *UserWithZeroIdUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uwziu.defaults()
	if len(uwziu.hooks) == 0 {
		affected, err = uwziu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithZeroIdMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwziu.mutation = mutation
			affected, err = uwziu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uwziu.hooks) - 1; i >= 0; i-- {
			if uwziu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwziu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uwziu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwziu *UserWithZeroIdUpdate) SaveX(ctx context.Context) int {
	affected, err := uwziu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uwziu *UserWithZeroIdUpdate) Exec(ctx context.Context) error {
	_, err := uwziu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwziu *UserWithZeroIdUpdate) ExecX(ctx context.Context) {
	if err := uwziu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwziu *UserWithZeroIdUpdate) defaults() {
	if _, ok := uwziu.mutation.UpdateTime(); !ok {
		v := userwithzeroid.UpdateDefaultUpdateTime()
		uwziu.mutation.SetUpdateTime(v)
	}
}

func (uwziu *UserWithZeroIdUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithzeroid.Table,
			Columns: userwithzeroid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userwithzeroid.FieldID,
			},
		},
	}
	if ps := uwziu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwziu.mutation.UpdateTime(); ok {
		_spec.SetField(userwithzeroid.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uwziu.mutation.ChatOrChannelID(); ok {
		_spec.SetField(userwithzeroid.FieldChatOrChannelID, field.TypeInt64, value)
	}
	if value, ok := uwziu.mutation.AddedChatOrChannelID(); ok {
		_spec.AddField(userwithzeroid.FieldChatOrChannelID, field.TypeInt64, value)
	}
	if value, ok := uwziu.mutation.ChatOrChannelName(); ok {
		_spec.SetField(userwithzeroid.FieldChatOrChannelName, field.TypeString, value)
	}
	if value, ok := uwziu.mutation.IsGroup(); ok {
		_spec.SetField(userwithzeroid.FieldIsGroup, field.TypeBool, value)
	}
	if value, ok := uwziu.mutation.ChainID(); ok {
		_spec.SetField(userwithzeroid.FieldChainID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uwziu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwithzeroid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserWithZeroIdUpdateOne is the builder for updating a single UserWithZeroId entity.
type UserWithZeroIdUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserWithZeroIdMutation
}

// SetUpdateTime sets the "update_time" field.
func (uwziuo *UserWithZeroIdUpdateOne) SetUpdateTime(t time.Time) *UserWithZeroIdUpdateOne {
	uwziuo.mutation.SetUpdateTime(t)
	return uwziuo
}

// SetChatOrChannelID sets the "chat_or_channel_id" field.
func (uwziuo *UserWithZeroIdUpdateOne) SetChatOrChannelID(i int64) *UserWithZeroIdUpdateOne {
	uwziuo.mutation.ResetChatOrChannelID()
	uwziuo.mutation.SetChatOrChannelID(i)
	return uwziuo
}

// AddChatOrChannelID adds i to the "chat_or_channel_id" field.
func (uwziuo *UserWithZeroIdUpdateOne) AddChatOrChannelID(i int64) *UserWithZeroIdUpdateOne {
	uwziuo.mutation.AddChatOrChannelID(i)
	return uwziuo
}

// SetChatOrChannelName sets the "chat_or_channel_name" field.
func (uwziuo *UserWithZeroIdUpdateOne) SetChatOrChannelName(s string) *UserWithZeroIdUpdateOne {
	uwziuo.mutation.SetChatOrChannelName(s)
	return uwziuo
}

// SetIsGroup sets the "is_group" field.
func (uwziuo *UserWithZeroIdUpdateOne) SetIsGroup(b bool) *UserWithZeroIdUpdateOne {
	uwziuo.mutation.SetIsGroup(b)
	return uwziuo
}

// SetChainID sets the "chain_id" field.
func (uwziuo *UserWithZeroIdUpdateOne) SetChainID(s string) *UserWithZeroIdUpdateOne {
	uwziuo.mutation.SetChainID(s)
	return uwziuo
}

// Mutation returns the UserWithZeroIdMutation object of the builder.
func (uwziuo *UserWithZeroIdUpdateOne) Mutation() *UserWithZeroIdMutation {
	return uwziuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uwziuo *UserWithZeroIdUpdateOne) Select(field string, fields ...string) *UserWithZeroIdUpdateOne {
	uwziuo.fields = append([]string{field}, fields...)
	return uwziuo
}

// Save executes the query and returns the updated UserWithZeroId entity.
func (uwziuo *UserWithZeroIdUpdateOne) Save(ctx context.Context) (*UserWithZeroId, error) {
	var (
		err  error
		node *UserWithZeroId
	)
	uwziuo.defaults()
	if len(uwziuo.hooks) == 0 {
		node, err = uwziuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserWithZeroIdMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uwziuo.mutation = mutation
			node, err = uwziuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uwziuo.hooks) - 1; i >= 0; i-- {
			if uwziuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uwziuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uwziuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserWithZeroId)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserWithZeroIdMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uwziuo *UserWithZeroIdUpdateOne) SaveX(ctx context.Context) *UserWithZeroId {
	node, err := uwziuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uwziuo *UserWithZeroIdUpdateOne) Exec(ctx context.Context) error {
	_, err := uwziuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uwziuo *UserWithZeroIdUpdateOne) ExecX(ctx context.Context) {
	if err := uwziuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uwziuo *UserWithZeroIdUpdateOne) defaults() {
	if _, ok := uwziuo.mutation.UpdateTime(); !ok {
		v := userwithzeroid.UpdateDefaultUpdateTime()
		uwziuo.mutation.SetUpdateTime(v)
	}
}

func (uwziuo *UserWithZeroIdUpdateOne) sqlSave(ctx context.Context) (_node *UserWithZeroId, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwithzeroid.Table,
			Columns: userwithzeroid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userwithzeroid.FieldID,
			},
		},
	}
	id, ok := uwziuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserWithZeroId.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uwziuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userwithzeroid.FieldID)
		for _, f := range fields {
			if !userwithzeroid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userwithzeroid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uwziuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uwziuo.mutation.UpdateTime(); ok {
		_spec.SetField(userwithzeroid.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uwziuo.mutation.ChatOrChannelID(); ok {
		_spec.SetField(userwithzeroid.FieldChatOrChannelID, field.TypeInt64, value)
	}
	if value, ok := uwziuo.mutation.AddedChatOrChannelID(); ok {
		_spec.AddField(userwithzeroid.FieldChatOrChannelID, field.TypeInt64, value)
	}
	if value, ok := uwziuo.mutation.ChatOrChannelName(); ok {
		_spec.SetField(userwithzeroid.FieldChatOrChannelName, field.TypeString, value)
	}
	if value, ok := uwziuo.mutation.IsGroup(); ok {
		_spec.SetField(userwithzeroid.FieldIsGroup, field.TypeBool, value)
	}
	if value, ok := uwziuo.mutation.ChainID(); ok {
		_spec.SetField(userwithzeroid.FieldChainID, field.TypeString, value)
	}
	_node = &UserWithZeroId{config: uwziuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uwziuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userwithzeroid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
