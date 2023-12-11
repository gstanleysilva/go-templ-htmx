package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gstanleysilva/go-templ-htmx-bulma/templates/pages"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home().Render(r.Context(), w)
	})

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}
