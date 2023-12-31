// Code generated by ent, DO NOT EDIT.

package author

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wtkeqrf0/TestGRPC/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldID, id))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldFirstName, v))
}

// Surname applies equality check predicate on the "surname" field. It's identical to SurnameEQ.
func Surname(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldSurname, v))
}

// Patronymic applies equality check predicate on the "patronymic" field. It's identical to PatronymicEQ.
func Patronymic(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldPatronymic, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Author {
	return predicate.Author(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Author {
	return predicate.Author(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Author {
	return predicate.Author(sql.FieldContainsFold(FieldFirstName, v))
}

// SurnameEQ applies the EQ predicate on the "surname" field.
func SurnameEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldSurname, v))
}

// SurnameNEQ applies the NEQ predicate on the "surname" field.
func SurnameNEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldSurname, v))
}

// SurnameIn applies the In predicate on the "surname" field.
func SurnameIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldSurname, vs...))
}

// SurnameNotIn applies the NotIn predicate on the "surname" field.
func SurnameNotIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldSurname, vs...))
}

// SurnameGT applies the GT predicate on the "surname" field.
func SurnameGT(v string) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldSurname, v))
}

// SurnameGTE applies the GTE predicate on the "surname" field.
func SurnameGTE(v string) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldSurname, v))
}

// SurnameLT applies the LT predicate on the "surname" field.
func SurnameLT(v string) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldSurname, v))
}

// SurnameLTE applies the LTE predicate on the "surname" field.
func SurnameLTE(v string) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldSurname, v))
}

// SurnameContains applies the Contains predicate on the "surname" field.
func SurnameContains(v string) predicate.Author {
	return predicate.Author(sql.FieldContains(FieldSurname, v))
}

// SurnameHasPrefix applies the HasPrefix predicate on the "surname" field.
func SurnameHasPrefix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasPrefix(FieldSurname, v))
}

// SurnameHasSuffix applies the HasSuffix predicate on the "surname" field.
func SurnameHasSuffix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasSuffix(FieldSurname, v))
}

// SurnameEqualFold applies the EqualFold predicate on the "surname" field.
func SurnameEqualFold(v string) predicate.Author {
	return predicate.Author(sql.FieldEqualFold(FieldSurname, v))
}

// SurnameContainsFold applies the ContainsFold predicate on the "surname" field.
func SurnameContainsFold(v string) predicate.Author {
	return predicate.Author(sql.FieldContainsFold(FieldSurname, v))
}

// PatronymicEQ applies the EQ predicate on the "patronymic" field.
func PatronymicEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldEQ(FieldPatronymic, v))
}

// PatronymicNEQ applies the NEQ predicate on the "patronymic" field.
func PatronymicNEQ(v string) predicate.Author {
	return predicate.Author(sql.FieldNEQ(FieldPatronymic, v))
}

// PatronymicIn applies the In predicate on the "patronymic" field.
func PatronymicIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldIn(FieldPatronymic, vs...))
}

// PatronymicNotIn applies the NotIn predicate on the "patronymic" field.
func PatronymicNotIn(vs ...string) predicate.Author {
	return predicate.Author(sql.FieldNotIn(FieldPatronymic, vs...))
}

// PatronymicGT applies the GT predicate on the "patronymic" field.
func PatronymicGT(v string) predicate.Author {
	return predicate.Author(sql.FieldGT(FieldPatronymic, v))
}

// PatronymicGTE applies the GTE predicate on the "patronymic" field.
func PatronymicGTE(v string) predicate.Author {
	return predicate.Author(sql.FieldGTE(FieldPatronymic, v))
}

// PatronymicLT applies the LT predicate on the "patronymic" field.
func PatronymicLT(v string) predicate.Author {
	return predicate.Author(sql.FieldLT(FieldPatronymic, v))
}

// PatronymicLTE applies the LTE predicate on the "patronymic" field.
func PatronymicLTE(v string) predicate.Author {
	return predicate.Author(sql.FieldLTE(FieldPatronymic, v))
}

// PatronymicContains applies the Contains predicate on the "patronymic" field.
func PatronymicContains(v string) predicate.Author {
	return predicate.Author(sql.FieldContains(FieldPatronymic, v))
}

// PatronymicHasPrefix applies the HasPrefix predicate on the "patronymic" field.
func PatronymicHasPrefix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasPrefix(FieldPatronymic, v))
}

// PatronymicHasSuffix applies the HasSuffix predicate on the "patronymic" field.
func PatronymicHasSuffix(v string) predicate.Author {
	return predicate.Author(sql.FieldHasSuffix(FieldPatronymic, v))
}

// PatronymicIsNil applies the IsNil predicate on the "patronymic" field.
func PatronymicIsNil() predicate.Author {
	return predicate.Author(sql.FieldIsNull(FieldPatronymic))
}

// PatronymicNotNil applies the NotNil predicate on the "patronymic" field.
func PatronymicNotNil() predicate.Author {
	return predicate.Author(sql.FieldNotNull(FieldPatronymic))
}

// PatronymicEqualFold applies the EqualFold predicate on the "patronymic" field.
func PatronymicEqualFold(v string) predicate.Author {
	return predicate.Author(sql.FieldEqualFold(FieldPatronymic, v))
}

// PatronymicContainsFold applies the ContainsFold predicate on the "patronymic" field.
func PatronymicContainsFold(v string) predicate.Author {
	return predicate.Author(sql.FieldContainsFold(FieldPatronymic, v))
}

// HasBooks applies the HasEdge predicate on the "books" edge.
func HasBooks() predicate.Author {
	return predicate.Author(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BooksTable, BooksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBooksWith applies the HasEdge predicate on the "books" edge with a given conditions (other predicates).
func HasBooksWith(preds ...predicate.Book) predicate.Author {
	return predicate.Author(func(s *sql.Selector) {
		step := newBooksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Author) predicate.Author {
	return predicate.Author(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Author) predicate.Author {
	return predicate.Author(func(s *sql.Selector) {
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
func Not(p predicate.Author) predicate.Author {
	return predicate.Author(func(s *sql.Selector) {
		p(s.Not())
	})
}
