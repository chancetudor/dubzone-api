package middleware

import "github.com/sirupsen/logrus"

type Middleware struct {
	log *logrus.Logger
}

func NewMiddleware(l *logrus.Logger) *Middleware {
	return &Middleware{log: l}
}

// LoadoutKey is the key to be used for use in validating request bodies.
// When we store the internal representation of a JSON struct,
// we use the context.WithValue function, which takes this as its key.
type LoadoutKey struct{}

// CatKey is the key to be used for use in validating category parameters.
// When a category parameter is correct, we add the category to the request context,
// for use in the next handler.
type CatKey struct{}

// NameKey is the key to be used for use in validating name parameters.
// When a name parameter is correct, we add the name to the request context,
// for use in the next handler.
type NameKey struct{}
