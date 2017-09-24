package main

import "fmt"

type contacts struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contacts
}

func main() {

	p := person{firstName: "Dmitry", lastName: "Borisovets"}
	d := person{"Zmitser", "Lisitsin", contacts{"1@2.by", 4234342}}
	var a person
	p.console()
	d.console()
	a.firstName = "Lev"
	a.lastName = "ITransition"
	a.updateFirstName("Serg")
	k := a.firstName
	fmt.Println(*&k)
}

func (p *person) updateFirstName(firstName string)  {
	p.firstName = firstName
}

func (p person) console()  {
	fmt.Printf("%+v\n", p)
}
