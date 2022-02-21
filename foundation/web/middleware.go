package web

// Middleware is wrapping Handlers with additional business logic
type Middleware func(h Handler) Handler

// WrapMiddleware will take slice of all Middlewares wrap them around and return modified handler
func WrapMiddleware(mws []Middleware, h Handler) Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h

}
