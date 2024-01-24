package main

import (
	"fmt"

	"learn/exceptions"
	"learn/greetings"
)

func Say(msg string) {
	fmt.Println(msg)
}

func main() {
	names := []string{
		"Sagar", "Monokai", "Gilfoyle", "Summon",
	}

	msgs, err := greetings.GreetMany(names)
	exceptions.HandleErr(err)
	fmt.Println(msgs)
}
