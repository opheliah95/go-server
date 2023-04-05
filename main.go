package main

import (
	"fmt"
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

}
