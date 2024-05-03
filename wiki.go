package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct{
	Title string
	Body []byte
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func (p *Page)save() error{
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p,_:= loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}