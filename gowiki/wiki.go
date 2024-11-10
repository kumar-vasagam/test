package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{fileName, content}, nil
}
