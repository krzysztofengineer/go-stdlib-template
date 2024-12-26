package database

import (
	"database/sql"
	"log/slog"
)

type Migration struct {
	Name    string
	Version int
	Query   string
	Args    []any
}

var migrations []Migration

func PushMigration(name string, version int, query string, args ...any) {
	migrations = append(migrations, Migration{
		Name:    name,
		Version: version,
		Query:   query,
		Args:    args,
	})
}

func MustMigrate(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS migrations(
			version int NOT NULL,
			name varchar(255) NOT NULL
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	query = `
		SELECT COALESCE(MAX(version), 0) as version FROM migrations;
	`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	rows.Next()
	defer rows.Close()

	var version int
	err = rows.Scan(&version)
	if err != nil {
		panic(err)
	}

	rows.Close()

	for _, m := range migrations {
		if version >= m.Version {
			continue
		}

		_, err := db.Exec(m.Query)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			INSERT INTO migrations (version, name) VALUES (?, ?);
		`, m.Version, m.Name)

		if err != nil {
			panic(err)
		}

		version++
		slog.Info("migrated", "name", m.Name)
	}

	slog.Info("migrations complete", "version", version)
}
