package greetings

import (
	"errors"
	"fmt"
	"math/rand"

	"learn/exceptions"
)

func ValidateName(name string) (bool, error) {
	if name == "" {
		return false, errors.New("no name specified")
	}

	return true, nil
}

func randomMsg() string {
	// slice of message patterns
	msgs := []string{
		"Hello %v, how its going",
		"Yo, whats up to you %v",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	randIndx := rand.Intn(len(msgs))
	return msgs[randIndx]
}

func Greet(name string) (string, error) {
	// handle errors
	if name == "" {
		return "", errors.New("no name specified")
	}

	message := fmt.Sprintf(randomMsg(), name)
	return message, nil
}

func GreetMany(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greet(name)
		exceptions.HandleErr(err)
		messages[name] = message
	}

	return messages, nil
}
