package mysql

import (
	"context"
	"testing"
)

func TestClient_AuthorsByBook(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	cl, err := Open(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer cl.Close()

	author, err := cl.cl.Author.Create().SetFirstName("first").SetSurname("last").Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer cl.cl.Author.DeleteOne(author).Exec(ctx)

	t.Log(author)

	book, err := cl.cl.Book.Create().SetID("My Big Adventure").AddAuthors(author).Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer cl.cl.Book.DeleteOne(book).Exec(ctx)

	authors, err := cl.AuthorsByBook(ctx, book.ID)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range authors {
		if v.ID == author.ID {
			return
		}
	}
	t.Fatalf("author was not found: %+v", author)
}

func TestClient_BooksByAuthor(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	cl, err := Open(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer cl.Close()

	a, err := cl.cl.Author.Create().SetFirstName("first").SetSurname("last").Save(ctx)
	if err != nil {
		t.Error(err)
	}
	defer cl.cl.Author.DeleteOne(a).Exec(ctx)

	b, err := cl.cl.Book.Create().SetID("My Big Adventure").AddAuthors(a).Save(ctx)
	if err != nil {
		t.Error(err)
	}
	defer cl.cl.Book.DeleteOne(b).Exec(ctx)

	t.Log(b)

	books, err := cl.BooksByAuthor(ctx, a.ID)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range books {
		if v.ID == b.ID {
			return
		}
	}
	t.Fatalf("author was not found: %+v", a)
}
