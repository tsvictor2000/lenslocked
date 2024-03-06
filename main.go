package main

import (
	"fmt"
	"lenslocked/controllers"
	"lenslocked/templates"
	"lenslocked/views"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	r.Get("/signup", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))))


	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Stearting the server on :3000...")

	http.ListenAndServe(":3000", r)

}
