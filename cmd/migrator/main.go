package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", "up", "up/down migrations")
}

func main() {
	flag.Parse()
	dbURL := os.Getenv("dbURL")
	if dbURL == "" {
		panic("dbURL should be provided")
	}

	migratePath := os.Getenv("migratePath")
	if migratePath == "" {
		panic("migratePath should be provided")
	}

	m, err := migrate.New(fmt.Sprintf("file://%s", migratePath), dbURL)
	if err != nil {
		panic(err)
	}

	switch mode {
	case "up":
		err = m.Up()
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				err = nil
			}
		}
	default:
		err = m.Down()
	}
	if err != nil {
		panic(err)
	}
}
