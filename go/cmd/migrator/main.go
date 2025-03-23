package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "sp", "./storage/sso.db", "path to storage")
	flag.StringVar(&migrationsPath, "mp", "./migrations", "path to migrations")
	flag.StringVar(&migrationsTable, "mt", "migrations", "name of migrations table")
	flag.Parse()

	if storagePath == "" {
		panic("sp (storage path) is empty")
	}
	if migrationsPath == "" {
		panic("mp (migrations path) is empty")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("nothing to migrate")
			return
		}
		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
