package grpc

import (
	"context"
	pb "github.com/wtkeqrf0/TestGRPC/proto"
)

// BooksByAuthor method gets all the books of the author.
func (s *Server) BooksByAuthor(ctx context.Context, in *pb.GetAuthor) (*pb.Books, error) {
	books, err := s.cl.BooksByAuthor(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Books_Book, len(books))
	for i := range res {
		book := books[i]
		res[i] = &pb.Books_Book{
			Title: book.ID,
		}
	}

	return &pb.Books{Books: res}, nil
}

// AuthorsByBook method gets all the authors of the book.
func (s *Server) AuthorsByBook(ctx context.Context, in *pb.GetBook) (*pb.Authors, error) {
	authors, err := s.cl.AuthorsByBook(ctx, in.Title)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Authors_Author, len(authors))
	for i := range res {
		author := authors[i]
		res[i] = &pb.Authors_Author{
			FirstName:  author.FirstName,
			Surname:    author.Surname,
			Patronymic: author.Patronymic,
		}
	}

	return &pb.Authors{Authors: res}, nil
}
