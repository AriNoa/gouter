package gouter

import (
	"errors"
)

type Handler interface {
	Handle(arg string)
}

type Router struct {
	childern map[string]*Router
	handlers map[string]Handler
}

func (r *Router) AddRouter(command string, router *Router) error {
	if _, ok := r.childern[command]; ok {
		return errors.New("Router for the command already exists")
	}

	r.childern[command] = router

	return nil
}

func (r *Router) AddHandler(command string, handler Handler) error {
	if _, ok := r.handlers[command]; ok {
		return errors.New("Handler for the command already exists")
	}

	r.handlers[command] = handler

	return nil
}
