package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func printPage(s Page) string {
	return s.Title
}

func (s *Page) savePage() error {
	filename := s.Title + ".txt"
	return os.WriteFile(filename, s.Body, 0600)
}

func loadPage (title string) (*Page, error) {
	filename := title + ".txt"
	body, err :=os.ReadFile(filename)
	fmt.Println("error is: ", err)
	if err != nil {
		return nil, err
	}
	return &Page{ Title: title, Body : body}, nil
}

func handler (w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "only get method is supported", http.StatusMethodNotAllowed)
	}
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hi there, You entered / %s for request!", r.URL.Path[1:])
}

func formHandler (w http.ResponseWriter, r *http.Request) {
	// only take post method
	if r.Method != "POST" {
		http.Error(w, "only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "There is some error in your form liks %v", err)
	} else{
		fmt.Fprintf(w, "post request successful!")
		fname := r.FormValue("fname")
		lname:= r.FormValue("lname")
		fmt.Fprintf(w, "first name: %s and last name is %s", fname, lname)
	}
	


}
func main() {
	body, _ := os.ReadFile("./static/index.html")
	page := &Page{
		"my title",
		body,
	}

	str := string(page.Body)
	fmt.Println(str)
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.savePage()

	p2, _ :=loadPage("TestPage")
	fmt.Println(string(p2.Body))

	// some file handling function
	http.HandleFunc("/", handler)
	http.HandleFunc("/form", formHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
