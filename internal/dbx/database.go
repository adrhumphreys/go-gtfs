package dbx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"log"
)

const driverName = "mysql"
const dataSourceName = "root:@tcp(localhost:3306)/gtfs?parseTime=true"

func InitDB () {
	db, err := sqlx.Open(driverName, dataSourceName)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	performMigration(db)
}

func performMigration(db *sqlx.DB) {
	migrations := &migrate.MemoryMigrationSource{Migrations: GetMigrations()}

	// cheeky bee key drop tables and start again
	n, err := migrate.Exec(db.DB, "mysql", migrations, migrate.Down)
	n, err = migrate.Exec(db.DB, "mysql", migrations, migrate.Up)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("🛴 %v migration(s) performed\n", n)
}

func Connect() *sqlx.DB {
	return sqlx.MustConnect(driverName, dataSourceName)
}