package main

import (
	//"net/http"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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
	Source := rand.NewSource(time.Now().UnixNano())
	Generator := rand.New(Source)
	Number := Generator.Intn(4) + 1
	page_name := "page" + strconv.Itoa(Number)
	webpage, _ := loadPage(page_name)
	fmt.Fprintf(w, string(webpage.Body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
