package main

import "fmt"

type person struct {
	name string
	age  int
}

func neewPerson(name string) *person {
	p := person{name: name}
	p.age = 42 // как-то разыменывать не надо, только в С нао ))
	return &p
}

func main_5() {
	fmt.Println(person{"Bob", 20}) // позиционно

	fmt.Println(person{name: "Alice", age: 30}) // инициализируем по именам
	dog := struct {                             // ну или можно заглушить _
		name   string
		isGood bool
	}{
		"Rex",
		true}

	fmt.Println(dog)

}
