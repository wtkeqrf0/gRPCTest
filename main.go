package main

import (
	"context"
	"github.com/wtkeqrf0/TestGRPC/grpc"
	"github.com/wtkeqrf0/TestGRPC/mysql"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	client, err := mysql.Open(ctx, os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer(client)

	port := uint16(3000)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err = srv.Serve(bufconn.Listen(1024 * 1024)); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Server Started On Port %d\n", port)

	<-quit

	log.Println("Server Is Shutting Down ...")

	srv.Stop()

	if err = client.Close(); err != nil {
		log.Fatalf("PostgreSQL Connection Shutdown Failed: %v", err)
	}
}
