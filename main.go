package main

import (
	"log"
	"log/slog"
	"net/http"
	"template/view"
)

func main() {
	r := http.NewServeMux()

	t := view.Must("layouts/base.html", "pages/home.html")
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	})

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("listening on http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
