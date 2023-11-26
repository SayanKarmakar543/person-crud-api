package persondetails

import (
	"fmt"
)

type Person struct {
	Id     string `json:"ID"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Mobile string `json:"mobile"`
}

// var p []Person

func (p *Person) InsertPersonDetails() {
	fmt.Print("Enter Person ID: ")
	fmt.Scan(&p.Id)
	fmt.Print("Enter Person Name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter Person Age: ")
	fmt.Scan(&p.Age)
	fmt.Print("Enter Person Mobile: ")
	fmt.Scan(&p.Mobile)
}

func (p *Person) UpdatePersonDetails() {
	fmt.Print("Enter Person ID: ")
	fmt.Scan(&p.Id)
	fmt.Print("Enter Person Name: ")
	fmt.Scan(&p.Name)
	fmt.Print("Enter Person Age: ")
	fmt.Scan(&p.Age)
	fmt.Print("Enter Person Mobile: ")
	fmt.Scan(&p.Mobile)
}

func (p Person) DisplayPersonDetails() {
	fmt.Println(p)
}
