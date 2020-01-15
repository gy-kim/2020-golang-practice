package mux

import "net/http"

// Route stores information to match a request and build URLs.
type Route struct {
	handler     http.Handler
	buldOnly    bool
	name        string
	err         error
	namedRoutes map[string]*Route

	routeConf
}

// SkipClean reports whether path cleaning is enabled for this route via
// Router.SkipClean.
func (r *Route) SkipClean() bool {
	return r.skipClean
}

type matcher interface {
	Match(*http.Request, *RouteMatch) bool
}

// BuildVarsFunc is the function signature used by custom build variable
// functions (which can modify route variables before a route's URL is built)
type BuildVarsFunc func(map[string]string) map[string]string
