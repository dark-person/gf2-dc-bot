package main

import (
	"embed"

	"github.com/dark-person/lazydb"
)

//go:embed all:schema
var schema embed.FS

// Database instance to use.
var db *lazydb.LazyDB

// Setup database modules and create if necessary.
func setup() error {
	// Init db
	db = lazydb.New(
		lazydb.DbPath("data.db"),         // Database path
		lazydb.Migrate(schema, "schema"), // Migration schema location
		lazydb.BackupDir("./backup"),     // Set auto backup directory
	)

	// Connect to db, which will create file if necessary
	err := db.Connect()
	if err != nil {
		return err
	}

	// Migration performed
	_, err = db.Migrate()
	if err != nil {
		return err
	}

	return nil
}
