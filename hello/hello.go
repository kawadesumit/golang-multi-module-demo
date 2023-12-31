package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	fmt.Println("Starting server")
	greetings.SayHello()

	// msg, err := greetings.SayHelloV2("Aania Kawade")
	msg, err := greetings.SayHelloV2("person1")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	output := greetings.SayHelloMulti([]string{"sumit", "aania", "abhilasha", ""})
	fmt.Println(output)

}
