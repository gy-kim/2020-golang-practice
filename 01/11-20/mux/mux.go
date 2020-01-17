package mux

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"regexp"
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

	regexp routeRegexpGroup

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

func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash expept for root;
	// put the trailing slash back of necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}
	return np
}

func uniqueVars(s1, s2 []string) error {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				return fmt.Errorf("mux: duplicated route variable %q", v2)
			}
		}
	}
	return nil
}

func checkPairs(pairs ...string) (int, error) {
	length := len(pairs)
	if length%2 != 0 {
		return length, fmt.Errorf("mux: number of parameters must be multiple of 2, got %v", pairs)
	}
	return length, nil
}

func mapFromPairsToString(pairs ...string) (map[string]string, error) {
	length, err := checkPairs(pairs...)
	if err != nil {
		return nil, err
	}
	m := make(map[string]string, length/2)
	for i := 0; i < length; i += 2 {
		m[pairs[i]] = pairs[i+1]
	}
	return m, nil
}

func matchInArray(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func matchMapWithString(toCheck map[string]string, toMatch map[string][]string, canonicalKey bool) bool {
	for k, v := range toCheck {
		if canonicalKey {
			k = http.CanonicalHeaderKey(k)
		}
		if values := toMatch[k]; values == nil {
			return false
		} else if v != "" {
			valueExists := false
			for _, value := range values {
				if v == value {
					valueExists = true
					break
				}
			}
			if !valueExists {
				return false
			}
		}
	}
	return true
}

func matchMapWithRegex(toCheck map[string]*regexp.Regexp, toMatch map[string][]string, canonicalKey bool) bool {
	for k, v := range toCheck {
		if canonicalKey {
			k = http.CanonicalHeaderKey(k)
		}
		if values := toMatch[k]; values == nil {
			return false
		} else if v != nil {
			valueExists := false
			for _, value := range values {
				if v.MatchString(value) {
					valueExists = true
					break
				}
			}
			if !valueExists {
				return false
			}
		}
	}
	return true
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func methodNotAllowedHandler() http.Handler { return http.HandlerFunc(methodNotAllowed) }
