package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Ошибка парсинга %v", err)
		http.Error(w, "Произошла ошибка парсинга шаблона", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка выполнения шаблона %v", err)
		http.Error(w, "Произошла ошибка выполнения шаблона", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplpath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplpath)
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
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, `<h1>FAQ page</h1>
	<ul>
		<li><b>Куда бы вы хотели поехать отдыхать летом?</b> В Турцию</li>
		<li><b>Сколько вас человек?</b> Нас 4 человека</li>
		<li><b>Сколько будет стоит путевка?</b> Путевка на 4-х стоит 2 млн. тенге</li>
	</ul>
	`)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(pageNotFoundHandler)

	fmt.Println("Stearting the server on :3000...")

	http.ListenAndServe(":3000", r)
}
