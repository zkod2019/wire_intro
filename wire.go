//go:build wireinject
// +build wireinject

// ^ add that as the reverse of what is generated in wire_gen.go
// basically means that you want to build this file when you're asking for an inject
package main

import "github.com/google/wire"

func initializeEvent(msg string) (event, error) {
	wire.Build(newMessage, newGreeter, newEvent)
	return event{}, nil
}
