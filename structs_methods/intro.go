package main

import (
	"fmt"
	"strings"
)

type struct1 struct {
	il  int
	fl  float32
	str string
}

type Person struct {
	firstName, lastName string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	ms := new(struct1)
	ms.il = 10
	ms.fl = 15.5
	ms.str = "Chris"
	fmt.Println(ms)

	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	fmt.Println(pers1)

	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward2"
	upPerson(pers2)
	fmt.Println(pers2)

	pers3 := &Person{"Chris", "Woodward"}
	upPerson(pers3)
	fmt.Println(pers3)

}
