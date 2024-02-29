package main

import (
	"fmt"
	"lenslocked/controllers"
	"lenslocked/views"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("Ошибка парсинга %v", err)
		http.Error(w, "Произошла ошибка парсинга шаблона", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplpath)
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Page not found</h1>")

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func main() {
	// парсинг шаблонов
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(pageNotFoundHandler)

	fmt.Println("Stearting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
