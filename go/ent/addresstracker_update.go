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
	"github.com/shifty11/cosmos-notifier/ent/addresstracker"
	"github.com/shifty11/cosmos-notifier/ent/chain"
	"github.com/shifty11/cosmos-notifier/ent/chainproposal"
	"github.com/shifty11/cosmos-notifier/ent/discordchannel"
	"github.com/shifty11/cosmos-notifier/ent/predicate"
	"github.com/shifty11/cosmos-notifier/ent/telegramchat"
)

// AddressTrackerUpdate is the builder for updating AddressTracker entities.
type AddressTrackerUpdate struct {
	config
	hooks    []Hook
	mutation *AddressTrackerMutation
}

// Where appends a list predicates to the AddressTrackerUpdate builder.
func (atu *AddressTrackerUpdate) Where(ps ...predicate.AddressTracker) *AddressTrackerUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUpdateTime sets the "update_time" field.
func (atu *AddressTrackerUpdate) SetUpdateTime(t time.Time) *AddressTrackerUpdate {
	atu.mutation.SetUpdateTime(t)
	return atu
}

// SetAddress sets the "address" field.
func (atu *AddressTrackerUpdate) SetAddress(s string) *AddressTrackerUpdate {
	atu.mutation.SetAddress(s)
	return atu
}

// SetNotificationInterval sets the "notification_interval" field.
func (atu *AddressTrackerUpdate) SetNotificationInterval(i int64) *AddressTrackerUpdate {
	atu.mutation.ResetNotificationInterval()
	atu.mutation.SetNotificationInterval(i)
	return atu
}

// AddNotificationInterval adds i to the "notification_interval" field.
func (atu *AddressTrackerUpdate) AddNotificationInterval(i int64) *AddressTrackerUpdate {
	atu.mutation.AddNotificationInterval(i)
	return atu
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (atu *AddressTrackerUpdate) SetChainID(id int) *AddressTrackerUpdate {
	atu.mutation.SetChainID(id)
	return atu
}

// SetChain sets the "chain" edge to the Chain entity.
func (atu *AddressTrackerUpdate) SetChain(c *Chain) *AddressTrackerUpdate {
	return atu.SetChainID(c.ID)
}

// SetDiscordChannelID sets the "discord_channel" edge to the DiscordChannel entity by ID.
func (atu *AddressTrackerUpdate) SetDiscordChannelID(id int) *AddressTrackerUpdate {
	atu.mutation.SetDiscordChannelID(id)
	return atu
}

// SetNillableDiscordChannelID sets the "discord_channel" edge to the DiscordChannel entity by ID if the given value is not nil.
func (atu *AddressTrackerUpdate) SetNillableDiscordChannelID(id *int) *AddressTrackerUpdate {
	if id != nil {
		atu = atu.SetDiscordChannelID(*id)
	}
	return atu
}

// SetDiscordChannel sets the "discord_channel" edge to the DiscordChannel entity.
func (atu *AddressTrackerUpdate) SetDiscordChannel(d *DiscordChannel) *AddressTrackerUpdate {
	return atu.SetDiscordChannelID(d.ID)
}

// SetTelegramChatID sets the "telegram_chat" edge to the TelegramChat entity by ID.
func (atu *AddressTrackerUpdate) SetTelegramChatID(id int) *AddressTrackerUpdate {
	atu.mutation.SetTelegramChatID(id)
	return atu
}

// SetNillableTelegramChatID sets the "telegram_chat" edge to the TelegramChat entity by ID if the given value is not nil.
func (atu *AddressTrackerUpdate) SetNillableTelegramChatID(id *int) *AddressTrackerUpdate {
	if id != nil {
		atu = atu.SetTelegramChatID(*id)
	}
	return atu
}

// SetTelegramChat sets the "telegram_chat" edge to the TelegramChat entity.
func (atu *AddressTrackerUpdate) SetTelegramChat(t *TelegramChat) *AddressTrackerUpdate {
	return atu.SetTelegramChatID(t.ID)
}

// AddChainProposalIDs adds the "chain_proposals" edge to the ChainProposal entity by IDs.
func (atu *AddressTrackerUpdate) AddChainProposalIDs(ids ...int) *AddressTrackerUpdate {
	atu.mutation.AddChainProposalIDs(ids...)
	return atu
}

// AddChainProposals adds the "chain_proposals" edges to the ChainProposal entity.
func (atu *AddressTrackerUpdate) AddChainProposals(c ...*ChainProposal) *AddressTrackerUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return atu.AddChainProposalIDs(ids...)
}

