package main

import (
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"template/database"
	"template/handler"
)

var (
	//go:embed static
	staticFS embed.FS
)

func main() {
	db := database.MustNew("file:test.db")

	database.PushMigration("0001_init", 1, `
		CREATE TABLE users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email VARCHAR(255) NOT NULL
		);
	`)
	database.PushMigration("0002_alter", 2, `
		ALTER TABLE users
		ADD name VARCHAR(255)
	`)
	database.MustMigrate(db)

	r := http.NewServeMux()

	homeHandler := handler.NewHomeHandler(db)
	r.HandleFunc("GET /{$}", homeHandler.Home())
	r.HandleFunc("POST /test", homeHandler.Test())

	staticSubDir, _ := fs.Sub(staticFS, "static")
	fs := http.FileServerFS(staticSubDir)
	r.Handle("/", http.StripPrefix("/static/", fs))

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("listening on http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
