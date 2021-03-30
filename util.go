package gouter

import "strings"

// DetachCommandFrom function detaches the command from the string and returns it and its arguments
func DetachCommandFrom(str string) (command string, argument string) {
	slice := strings.Split(str, " ")

	command = slice[0]
	argument = strings.TrimLeft(strings.TrimLeft(str, command), " ")

	return
}
