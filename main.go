package main

import (
	"fmt"
	"lenslocked/controllers"
	"lenslocked/models"
	"lenslocked/templates"
	"lenslocked/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	// парсинг шаблонов
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Get("/dogovors", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "doglist.gohtml", "tailwind.gohtml"))))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/signin", usersC.SignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Stearting the server on :3000...")

	csrfKey := "gjtnv57dkr9cbw3KEHN1PAMX54kd84nb"
	csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	http.ListenAndServe(":3000", csrfMw(r))

}