// Mutation returns the AddressTrackerMutation object of the builder.
func (atu *AddressTrackerUpdate) Mutation() *AddressTrackerMutation {
	return atu.mutation
}

// ClearChain clears the "chain" edge to the Chain entity.
func (atu *AddressTrackerUpdate) ClearChain() *AddressTrackerUpdate {
	atu.mutation.ClearChain()
	return atu
}

// ClearDiscordChannel clears the "discord_channel" edge to the DiscordChannel entity.
func (atu *AddressTrackerUpdate) ClearDiscordChannel() *AddressTrackerUpdate {
	atu.mutation.ClearDiscordChannel()
	return atu
}

// ClearTelegramChat clears the "telegram_chat" edge to the TelegramChat entity.
func (atu *AddressTrackerUpdate) ClearTelegramChat() *AddressTrackerUpdate {
	atu.mutation.ClearTelegramChat()
	return atu
}

// ClearChainProposals clears all "chain_proposals" edges to the ChainProposal entity.
func (atu *AddressTrackerUpdate) ClearChainProposals() *AddressTrackerUpdate {
	atu.mutation.ClearChainProposals()
	return atu
}

// RemoveChainProposalIDs removes the "chain_proposals" edge to ChainProposal entities by IDs.
func (atu *AddressTrackerUpdate) RemoveChainProposalIDs(ids ...int) *AddressTrackerUpdate {
	atu.mutation.RemoveChainProposalIDs(ids...)
	return atu
}

