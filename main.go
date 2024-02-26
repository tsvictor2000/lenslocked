package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my great site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href='mailto:3_7_11@mail.ru'>3_7_11@mail.ru</a></p>")
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
		<li><b>1 Куда бы вы хотели поехать отдыхать летом?</b> В Турцию</li>
		<li><b>2 Сколько вас человек?</b> Нас 4 человека</li>
		<li><b>3 Сколько будет стоит путевка?</b> Путевка на 4-х стоит 2 млн. тенге</li>
	</ul>
	`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, r.URL.Path)
// 	// fmt.Fprintln(w, r.URL.RawPath)

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		pageNotFoundHandler(w, r)
// 	}
// }

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		pageNotFoundHandler(w, r)
// 	}
// }

func main() {
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	// var router Router
	// router = pathHandler
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(pageNotFoundHandler)
	fmt.Println("Stearting the server on :3000...")
	// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	http.ListenAndServe(":3000", r)
}
