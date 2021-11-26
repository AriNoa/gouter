package gouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetachCommandFromEmpty(t *testing.T) {
	cmd, arg := DetachCommandFrom("")

	assert.Equal(t, cmd, "")
	assert.Equal(t, arg, "")
}

func TestDetachCommandFromCommandOnly(t *testing.T) {
	cmd, arg := DetachCommandFrom("command")

	assert.Equal(t, cmd, "command")
	assert.Equal(t, arg, "")
}

func TestDetachCommandFromCommands(t *testing.T) {
	cmd, arg := DetachCommandFrom("command arg")

	assert.Equal(t, cmd, "command")
	assert.Equal(t, arg, "arg")
}
