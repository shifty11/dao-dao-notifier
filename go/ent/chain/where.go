// Code generated by ent, DO NOT EDIT.

package chain

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shifty11/dao-dao-notifier/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ChainID applies equality check predicate on the "chain_id" field. It's identical to ChainIDEQ.
func ChainID(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// PrettyName applies equality check predicate on the "pretty_name" field. It's identical to PrettyNameEQ.
func PrettyName(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrettyName), v))
	})
}

// IsEnabled applies equality check predicate on the "is_enabled" field. It's identical to IsEnabledEQ.
func IsEnabled(v bool) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsEnabled), v))
	})
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageURL), v))
	})
}

// ThumbnailURL applies equality check predicate on the "thumbnail_url" field. It's identical to ThumbnailURLEQ.
func ThumbnailURL(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnailURL), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ChainIDEQ applies the EQ predicate on the "chain_id" field.
func ChainIDEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainID), v))
	})
}

// ChainIDNEQ applies the NEQ predicate on the "chain_id" field.
func ChainIDNEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainID), v))
	})
}

// ChainIDIn applies the In predicate on the "chain_id" field.
func ChainIDIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChainID), v...))
	})
}

// ChainIDNotIn applies the NotIn predicate on the "chain_id" field.
func ChainIDNotIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChainID), v...))
	})
}

// ChainIDGT applies the GT predicate on the "chain_id" field.
func ChainIDGT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainID), v))
	})
}

// ChainIDGTE applies the GTE predicate on the "chain_id" field.
func ChainIDGTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainID), v))
	})
}

// ChainIDLT applies the LT predicate on the "chain_id" field.
func ChainIDLT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainID), v))
	})
}

// ChainIDLTE applies the LTE predicate on the "chain_id" field.
func ChainIDLTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainID), v))
	})
}

// ChainIDContains applies the Contains predicate on the "chain_id" field.
func ChainIDContains(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChainID), v))
	})
}

// ChainIDHasPrefix applies the HasPrefix predicate on the "chain_id" field.
func ChainIDHasPrefix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChainID), v))
	})
}

// ChainIDHasSuffix applies the HasSuffix predicate on the "chain_id" field.
func ChainIDHasSuffix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChainID), v))
	})
}

// ChainIDEqualFold applies the EqualFold predicate on the "chain_id" field.
func ChainIDEqualFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChainID), v))
	})
}

// ChainIDContainsFold applies the ContainsFold predicate on the "chain_id" field.
func ChainIDContainsFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChainID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PrettyNameEQ applies the EQ predicate on the "pretty_name" field.
func PrettyNameEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrettyName), v))
	})
}

// PrettyNameNEQ applies the NEQ predicate on the "pretty_name" field.
func PrettyNameNEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrettyName), v))
	})
}

// PrettyNameIn applies the In predicate on the "pretty_name" field.
func PrettyNameIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrettyName), v...))
	})
}

// PrettyNameNotIn applies the NotIn predicate on the "pretty_name" field.
func PrettyNameNotIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrettyName), v...))
	})
}

// PrettyNameGT applies the GT predicate on the "pretty_name" field.
func PrettyNameGT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrettyName), v))
	})
}

// PrettyNameGTE applies the GTE predicate on the "pretty_name" field.
func PrettyNameGTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrettyName), v))
	})
}

// PrettyNameLT applies the LT predicate on the "pretty_name" field.
func PrettyNameLT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrettyName), v))
	})
}

// PrettyNameLTE applies the LTE predicate on the "pretty_name" field.
func PrettyNameLTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrettyName), v))
	})
}

// PrettyNameContains applies the Contains predicate on the "pretty_name" field.
func PrettyNameContains(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrettyName), v))
	})
}

// PrettyNameHasPrefix applies the HasPrefix predicate on the "pretty_name" field.
func PrettyNameHasPrefix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrettyName), v))
	})
}

// PrettyNameHasSuffix applies the HasSuffix predicate on the "pretty_name" field.
func PrettyNameHasSuffix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrettyName), v))
	})
}

