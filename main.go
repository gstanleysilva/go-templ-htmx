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

var userIDCount int = 2

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	r.Get("/", GetHome)
	r.Get("/users/{userID}", DisplayUser)
	r.Delete("/users/{userID}", RemoveUser)
	r.Post("/users", AddUser)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	userIDCount++
	user := models.User{
		ID:   userIDCount,
		Name: fmt.Sprintf("User %d", userIDCount),
		Age:  userIDCount * 8,
	}
	oData.AddUser(user)
	components.UserList(oData.Users).Render(r.Context(), w)
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

func DisplayUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := oData.GetUser(userId)
	components.DisplayUser(*user).Render(r.Context(), w)
}
