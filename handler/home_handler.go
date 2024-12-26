package handler

import (
	"database/sql"
	"net/http"
	"template/view"
)

type HomeHandler struct {
	DB *sql.DB
}

func NewHomeHandler(db *sql.DB) *HomeHandler {
	return &HomeHandler{
		DB: db,
	}
}

func (*HomeHandler) Home() http.HandlerFunc {
	t := view.MustNew("layouts/base.html", "pages/home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}

func (*HomeHandler) Test() http.HandlerFunc {
	t := view.MustNew("pages/home.html")

	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "test", nil)
	}
}
