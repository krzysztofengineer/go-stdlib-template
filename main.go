package main

import (
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"template/view"
)

var (
	//go:embed static
	staticFS embed.FS
)

func main() {
	r := http.NewServeMux()

	t := view.Must("layouts/base.html", "pages/home.html")
	r.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	})

	r.HandleFunc("POST /test", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "test", nil)
	})

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