// RemoveChainProposals removes "chain_proposals" edges to ChainProposal entities.
func (atu *AddressTrackerUpdate) RemoveChainProposals(c ...*ChainProposal) *AddressTrackerUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return atu.RemoveChainProposalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AddressTrackerUpdate) Save(ctx context.Context) (int, error) {
	atu.defaults()
	return withHooks[int, AddressTrackerMutation](ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AddressTrackerUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AddressTrackerUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AddressTrackerUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *AddressTrackerUpdate) defaults() {
	if _, ok := atu.mutation.UpdateTime(); !ok {
		v := addresstracker.UpdateDefaultUpdateTime()
		atu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *AddressTrackerUpdate) check() error {
	if _, ok := atu.mutation.ChainID(); atu.mutation.ChainCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AddressTracker.chain"`)
	}
	return nil
}

func (atu *AddressTrackerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(addresstracker.Table, addresstracker.Columns, sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.UpdateTime(); ok {
		_spec.SetField(addresstracker.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := atu.mutation.Address(); ok {
		_spec.SetField(addresstracker.FieldAddress, field.TypeString, value)
	}
	if value, ok := atu.mutation.NotificationInterval(); ok {
		_spec.SetField(addresstracker.FieldNotificationInterval, field.TypeInt64, value)
	}
	if value, ok := atu.mutation.AddedNotificationInterval(); ok {
		_spec.AddField(addresstracker.FieldNotificationInterval, field.TypeInt64, value)
	}
	if atu.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.ChainTable,
			Columns: []string{addresstracker.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.ChainTable,
			Columns: []string{addresstracker.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.DiscordChannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.DiscordChannelTable,
			Columns: []string{addresstracker.DiscordChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.DiscordChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.DiscordChannelTable,
			Columns: []string{addresstracker.DiscordChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.TelegramChatCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.TelegramChatTable,
			Columns: []string{addresstracker.TelegramChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.TelegramChatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.TelegramChatTable,
			Columns: []string{addresstracker.TelegramChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RemovedChainProposalsIDs(); len(nodes) > 0 && !atu.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ChainProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{addresstracker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AddressTrackerUpdateOne is the builder for updating a single AddressTracker entity.
type AddressTrackerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AddressTrackerMutation
}

// SetUpdateTime sets the "update_time" field.
func (atuo *AddressTrackerUpdateOne) SetUpdateTime(t time.Time) *AddressTrackerUpdateOne {
	atuo.mutation.SetUpdateTime(t)
	return atuo
}

// SetAddress sets the "address" field.
func (atuo *AddressTrackerUpdateOne) SetAddress(s string) *AddressTrackerUpdateOne {
	atuo.mutation.SetAddress(s)
	return atuo
}

// SetNotificationInterval sets the "notification_interval" field.
func (atuo *AddressTrackerUpdateOne) SetNotificationInterval(i int64) *AddressTrackerUpdateOne {
	atuo.mutation.ResetNotificationInterval()
	atuo.mutation.SetNotificationInterval(i)
	return atuo
}

// AddNotificationInterval adds i to the "notification_interval" field.
func (atuo *AddressTrackerUpdateOne) AddNotificationInterval(i int64) *AddressTrackerUpdateOne {
	atuo.mutation.AddNotificationInterval(i)
	return atuo
}

// SetChainID sets the "chain" edge to the Chain entity by ID.
func (atuo *AddressTrackerUpdateOne) SetChainID(id int) *AddressTrackerUpdateOne {
	atuo.mutation.SetChainID(id)
	return atuo
}

// SetChain sets the "chain" edge to the Chain entity.
func (atuo *AddressTrackerUpdateOne) SetChain(c *Chain) *AddressTrackerUpdateOne {
	return atuo.SetChainID(c.ID)
}

// SetDiscordChannelID sets the "discord_channel" edge to the DiscordChannel entity by ID.
func (atuo *AddressTrackerUpdateOne) SetDiscordChannelID(id int) *AddressTrackerUpdateOne {
	atuo.mutation.SetDiscordChannelID(id)
	return atuo
}

// SetNillableDiscordChannelID sets the "discord_channel" edge to the DiscordChannel entity by ID if the given value is not nil.
func (atuo *AddressTrackerUpdateOne) SetNillableDiscordChannelID(id *int) *AddressTrackerUpdateOne {
	if id != nil {
		atuo = atuo.SetDiscordChannelID(*id)
	}
	return atuo
}

// SetDiscordChannel sets the "discord_channel" edge to the DiscordChannel entity.
func (atuo *AddressTrackerUpdateOne) SetDiscordChannel(d *DiscordChannel) *AddressTrackerUpdateOne {
	return atuo.SetDiscordChannelID(d.ID)
}

// SetTelegramChatID sets the "telegram_chat" edge to the TelegramChat entity by ID.
func (atuo *AddressTrackerUpdateOne) SetTelegramChatID(id int) *AddressTrackerUpdateOne {
	atuo.mutation.SetTelegramChatID(id)
	return atuo
}

// SetNillableTelegramChatID sets the "telegram_chat" edge to the TelegramChat entity by ID if the given value is not nil.
func (atuo *AddressTrackerUpdateOne) SetNillableTelegramChatID(id *int) *AddressTrackerUpdateOne {
	if id != nil {
		atuo = atuo.SetTelegramChatID(*id)
	}
	return atuo
}

// SetTelegramChat sets the "telegram_chat" edge to the TelegramChat entity.
func (atuo *AddressTrackerUpdateOne) SetTelegramChat(t *TelegramChat) *AddressTrackerUpdateOne {
	return atuo.SetTelegramChatID(t.ID)
}

// AddChainProposalIDs adds the "chain_proposals" edge to the ChainProposal entity by IDs.
func (atuo *AddressTrackerUpdateOne) AddChainProposalIDs(ids ...int) *AddressTrackerUpdateOne {
	atuo.mutation.AddChainProposalIDs(ids...)
	return atuo
}

// AddChainProposals adds the "chain_proposals" edges to the ChainProposal entity.
func (atuo *AddressTrackerUpdateOne) AddChainProposals(c ...*ChainProposal) *AddressTrackerUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return atuo.AddChainProposalIDs(ids...)
}

// Mutation returns the AddressTrackerMutation object of the builder.
func (atuo *AddressTrackerUpdateOne) Mutation() *AddressTrackerMutation {
	return atuo.mutation
}

// ClearChain clears the "chain" edge to the Chain entity.
func (atuo *AddressTrackerUpdateOne) ClearChain() *AddressTrackerUpdateOne {
	atuo.mutation.ClearChain()
	return atuo
}

// ClearDiscordChannel clears the "discord_channel" edge to the DiscordChannel entity.
func (atuo *AddressTrackerUpdateOne) ClearDiscordChannel() *AddressTrackerUpdateOne {
	atuo.mutation.ClearDiscordChannel()
	return atuo
}

// ClearTelegramChat clears the "telegram_chat" edge to the TelegramChat entity.
func (atuo *AddressTrackerUpdateOne) ClearTelegramChat() *AddressTrackerUpdateOne {
	atuo.mutation.ClearTelegramChat()
	return atuo
}

// ClearChainProposals clears all "chain_proposals" edges to the ChainProposal entity.
func (atuo *AddressTrackerUpdateOne) ClearChainProposals() *AddressTrackerUpdateOne {
	atuo.mutation.ClearChainProposals()
	return atuo
}

// RemoveChainProposalIDs removes the "chain_proposals" edge to ChainProposal entities by IDs.
func (atuo *AddressTrackerUpdateOne) RemoveChainProposalIDs(ids ...int) *AddressTrackerUpdateOne {
	atuo.mutation.RemoveChainProposalIDs(ids...)
	return atuo
}

// RemoveChainProposals removes "chain_proposals" edges to ChainProposal entities.
func (atuo *AddressTrackerUpdateOne) RemoveChainProposals(c ...*ChainProposal) *AddressTrackerUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return atuo.RemoveChainProposalIDs(ids...)
}

// Where appends a list predicates to the AddressTrackerUpdate builder.
func (atuo *AddressTrackerUpdateOne) Where(ps ...predicate.AddressTracker) *AddressTrackerUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AddressTrackerUpdateOne) Select(field string, fields ...string) *AddressTrackerUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AddressTracker entity.
func (atuo *AddressTrackerUpdateOne) Save(ctx context.Context) (*AddressTracker, error) {
	atuo.defaults()
	return withHooks[*AddressTracker, AddressTrackerMutation](ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AddressTrackerUpdateOne) SaveX(ctx context.Context) *AddressTracker {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AddressTrackerUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AddressTrackerUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *AddressTrackerUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdateTime(); !ok {
		v := addresstracker.UpdateDefaultUpdateTime()
		atuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *AddressTrackerUpdateOne) check() error {
	if _, ok := atuo.mutation.ChainID(); atuo.mutation.ChainCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AddressTracker.chain"`)
	}
	return nil
}

func (atuo *AddressTrackerUpdateOne) sqlSave(ctx context.Context) (_node *AddressTracker, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(addresstracker.Table, addresstracker.Columns, sqlgraph.NewFieldSpec(addresstracker.FieldID, field.TypeInt))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AddressTracker.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, addresstracker.FieldID)
		for _, f := range fields {
			if !addresstracker.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != addresstracker.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.UpdateTime(); ok {
		_spec.SetField(addresstracker.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := atuo.mutation.Address(); ok {
		_spec.SetField(addresstracker.FieldAddress, field.TypeString, value)
	}
	if value, ok := atuo.mutation.NotificationInterval(); ok {
		_spec.SetField(addresstracker.FieldNotificationInterval, field.TypeInt64, value)
	}
	if value, ok := atuo.mutation.AddedNotificationInterval(); ok {
		_spec.AddField(addresstracker.FieldNotificationInterval, field.TypeInt64, value)
	}
	if atuo.mutation.ChainCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.ChainTable,
			Columns: []string{addresstracker.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ChainIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.ChainTable,
			Columns: []string{addresstracker.ChainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chain.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.DiscordChannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.DiscordChannelTable,
			Columns: []string{addresstracker.DiscordChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.DiscordChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.DiscordChannelTable,
			Columns: []string{addresstracker.DiscordChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.TelegramChatCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.TelegramChatTable,
			Columns: []string{addresstracker.TelegramChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.TelegramChatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   addresstracker.TelegramChatTable,
			Columns: []string{addresstracker.TelegramChatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(telegramchat.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RemovedChainProposalsIDs(); len(nodes) > 0 && !atuo.mutation.ChainProposalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ChainProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   addresstracker.ChainProposalsTable,
			Columns: addresstracker.ChainProposalsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chainproposal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AddressTracker{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{addresstracker.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
