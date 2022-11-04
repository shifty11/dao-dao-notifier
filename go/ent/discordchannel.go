// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/dao-dao-notifier/ent/discordchannel"
)

// DiscordChannel is the model entity for the DiscordChannel schema.
type DiscordChannel struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// ChannelID holds the value of the "channel_id" field.
	ChannelID int64 `json:"channel_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsGroup holds the value of the "is_group" field.
	IsGroup bool `json:"is_group,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiscordChannelQuery when eager-loading is set.
	Edges DiscordChannelEdges `json:"edges"`
}

// DiscordChannelEdges holds the relations/edges for other nodes in the graph.
type DiscordChannelEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Contracts holds the value of the contracts edge.
	Contracts []*Contract `json:"contracts,omitempty"`
	// Chains holds the value of the chains edge.
	Chains []*Chain `json:"chains,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordChannelEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ContractsOrErr returns the Contracts value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordChannelEdges) ContractsOrErr() ([]*Contract, error) {
	if e.loadedTypes[1] {
		return e.Contracts, nil
	}
	return nil, &NotLoadedError{edge: "contracts"}
}

// ChainsOrErr returns the Chains value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordChannelEdges) ChainsOrErr() ([]*Chain, error) {
	if e.loadedTypes[2] {
		return e.Chains, nil
	}
	return nil, &NotLoadedError{edge: "chains"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DiscordChannel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case discordchannel.FieldIsGroup:
			values[i] = new(sql.NullBool)
		case discordchannel.FieldID, discordchannel.FieldChannelID:
			values[i] = new(sql.NullInt64)
		case discordchannel.FieldName:
			values[i] = new(sql.NullString)
		case discordchannel.FieldCreateTime, discordchannel.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DiscordChannel", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DiscordChannel fields.
func (dc *DiscordChannel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case discordchannel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dc.ID = int(value.Int64)
		case discordchannel.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				dc.CreateTime = value.Time
			}
		case discordchannel.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				dc.UpdateTime = value.Time
			}
		case discordchannel.FieldChannelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				dc.ChannelID = value.Int64
			}
		case discordchannel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dc.Name = value.String
			}
		case discordchannel.FieldIsGroup:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_group", values[i])
			} else if value.Valid {
				dc.IsGroup = value.Bool
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the DiscordChannel entity.
func (dc *DiscordChannel) QueryUsers() *UserQuery {
	return (&DiscordChannelClient{config: dc.config}).QueryUsers(dc)
}

// QueryContracts queries the "contracts" edge of the DiscordChannel entity.
func (dc *DiscordChannel) QueryContracts() *ContractQuery {
	return (&DiscordChannelClient{config: dc.config}).QueryContracts(dc)
}

// QueryChains queries the "chains" edge of the DiscordChannel entity.
func (dc *DiscordChannel) QueryChains() *ChainQuery {
	return (&DiscordChannelClient{config: dc.config}).QueryChains(dc)
}

// Update returns a builder for updating this DiscordChannel.
// Note that you need to call DiscordChannel.Unwrap() before calling this method if this DiscordChannel
// was returned from a transaction, and the transaction was committed or rolled back.
func (dc *DiscordChannel) Update() *DiscordChannelUpdateOne {
	return (&DiscordChannelClient{config: dc.config}).UpdateOne(dc)
}

// Unwrap unwraps the DiscordChannel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dc *DiscordChannel) Unwrap() *DiscordChannel {
	_tx, ok := dc.config.driver.(*txDriver)
	if !ok {
		panic("ent: DiscordChannel is not a transactional entity")
	}
	dc.config.driver = _tx.drv
	return dc
}

// String implements the fmt.Stringer.
func (dc *DiscordChannel) String() string {
	var builder strings.Builder
	builder.WriteString("DiscordChannel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(dc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(dc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("channel_id=")
	builder.WriteString(fmt.Sprintf("%v", dc.ChannelID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(dc.Name)
	builder.WriteString(", ")
	builder.WriteString("is_group=")
	builder.WriteString(fmt.Sprintf("%v", dc.IsGroup))
	builder.WriteByte(')')
	return builder.String()
}

// DiscordChannels is a parsable slice of DiscordChannel.
type DiscordChannels []*DiscordChannel

func (dc DiscordChannels) config(cfg config) {
	for _i := range dc {
		dc[_i].config = cfg
	}
}
