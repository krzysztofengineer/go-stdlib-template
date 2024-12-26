package main

import (
	"log"
	"log/slog"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	s := http.Server{
		Addr:    ":4000",
		Handler: r,
	}

	slog.Info("listening on http://localhost:4000")

	log.Fatal(s.ListenAndServe())
}
