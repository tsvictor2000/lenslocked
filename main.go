package main

import (
	"fmt"
	"lenslocked/controllers"
	"lenslocked/views"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func main() {
	// парсинг шаблонов
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Stearting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
