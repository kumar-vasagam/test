package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	simpleTemplate()
	RandomExamples()
	letter()
}

type Inventory struct {
	Count    int
	Material string
}

func simpleTemplate() {
	sweaters := Inventory{12, "Wool"}
	text := "{{.Count}} items were made of {{.Material}} \n"
	t, err := template.New("Test").Parse(text)
	if err != nil {
		fmt.Printf("Error in parsing %v \n", err.Error())
	}
	err = t.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	textWithWhiteSpaces := "So, {{- .Count}} items were made of {{.Material -}} ,isn't that great? \n"
	t2, err := template.New("T2").Parse(textWithWhiteSpaces)
	if err != nil {
		fmt.Printf("Error in parsing %v \n", err.Error())
	}
	err = t2.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

type Nature struct {
	Description string
}

func RandomExamples() {
	//To output "output"

	s := Nature{"Metal"}
	nice := "Nice!"

	t, _ := template.New("Example").Parse("Nature is {{.Description}} \n")
	t.Execute(os.Stdout, s)

	t1, _ := template.New("Example1").Parse("Nature is {{.}}")
	t1.Execute(os.Stdout, nice)
}

func letter() {
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	attendees := []Recipient{
		{"Tom", "Kettle", true},
		{"Peter", "Chocolates", true},
		{"Mary", "", true},
		{"Elizabeth", "", false},
	}

	letter := ` Dear {{.Name}}, 
				{{if .Attended}}
				It was a pleasure to see you at the wedding.
				{{with .Gift}}
				Thanks for the wonderful {{.}}
				{{end}}
				{{else}}
				Sorry you could not come
				{{end}}
				`
	t, err := template.New("Letter").Parse(letter)
	if err != nil {
		fmt.Printf("error in parsing %v", err.Error())
	}

	for _, a := range attendees {
		t.Execute(os.Stdout, a)
	}

}
