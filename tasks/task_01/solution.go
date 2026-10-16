package main

import "strings"

func greet(name string) string {
	return formatGreeting(strings.TrimSpace(name))
}

func formatGreeting(name string) string {
	if name == "" {
		name = "World"
	}

	return "Hello, " + name + "!"
}
