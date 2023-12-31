package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func SayHello() {
	fmt.Println("hello, world!")

}

func SayHelloV2(name string) (string, error) {
	//missing name. return error
	if name == "" {
		return "", errors.New("name is mandatory")
	}

	return fmt.Sprintf(randomMessage(), name), nil

}

func SayHelloMulti(names []string) map[string]string {
	output := make(map[string]string, len(names))

	for _, name := range names {
		fmt.Printf("Processing %v\n", name)

		msg, err := SayHelloV2(name)
		if err != nil {
			output[name] = "NOT_FOUND"
		} else {
			output[name] = msg
		}
	}

	return output
}

func randomMessage() string {
	//slice of message formats
	messages := []string{
		"Hello %v",
		"great to see you %v",
		"welcome back %v",
	}

	return messages[rand.Intn(len(messages))]

}
