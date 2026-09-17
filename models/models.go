package models

import (
	"log"

	"todo/migrations"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

// Migrate aplica las migraciones pendientes sobre la conexión DB.
func Migrate() error {
	box, err := pop.NewMigrationBox(migrations.FS(), DB)
	if err != nil {
		return err
	}
	return box.Up()
}
