package mysql

import (
	"context"
	"github.com/wtkeqrf0/TestGRPC/ent"
)

// Client struct provides the ability to interact with the database.
type Client struct {
	cl *ent.Client
}

// AuthorsByBook method gets all the authors of the book.
func (c *Client) AuthorsByBook(ctx context.Context, bookTitle string) (ent.Authors, error) {
	b, err := c.cl.Book.Get(ctx, bookTitle)
	if err != nil {
		return nil, err
	}

	return b.QueryAuthors().All(ctx)
}

// BooksByAuthor method gets all the books of the author.
func (c *Client) BooksByAuthor(ctx context.Context, authorId int) (ent.Books, error) {
	ar, err := c.cl.Author.Get(ctx, authorId)
	if err != nil {
		return nil, err
	}

	return ar.QueryBooks().All(ctx)
}
