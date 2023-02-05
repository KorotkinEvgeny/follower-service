package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func Migrate(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		log.Panicf("failed to init migrations: %s", err.Error())
		return err
	}

	err = m.Up()

	if err != nil {
		if err.Error() == "no change" {
			log.Info("no new migrations to run")
			return nil
		} else {
			log.Panicf("failed to init migrations: %s", err.Error())
			return err
		}
	}

	return nil
}
