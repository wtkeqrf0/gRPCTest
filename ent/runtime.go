// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/wtkeqrf0/TestGRPC/ent/author"
	"github.com/wtkeqrf0/TestGRPC/ent/book"
	"github.com/wtkeqrf0/TestGRPC/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authorFields := schema.Author{}.Fields()
	_ = authorFields
	// authorDescFirstName is the schema descriptor for first_name field.
	authorDescFirstName := authorFields[0].Descriptor()
	// author.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	author.FirstNameValidator = authorDescFirstName.Validators[0].(func(string) error)
	// authorDescSurname is the schema descriptor for surname field.
	authorDescSurname := authorFields[1].Descriptor()
	// author.SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	author.SurnameValidator = authorDescSurname.Validators[0].(func(string) error)
	// authorDescPatronymic is the schema descriptor for patronymic field.
	authorDescPatronymic := authorFields[2].Descriptor()
	// author.PatronymicValidator is a validator for the "patronymic" field. It is called by the builders before save.
	author.PatronymicValidator = authorDescPatronymic.Validators[0].(func(string) error)
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescID is the schema descriptor for id field.
	bookDescID := bookFields[0].Descriptor()
	// book.IDValidator is a validator for the "id" field. It is called by the builders before save.
	book.IDValidator = bookDescID.Validators[0].(func(string) error)
}