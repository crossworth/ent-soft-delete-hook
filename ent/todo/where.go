// Code generated by entc, DO NOT EDIT.

package todo

import (
	"time"

	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DeletedTime applies equality check predicate on the "deleted_time" field. It's identical to DeletedTimeEQ.
func DeletedTime(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// DeletedTimeEQ applies the EQ predicate on the "deleted_time" field.
func DeletedTimeEQ(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedTime), v))
	})
}

// DeletedTimeNEQ applies the NEQ predicate on the "deleted_time" field.
func DeletedTimeNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedTime), v))
	})
}

// DeletedTimeIn applies the In predicate on the "deleted_time" field.
func DeletedTimeIn(vs ...time.Time) predicate.Todo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedTime), v...))
	})
}

// DeletedTimeNotIn applies the NotIn predicate on the "deleted_time" field.
func DeletedTimeNotIn(vs ...time.Time) predicate.Todo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedTime), v...))
	})
}

// DeletedTimeGT applies the GT predicate on the "deleted_time" field.
func DeletedTimeGT(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedTime), v))
	})
}

// DeletedTimeGTE applies the GTE predicate on the "deleted_time" field.
func DeletedTimeGTE(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedTime), v))
	})
}

// DeletedTimeLT applies the LT predicate on the "deleted_time" field.
func DeletedTimeLT(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedTime), v))
	})
}

// DeletedTimeLTE applies the LTE predicate on the "deleted_time" field.
func DeletedTimeLTE(v time.Time) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedTime), v))
	})
}

// DeletedTimeIsNil applies the IsNil predicate on the "deleted_time" field.
func DeletedTimeIsNil() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedTime)))
	})
}

// DeletedTimeNotNil applies the NotNil predicate on the "deleted_time" field.
func DeletedTimeNotNil() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedTime)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Todo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Todo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Todo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
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
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}