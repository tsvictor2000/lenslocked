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
		views.Must(views.ParseFS(templates.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.Get("/dogovors", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "doglist.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Stearting the server on :3000...")

	http.ListenAndServe(":3000", r)

}
