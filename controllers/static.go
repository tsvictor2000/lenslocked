package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Куда бы вы хотели поехать отдыхать летом?",
			Answer:   "В Турцию",
		},
		{
			Question: "Сколько вас человек?",
			Answer:   "Нас 4 человека",
		},
		{
			Question: "Сколько будет стоит путевка?",
			Answer:   "Путевка на 4-х стоит 2 млн. тенге",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
