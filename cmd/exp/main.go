package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio string
}

func main() {
	fmt.Println("Experimental code!")

	user := User{
		Name: "Victor",
		Bio: `<script>alert("Hi")</script>`,
	}

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
