package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
	names := []string{"Gladys", "Samantha", "Darrin"}

	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(messages)
}