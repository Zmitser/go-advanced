package main

import "fmt"

type englishBot struct {

}
type spanishBot struct {
	
}

type bot interface {
	getGreeting() string
}

func main() {
	eb:=englishBot{}
	sb:=spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
	
}

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}

//func printGreeting(bot englishBot) {
//	fmt.Println(bot.getGreeting())
//}
//
//func printGreeting(bot spanishBot) {
//	fmt.Println(bot.getGreeting())
//}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (eb spanishBot) getGreeting() string {
	return "Hola!"
}