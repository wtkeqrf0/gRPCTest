package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/wtkeqrf0/TestGRPC/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Open new connection, start stats recorder and create the tables.
func Open(ctx context.Context, dsn string) (*Client, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error occurred while opening MySQL connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to connect to the MySQL database: %v", err)
	}
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.MySQL, db)
	cl := ent.NewClient(ent.Driver(drv))

	if err = cl.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("tables initialization failed: %v", err)
	}

	return &Client{cl}, nil
}

func (c *Client) Close() error {
	return c.cl.Close()
}
