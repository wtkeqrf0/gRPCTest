package grpc

import (
	"context"
	pb "github.com/wtkeqrf0/TestGRPC/proto"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_BooksByAuthor(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	dialer, err := runGRPCServer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewLibraryServiceClient(conn)

	resp, err := client.BooksByAuthor(ctx, &pb.GetAuthor{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Response: %+v", resp)
}

func TestServer_AuthorsByBook(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	ctx := context.Background()

	dialer, err := runGRPCServer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewLibraryServiceClient(conn)

	resp, err := client.AuthorsByBook(ctx, &pb.GetBook{Title: "Evening ogusok"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Response: %+v", resp)
}
