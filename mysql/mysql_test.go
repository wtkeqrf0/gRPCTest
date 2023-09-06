package mysql

import "os"

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
