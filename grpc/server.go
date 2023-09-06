package grpc

import (
	"context"
	"github.com/wtkeqrf0/TestGRPC/ent"
	pb "github.com/wtkeqrf0/TestGRPC/proto"
	"google.golang.org/grpc"
	"net"
)

// DBClient means a structure capable of obtaining data about books and their authors.
type DBClient interface {
	AuthorsByBook(ctx context.Context, bookTitle string) (ent.Authors, error)
	BooksByAuthor(ctx context.Context, authorId int) (ent.Books, error)
}

// Server struct implements the functions of the RPC server
type Server struct {
	cl  DBClient
	srv *grpc.Server
	pb.UnimplementedLibraryServiceServer
}

// NewServer creates a new RPC server.
func NewServer(client DBClient) *Server {
	return &Server{
		srv: grpc.NewServer(),
		cl:  client,
	}
}

// Serve RPC on specified net.Listener.
func (s *Server) Serve(lis net.Listener) error {

	pb.RegisterLibraryServiceServer(s.srv, s)

	return s.srv.Serve(lis)
}

// Stop the server by the graceful shutdown.
func (s *Server) Stop() {
	s.srv.GracefulStop()
}
