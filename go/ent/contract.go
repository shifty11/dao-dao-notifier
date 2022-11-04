// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/shifty11/dao-dao-notifier/ent/contract"
)

// Contract is the model entity for the Contract schema.
type Contract struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// ThumbnailURL holds the value of the "thumbnail_url" field.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// RPCEndpoint holds the value of the "rpc_endpoint" field.
	RPCEndpoint string `json:"rpc_endpoint,omitempty"`
	// ConfigVersion holds the value of the "config_version" field.
	ConfigVersion contract.ConfigVersion `json:"config_version,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContractQuery when eager-loading is set.
	Edges ContractEdges `json:"edges"`
}

// ContractEdges holds the relations/edges for other nodes in the graph.
type ContractEdges struct {
	// Proposals holds the value of the proposals edge.
	Proposals []*ContractProposal `json:"proposals,omitempty"`
	// TelegramChats holds the value of the telegram_chats edge.
	TelegramChats []*TelegramChat `json:"telegram_chats,omitempty"`
	// DiscordChannels holds the value of the discord_channels edge.
	DiscordChannels []*DiscordChannel `json:"discord_channels,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProposalsOrErr returns the Proposals value or an error if the edge
// was not loaded in eager-loading.
func (e ContractEdges) ProposalsOrErr() ([]*ContractProposal, error) {
	if e.loadedTypes[0] {
		return e.Proposals, nil
	}
	return nil, &NotLoadedError{edge: "proposals"}
}

// TelegramChatsOrErr returns the TelegramChats value or an error if the edge
// was not loaded in eager-loading.
func (e ContractEdges) TelegramChatsOrErr() ([]*TelegramChat, error) {
	if e.loadedTypes[1] {
		return e.TelegramChats, nil
	}
	return nil, &NotLoadedError{edge: "telegram_chats"}
}

// DiscordChannelsOrErr returns the DiscordChannels value or an error if the edge
// was not loaded in eager-loading.
func (e ContractEdges) DiscordChannelsOrErr() ([]*DiscordChannel, error) {
	if e.loadedTypes[2] {
		return e.DiscordChannels, nil
	}
	return nil, &NotLoadedError{edge: "discord_channels"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contract) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contract.FieldID:
			values[i] = new(sql.NullInt64)
		case contract.FieldAddress, contract.FieldName, contract.FieldDescription, contract.FieldImageURL, contract.FieldThumbnailURL, contract.FieldRPCEndpoint, contract.FieldConfigVersion:
			values[i] = new(sql.NullString)
		case contract.FieldCreateTime, contract.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Contract", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contract fields.
func (c *Contract) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case contract.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case contract.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case contract.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				c.Address = value.String
			}
		case contract.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case contract.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				c.Description = value.String
			}
		case contract.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				c.ImageURL = value.String
			}
		case contract.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				c.ThumbnailURL = value.String
			}
		case contract.FieldRPCEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rpc_endpoint", values[i])
			} else if value.Valid {
				c.RPCEndpoint = value.String
			}
		case contract.FieldConfigVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field config_version", values[i])
			} else if value.Valid {
				c.ConfigVersion = contract.ConfigVersion(value.String)
			}
		}
	}
	return nil
}

// QueryProposals queries the "proposals" edge of the Contract entity.
func (c *Contract) QueryProposals() *ContractProposalQuery {
	return (&ContractClient{config: c.config}).QueryProposals(c)
}

// QueryTelegramChats queries the "telegram_chats" edge of the Contract entity.
func (c *Contract) QueryTelegramChats() *TelegramChatQuery {
	return (&ContractClient{config: c.config}).QueryTelegramChats(c)
}

// QueryDiscordChannels queries the "discord_channels" edge of the Contract entity.
func (c *Contract) QueryDiscordChannels() *DiscordChannelQuery {
	return (&ContractClient{config: c.config}).QueryDiscordChannels(c)
}

// Update returns a builder for updating this Contract.
// Note that you need to call Contract.Unwrap() before calling this method if this Contract
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contract) Update() *ContractUpdateOne {
	return (&ContractClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Contract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contract) Unwrap() *Contract {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contract is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contract) String() string {
	var builder strings.Builder
	builder.WriteString("Contract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(c.Address)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(c.Description)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(c.ImageURL)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(c.ThumbnailURL)
	builder.WriteString(", ")
	builder.WriteString("rpc_endpoint=")
	builder.WriteString(c.RPCEndpoint)
	builder.WriteString(", ")
	builder.WriteString("config_version=")
	builder.WriteString(fmt.Sprintf("%v", c.ConfigVersion))
	builder.WriteByte(')')
	return builder.String()
}

// Contracts is a parsable slice of Contract.
type Contracts []*Contract

func (c Contracts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
