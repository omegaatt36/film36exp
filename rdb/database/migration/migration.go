package migration

import (
	"github.com/omegaatt36/film36exp/logging"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var migrationOptions = gormigrate.Options{
	UseTransaction: true,
}

// Migrator runs migration.
type Migrator struct {
	db         *gorm.DB
	migrations []*gormigrate.Migration
}

// NewMigrator creates migrator.
func NewMigrator(db *gorm.DB, migrations []*gormigrate.Migration) *Migrator {
	return &Migrator{db: db, migrations: migrations}
}

func (m *Migrator) upgradeLatestMigrate() error {
	mg := gormigrate.New(m.db, &migrationOptions, m.migrations)
	err := mg.Migrate()
	if err != nil {
		return err
	}
	return nil
}

// Upgrade upgrades db schema version.
func (m *Migrator) Upgrade() error {
	if len(m.migrations) == 0 {
		return nil
	}

	if err := m.upgradeLatestMigrate(); err != nil {
		return errors.Wrap(err, "update to latest failed")
	}

	logging.Infof("upgraded to version \"%s\"", m.migrations[len(m.migrations)-1].ID)
	return nil
}

// Rollback rollbacks the last migration.
func (m *Migrator) Rollback() error {
	mg := gormigrate.New(m.db, &migrationOptions, m.migrations)
	if err := mg.RollbackLast(); err != nil {
		return errors.Wrap(err, "rollback to last failed")
	}

	logging.Info("rollback to last version success")
	return nil
}