// PrettyNameEqualFold applies the EqualFold predicate on the "pretty_name" field.
func PrettyNameEqualFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrettyName), v))
	})
}

// PrettyNameContainsFold applies the ContainsFold predicate on the "pretty_name" field.
func PrettyNameContainsFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrettyName), v))
	})
}

// IsEnabledEQ applies the EQ predicate on the "is_enabled" field.
func IsEnabledEQ(v bool) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsEnabled), v))
	})
}

// IsEnabledNEQ applies the NEQ predicate on the "is_enabled" field.
func IsEnabledNEQ(v bool) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsEnabled), v))
	})
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageURL), v))
	})
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageURL), v))
	})
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldImageURL), v...))
	})
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldImageURL), v...))
	})
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImageURL), v))
	})
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImageURL), v))
	})
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImageURL), v))
	})
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImageURL), v))
	})
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImageURL), v))
	})
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImageURL), v))
	})
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImageURL), v))
	})
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImageURL), v))
	})
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImageURL), v))
	})
}

// ThumbnailURLEQ applies the EQ predicate on the "thumbnail_url" field.
func ThumbnailURLEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLNEQ applies the NEQ predicate on the "thumbnail_url" field.
func ThumbnailURLNEQ(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLIn applies the In predicate on the "thumbnail_url" field.
func ThumbnailURLIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldThumbnailURL), v...))
	})
}

// ThumbnailURLNotIn applies the NotIn predicate on the "thumbnail_url" field.
func ThumbnailURLNotIn(vs ...string) predicate.Chain {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldThumbnailURL), v...))
	})
}

// ThumbnailURLGT applies the GT predicate on the "thumbnail_url" field.
func ThumbnailURLGT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLGTE applies the GTE predicate on the "thumbnail_url" field.
func ThumbnailURLGTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLLT applies the LT predicate on the "thumbnail_url" field.
func ThumbnailURLLT(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLLTE applies the LTE predicate on the "thumbnail_url" field.
func ThumbnailURLLTE(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLContains applies the Contains predicate on the "thumbnail_url" field.
func ThumbnailURLContains(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLHasPrefix applies the HasPrefix predicate on the "thumbnail_url" field.
func ThumbnailURLHasPrefix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLHasSuffix applies the HasSuffix predicate on the "thumbnail_url" field.
func ThumbnailURLHasSuffix(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLEqualFold applies the EqualFold predicate on the "thumbnail_url" field.
func ThumbnailURLEqualFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThumbnailURL), v))
	})
}

// ThumbnailURLContainsFold applies the ContainsFold predicate on the "thumbnail_url" field.
func ThumbnailURLContainsFold(v string) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThumbnailURL), v))
	})
}

// HasChainProposals applies the HasEdge predicate on the "chain_proposals" edge.
func HasChainProposals() predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChainProposalsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChainProposalsTable, ChainProposalsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChainProposalsWith applies the HasEdge predicate on the "chain_proposals" edge with a given conditions (other predicates).
func HasChainProposalsWith(preds ...predicate.ChainProposal) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChainProposalsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChainProposalsTable, ChainProposalsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTelegramChats applies the HasEdge predicate on the "telegram_chats" edge.
func HasTelegramChats() predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TelegramChatsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TelegramChatsTable, TelegramChatsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTelegramChatsWith applies the HasEdge predicate on the "telegram_chats" edge with a given conditions (other predicates).
func HasTelegramChatsWith(preds ...predicate.TelegramChat) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TelegramChatsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TelegramChatsTable, TelegramChatsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDiscordChannels applies the HasEdge predicate on the "discord_channels" edge.
func HasDiscordChannels() predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiscordChannelsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DiscordChannelsTable, DiscordChannelsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiscordChannelsWith applies the HasEdge predicate on the "discord_channels" edge with a given conditions (other predicates).
func HasDiscordChannelsWith(preds ...predicate.DiscordChannel) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiscordChannelsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, DiscordChannelsTable, DiscordChannelsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Chain) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Chain) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Chain) predicate.Chain {
	return predicate.Chain(func(s *sql.Selector) {
		p(s.Not())
	})
}
