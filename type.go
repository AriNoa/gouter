package gouter

import (
	"errors"
)

// Handler is interface for command execution
type Handler interface {
	Handle(arg string)
}

// Router is structure that router for commands
type Router struct {
	childern map[string]*Router
	handlers map[string]Handler
}

// New returns new router
func New() *Router {
	return &Router{
		map[string]*Router{},
		map[string]Handler{},
	}
}

// AddRouter adds the child router
func (r *Router) AddRouter(command string, router *Router) error {
	if _, ok := r.childern[command]; ok {
		return errors.New("Router for the command already exists")
	}

	r.childern[command] = router

	return nil
}

// AddHandler adds the handler
func (r *Router) AddHandler(command string, handler Handler) error {
	if _, ok := r.handlers[command]; ok {
		return errors.New("Handler for the command already exists")
	}

	r.handlers[command] = handler

	return nil
}

// Route handles the command
func (r *Router) Route(str string) {
	cmd, arg := DetachCommandFrom(str)

	if handler, ok := r.handlers[cmd]; ok {
		handler.Handle(arg)
	}

	if child, ok := r.childern[cmd]; ok {
		child.Route(arg)
	}
}
