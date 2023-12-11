package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gstanleysilva/go-templ-htmx-bulma/templates/models"
	"github.com/gstanleysilva/go-templ-htmx-bulma/templates/pages"
)

var oData models.GlobalData = models.GlobalData{
	Admin: models.User{
		ID:   0,
		Name: "Admin",
		Age:  30,
	},
	Users: []models.User{
		{

			ID:   1,
			Name: "User 1",
			Age:  20,
		},
		{
			ID:   2,
			Name: "User 2",
			Age:  25,
		},
	},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home(oData).Render(r.Context(), w)
	})

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}
