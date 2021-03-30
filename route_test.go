package gouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHandler struct {
	called   bool
	argument string
}

func (h *TestHandler) Handle(arg string) {
	h.called = true
	h.argument = arg
}

func TestAddRouter(t *testing.T) {
	r := New()
	c := New()

	r.AddRouter("test", c)

	child, ok := r.childern["test"]

	assert.True(t, ok, "router does not have a router for the \"test\" command")
	assert.Equal(t, c, child)
}

func TestAddHandler(t *testing.T) {
	r := New()
	h := TestHandler{}

	r.AddHandler("test", &h)

	handler, ok := r.handlers["test"]

	assert.True(t, ok, "router does not have a handler for the \"test\" command")
	assert.Equal(t, &h, handler)
}

func TestDuplicateAddition(t *testing.T) {
	r := New()
	child1 := New()
	child2 := New()
	handler1 := TestHandler{}
	handler2 := TestHandler{}

	var err error

	err = r.AddRouter("child", child1)
	assert.Nil(t, err)

	err = r.AddRouter("child", child2)
	assert.NotNil(t, err)

	err = r.AddHandler("handler", &handler1)
	assert.Nil(t, err)

	err = r.AddHandler("handler", &handler2)
	assert.NotNil(t, err)
}

func TestRoute(t *testing.T) {
	r := New()
	c := New()
	h := TestHandler{}

	c.AddHandler("test", &h)
	r.AddRouter("route", c)

	r.Route("route")
	assert.False(t, h.called)

	r.Route("route test arg")
	assert.True(t, h.called)
	assert.Equal(t, h.argument, "arg")
}
