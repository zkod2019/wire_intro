package main

// tutorial from: https://www.youtube.com/watch?v=yjCIeT5eWBo&ab_channel=MileHighGophers

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	msg := "Welcome!"
	e, err := initializeEvent(msg)
	if err != nil {
		log.Fatalf("FAILURE: %v", err)
	}
	// m := newMessage("welcome!")
	// g := newGreeter(m)
	// e := newEvent(g)

	e.start()
}

func newMessage(msg string) message {
	return message{value: msg}
}

type message struct {
	value string
}

func (m message) String() string {
	return fmt.Sprintf("[OFFICIAL MESSAGE]: %s", m.value)
}

func newGreeter(m message) greeter {
	return greeter{message: m}
}

type greeter struct {
	message message
}

func (g greeter) greet() string {
	return g.message.String()
}

func newEvent(g greeter) (event, error) {
	// return event{greeter: g}
	return event{}, errors.New("failed event")
}

type event struct {
	greeter greeter
}

func (e event) start() string {
	fmt.Println(e.greeter.greet())
	return e.greeter.greet()
}
