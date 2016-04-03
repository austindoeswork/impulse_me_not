package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"database/sql"
	"fmt"
	"github.com/ConnorFoody/modl"
	//"github.com/jmoiron/modl"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"os"
	"strings"
	"time"
)

const (
	backupsDir    = "db/backups"
	migrationsDir = "db/migrations"
)

// Database is the db for code
type Database struct {
	path string

	db                    *sql.DB
	userMap               *modl.DbMap
	objectMap             *modl.DbMap
	userObjectRelationMap *modl.DbMap
}

// OpenDatabase opens the SQLite database at the given path, creating
// it if it doesn't exist, and runs any pending migrations
func OpenDatabase(path string) (*Database, error) {
	// use goose to auto-create the DB
	dbDriver := goose.DBDriver{"sqlite3",
		path,
		"github.com/mattn/go-sqlite3",
		&goose.Sqlite3Dialect{},
	}

	dbConf := goose.DBConf{MigrationsDir: migrationsDir,
		Env:    "prod",
		Driver: dbDriver,
	}

	target, err := goose.GetMostRecentDBVersion(migrationsDir)
	if err != nil {
		fmt.Println("a")
		return nil, err
	}

	err = goose.RunMigrations(&dbConf, migrationsDir, target)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	database := Database{path: path, db: db}
	database.mapTables()

	return &database, nil
}

// Close the DB
func (database *Database) Close() {
	database.db.Close()
}

// Backup creates a copy of the current database and saves it to the
// backups directory
func (database *Database) Backup(reason string) error {

	err := os.MkdirAll(backupsDir, 0755)
	if err != nil {
		return err
	}

	name := "testing"
	filename := fmt.Sprintf("%s/%s_%s_%s.db",
		backupsDir, strings.Replace(name, " ", "_", -1),
		time.Now().Format("20060102150405"), reason)

	src, err := os.Open(database.path)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return err
	}
	return nil
}

// Sets up table-object associations.
func (database *Database) mapTables() {
	dialect := new(modl.SqliteDialect)

	database.objectMap = modl.NewDbMap(database.db, dialect)
	database.objectMap.AddTableWithName(Object{},
		"object").SetKeys(true, "ID")

	database.userMap = modl.NewDbMap(database.db, dialect)
	database.userMap.AddTableWithName(User{},
		"user").SetKeys(true, "ID")

	database.userObjectRelationMap = modl.NewDbMap(database.db, dialect)
	database.userObjectRelationMap.AddTableWithName(UserObjectRelation{},
		"user_object_relation").SetKeys(true, "ID")
}
