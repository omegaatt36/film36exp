package database

import (
	"fmt"
	"log"
	"time"
)

var (
	// PostgresOpt is default connection option for postgres.
	PostgresOpt = ConnectOption{
		Dialect: "postgres",
		Host:    "localhost",
		DBName:  "testing",
		Port:    5432,
		User:    "tester",
		Pass:    "postgres",
	}

	// SQLiteOpt is shared in-memory database.
	SQLiteOpt = ConnectOption{
		Dialect: "sqlite3",
		Host:    "file::memory:?cache=shared",
	}
)

var count int

func randomDBName() string {
	count++
	return fmt.Sprintf("testing_%v_%d", time.Now().UnixNano(), count)
}

// TestingInitialize creates new db for testing.
func TestingInitialize(typ string, opt ConnectOption) {
	opt.Config.DisableForeignKeyConstraintWhenMigrating = true
	opt.Testing = true

	Initialize(typ, opt)

	if opt.Dialect != "postgres" {
		return
	}

	dbName := randomDBName()

	db := GetDB(typ)
	err := db.Exec("CREATE DATABASE " + dbName).Error
	if err != nil {
		log.Panicln(err)
	}

	opt.DBName = dbName
	log.Println("use db name:", dbName)

	store.Lock()
	for key, db := range store.dbs {
		if key != typ {
			continue
		}

		db.Close()
		delete(store.dbs, key)
	}
	store.Unlock()

	Initialize(typ, opt)
}

// TestingFinalize cleanups testing data.
func TestingFinalize() {
	Finalize()
}
