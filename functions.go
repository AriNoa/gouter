package gouter

import "strings"

func DetachCommandFrom(str string) (command string, argument string) {
	slice := strings.Split(str, " ")

	command = slice[0]
	argument = strings.TrimLeft(strings.TrimLeft(str, command), " ")

	return
}
