package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Quote struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Text   string `json:"text"`
}

var (
	theQuote Quote
)

func init() {
	theQuote = Quote{
		Author: "Rodney King",
		Title:  "",
		Text:   "Why can't we all just get along?",
	}
}

func (q Quote) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(theQuote)

	case "POST", "PUT":
		text, ok := r.URL.Query()["author"]
		if !ok || len(text[0]) < 1 {
			log.Println("Url Param 'text' is missing...")
			fmt.Fprint(w, "URL param 'text' is missing")
			return
		}
		author, ok := r.URL.Query()["author"]
		if !ok || len(author[0]) < 1 {
			log.Println("URL Param 'author' is missing")
		}
		title, ok := r.URL.Query()["title"]
		if !ok || len(title[0]) < 1 {
			log.Println("Url Param 'title' is missing")
		}

		theQuote.Author = strings.Join(author, ",")
		theQuote.Title = title[0]
		theQuote.Text = text[0]

	}

}
