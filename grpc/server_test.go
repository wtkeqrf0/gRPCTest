package grpc

import (
	"context"
	"github.com/wtkeqrf0/TestGRPC/mysql"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
)

var (
	dsn      string
	realTest bool
)

func init() {
	dsn = os.Getenv("MYSQL_DSN")

	if dsn != "" {
		realTest = true
	}
}

func runGRPCServer(ctx context.Context) (func(context.Context, string) (net.Conn, error), error) {
	cl, err := mysql.Open(ctx, dsn)
	if err != nil {
		return nil, err
	}

	srv := NewServer(cl)

	lis := bufconn.Listen(1024 * 1024)

	go func() {
		defer cl.Close()
		defer srv.Stop()
		if err = srv.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}, nil
}
