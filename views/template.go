package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("ошибка парсинга: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("Ошибка выполнения шаблона %v", err)
		http.Error(w, "Произошла ошибка выполнения шаблона", http.StatusInternalServerError)
		return
	}
}
