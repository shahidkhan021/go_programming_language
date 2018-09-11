package main

import (
	"fmt"
)

/* syntax of interface

type nameOfInterface interface {
	functionToBeInterfaced(listofargumenttype) (list_of_return_type)
}

*/
type bot interface {
	getGreeting(int, string) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
