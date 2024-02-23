package main

import "fmt"

type Person struct {
	Name      string
	DateBirth string
	Weight    string
	Mail      []Email
}

type Email struct {
	To   string
	Name string
}

func main() {
	person1 := Person{
		Name:      "Putra Yasa",
		DateBirth: "01-01-1997",
		Weight:    "80",
		Mail: []Email{
			{
				To:   "putu.yasa2@gmail.com",
				Name: "Putra Yasa (PGPY)",
			},
		},
	}

	fmt.Println(person1.Name)
	fmt.Println(person1.Mail)

	persons := []Person{
		{Name: "Saya1", DateBirth: "01-01-1997", Weight: "80"},
		{Name: "Saya2", DateBirth: "01-01-1999", Weight: "60"},
	}

	fmt.Println(persons)
}
