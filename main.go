package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gstanleysilva/go-templ-htmx-bulma/models"
	"github.com/gstanleysilva/go-templ-htmx-bulma/templates/components"
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

	r.Get("/", GetHome)
	r.Delete("/users/{userID}", RemoveUser)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	pages.Home(oData).Render(r.Context(), w)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oData.RemoveUser(id)
	components.UserList(oData.Users).Render(r.Context(), w)
}
