package gouter

type HandlerFunc func(string)

type MiddlewareFunc func(HandlerFunc) HandlerFunc

type Router struct {
	handlers map[string]HandlerFunc
}

func New(name string) *Router {
	return &Router{
		handlers: map[string]HandlerFunc{},
	}
}

func (r *Router) UseChild(cmd string, router *Router) *Router {
	r.handlers[cmd] = router.Handle

	return r
}

func (r *Router) UseHandler(cmd string, handler HandlerFunc) *Router {
	r.handlers[cmd] = handler

	return r
}

func (r *Router) UseMiddleware(middleware MiddlewareFunc) *Router {
	for cmd, handler := range r.handlers {
		r.handlers[cmd] = middleware(handler)
	}

	return r
}

func (r *Router) Handle(arg string) {
	cmd, arg := DetachCommandFrom(arg)

	handler, ok := r.handlers[cmd]
	if !ok {
		return
	}

	handler(arg)
}
