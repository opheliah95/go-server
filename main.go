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

}
