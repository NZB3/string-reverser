package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	ReversedString string
}

func HandlerMain(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./page/index.html")
	if err != nil {
		log.Fatal("error parse template")
	}

	var stringToReverse string
	if r.Method == "POST" {
		r.ParseForm()
		stringToReverse = r.Form.Get("stringToReverse")
		log.Println("string to reverse: ", stringToReverse)
	}
	reversedString := reverseString(stringToReverse)
	log.Println("reversed string: ", reversedString)

	data := PageData{
		ReversedString: reversedString,
	}

	log.Println("data: ", data)
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", HandlerMain)
	http.ListenAndServe(":8080", nil)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
