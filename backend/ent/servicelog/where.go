// Code generated by ent, DO NOT EDIT.

package servicelog

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ochanoco/proxy/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Headers applies equality check predicate on the "headers" field. It's identical to HeadersEQ.
func Headers(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeaders), v))
	})
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// HeadersEQ applies the EQ predicate on the "headers" field.
func HeadersEQ(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeaders), v))
	})
}

// HeadersNEQ applies the NEQ predicate on the "headers" field.
func HeadersNEQ(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeaders), v))
	})
}

// HeadersIn applies the In predicate on the "headers" field.
func HeadersIn(vs ...string) predicate.ServiceLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHeaders), v...))
	})
}

// HeadersNotIn applies the NotIn predicate on the "headers" field.
func HeadersNotIn(vs ...string) predicate.ServiceLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHeaders), v...))
	})
}

// HeadersGT applies the GT predicate on the "headers" field.
func HeadersGT(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHeaders), v))
	})
}

// HeadersGTE applies the GTE predicate on the "headers" field.
func HeadersGTE(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHeaders), v))
	})
}

// HeadersLT applies the LT predicate on the "headers" field.
func HeadersLT(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHeaders), v))
	})
}

// HeadersLTE applies the LTE predicate on the "headers" field.
func HeadersLTE(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHeaders), v))
	})
}

// HeadersContains applies the Contains predicate on the "headers" field.
func HeadersContains(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHeaders), v))
	})
}

// HeadersHasPrefix applies the HasPrefix predicate on the "headers" field.
func HeadersHasPrefix(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHeaders), v))
	})
}

// HeadersHasSuffix applies the HasSuffix predicate on the "headers" field.
func HeadersHasSuffix(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHeaders), v))
	})
}

// HeadersEqualFold applies the EqualFold predicate on the "headers" field.
func HeadersEqualFold(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHeaders), v))
	})
}

// HeadersContainsFold applies the ContainsFold predicate on the "headers" field.
func HeadersContainsFold(v string) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHeaders), v))
	})
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBody), v))
	})
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...[]byte) predicate.ServiceLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBody), v...))
	})
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...[]byte) predicate.ServiceLog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBody), v...))
	})
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBody), v))
	})
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBody), v))
	})
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBody), v))
	})
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v []byte) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBody), v))
	})
}

// BodyIsNil applies the IsNil predicate on the "body" field.
func BodyIsNil() predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBody)))
	})
}

// BodyNotNil applies the NotNil predicate on the "body" field.
func BodyNotNil() predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBody)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ServiceLog) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ServiceLog) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
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
func Not(p predicate.ServiceLog) predicate.ServiceLog {
	return predicate.ServiceLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
