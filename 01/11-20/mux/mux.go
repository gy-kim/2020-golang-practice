package mux

import (
	"errors"
	"net/http"
)

var (
	// ErrMethodMismatch is returned when the method in the request does not match
	// the method defined against the route.
	ErrMethodMismatch = errors.New("method is not allowed")

	// ErrNotFound is returned when no route match is found.
	ErrNotFound = errors.New("no matching route was found")
)

// Router registers routes to be matched and dispatches a handler.
type Router struct {
	// Configuratble handler to be used when no route matches.
	NotFoundHandler http.Handler

	//  Configurable Handler to be used when the request method does not match the route.
	MathodNotAllowedHandler http.Handler

	// Routes to be matched, in order.
	routes []*Route

	namedRoutes map[string]*Route

	KeepContext bool

	middlewares []middleware

	routeConf
}

// common route configuration shared between `Router` and `Route`
type routeConf struct {
	useEncodedPath bool

	strictSlash bool

	skipClean bool

	regex routeRegexGroup

	matchers []matcher

	buildScheme string

	buildVarsFunc BuildVarsFunc
}

// RouteMatch stores information about a matched route.
type RouteMatch struct {
	Route   *Route
	Handler http.Handler
	Vars    map[string]string

	MatchErr error
}
