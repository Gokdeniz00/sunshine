package main

import (
	//"net/http"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	webpage, _ := loadPage("page2")
	fmt.Fprintf(w, string(webpage.Body))
}
func randNumber() int {
	Source := rand.NewSource(time.Now().UnixNano())
	Generator := rand.New(Source)
	Number := Generator.Intn(5)
	return Number
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
