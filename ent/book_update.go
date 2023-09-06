// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/TestGRPC/ent/author"
	"github.com/wtkeqrf0/TestGRPC/ent/book"
	"github.com/wtkeqrf0/TestGRPC/ent/predicate"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// AddAuthorIDs adds the "authors" edge to the Author entity by IDs.
func (bu *BookUpdate) AddAuthorIDs(ids ...int) *BookUpdate {
	bu.mutation.AddAuthorIDs(ids...)
	return bu
}

// AddAuthors adds the "authors" edges to the Author entity.
func (bu *BookUpdate) AddAuthors(a ...*Author) *BookUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.AddAuthorIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// ClearAuthors clears all "authors" edges to the Author entity.
func (bu *BookUpdate) ClearAuthors() *BookUpdate {
	bu.mutation.ClearAuthors()
	return bu
}

// RemoveAuthorIDs removes the "authors" edge to Author entities by IDs.
func (bu *BookUpdate) RemoveAuthorIDs(ids ...int) *BookUpdate {
	bu.mutation.RemoveAuthorIDs(ids...)
	return bu
}

// RemoveAuthors removes "authors" edges to Author entities.
func (bu *BookUpdate) RemoveAuthors(a ...*Author) *BookUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.RemoveAuthorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeString))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if bu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !bu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// AddAuthorIDs adds the "authors" edge to the Author entity by IDs.
func (buo *BookUpdateOne) AddAuthorIDs(ids ...int) *BookUpdateOne {
	buo.mutation.AddAuthorIDs(ids...)
	return buo
}

// AddAuthors adds the "authors" edges to the Author entity.
func (buo *BookUpdateOne) AddAuthors(a ...*Author) *BookUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.AddAuthorIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// ClearAuthors clears all "authors" edges to the Author entity.
func (buo *BookUpdateOne) ClearAuthors() *BookUpdateOne {
	buo.mutation.ClearAuthors()
	return buo
}

// RemoveAuthorIDs removes the "authors" edge to Author entities by IDs.
func (buo *BookUpdateOne) RemoveAuthorIDs(ids ...int) *BookUpdateOne {
	buo.mutation.RemoveAuthorIDs(ids...)
	return buo
}

// RemoveAuthors removes "authors" edges to Author entities.
func (buo *BookUpdateOne) RemoveAuthors(a ...*Author) *BookUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.RemoveAuthorIDs(ids...)
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeString))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if buo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !buo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}