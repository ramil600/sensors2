package main

import (
	_ "embed"
	"log"
	"os"

	"github.com/ardanlabs/darwin"
	"github.com/ramil600/sensors2/business/user/db"
)

//go:embed schema/schema.sql
var schema string

func main() {

	db, err := db.Open(db.DBcfg)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	driver, err := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})
	if err != nil {
		log.Fatal(err)
	}
	d := darwin.New(driver, darwin.ParseMigrations(schema))
	if err := d.Migrate(); err != nil {
		log.Fatal(err)
	}
}
