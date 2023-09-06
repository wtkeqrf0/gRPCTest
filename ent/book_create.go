// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wtkeqrf0/TestGRPC/ent/author"
	"github.com/wtkeqrf0/TestGRPC/ent/book"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetID sets the "id" field.
func (bc *BookCreate) SetID(s string) *BookCreate {
	bc.mutation.SetID(s)
	return bc
}

// AddAuthorIDs adds the "authors" edge to the Author entity by IDs.
func (bc *BookCreate) AddAuthorIDs(ids ...int) *BookCreate {
	bc.mutation.AddAuthorIDs(ids...)
	return bc
}

// AddAuthors adds the "authors" edges to the Author entity.
func (bc *BookCreate) AddAuthors(a ...*Author) *BookCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bc.AddAuthorIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if v, ok := bc.mutation.ID(); ok {
		if err := book.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Book.id": %w`, err)}
		}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Book.ID type: %T", _spec.ID.Value)
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(book.Table, sqlgraph.NewFieldSpec(book.FieldID, field.TypeString))
	)
	_spec.OnConflict = bc.conflict
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := bc.mutation.AuthorsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Book.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (bc *BookCreate) OnConflict(opts ...sql.ConflictOption) *BookUpsertOne {
	bc.conflict = opts
	return &BookUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bc *BookCreate) OnConflictColumns(columns ...string) *BookUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BookUpsertOne{
		create: bc,
	}
}

type (
	// BookUpsertOne is the builder for "upsert"-ing
	//  one Book node.
	BookUpsertOne struct {
		create *BookCreate
	}

	// BookUpsert is the "OnConflict" setter.
	BookUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(book.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BookUpsertOne) UpdateNewValues() *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(book.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Book.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BookUpsertOne) Ignore() *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BookUpsertOne) DoNothing() *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BookCreate.OnConflict
// documentation for more info.
func (u *BookUpsertOne) Update(set func(*BookUpsert)) *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BookUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *BookUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BookCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BookUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BookUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BookUpsertOne.ID is not supported by MySQL driver. Use BookUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BookUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	builders []*BookCreate
	conflict []sql.ConflictOption
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Book.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (bcb *BookCreateBulk) OnConflict(opts ...sql.ConflictOption) *BookUpsertBulk {
	bcb.conflict = opts
	return &BookUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bcb *BookCreateBulk) OnConflictColumns(columns ...string) *BookUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BookUpsertBulk{
		create: bcb,
	}
}

// BookUpsertBulk is the builder for "upsert"-ing
// a bulk of Book nodes.
type BookUpsertBulk struct {
	create *BookCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(book.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BookUpsertBulk) UpdateNewValues() *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(book.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BookUpsertBulk) Ignore() *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BookUpsertBulk) DoNothing() *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BookCreateBulk.OnConflict
// documentation for more info.
func (u *BookUpsertBulk) Update(set func(*BookUpsert)) *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BookUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *BookUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BookCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BookCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BookUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
