package grequests

import (
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	localUserAgent = "GRequests/0.10"

	// Default value for net.Dialer Timeout
	dialTimeout = 30 * time.Second

	// Default value for net.Dialer KeepAlive
	dialKeepAlive = 30 * time.Second

	// Default value for http.Transport TLSHandshakeTimeout
	tslHandshakeTimeout = 10 * time.Second

	// Default value for Request Timeout
	requestTimeout = 90 * time.Second
)

var (
	// ErrRedirectLimitExceed is the error returned when the request responded
	// with too many redirects
	ErrRedirectLimitExceed = errors.New("grequest: Request exceeded redirect count")

	// RedirectLimit is a tunable variable that specified how many times we can redirect in request to a redirect.
	RedirectLimit = 30

	//SensitiveHTTPHeaders is a map of sensitive HTTP headers that a user
	// doesn't want passed on a redirrect.
	SensitiveHTTPHeaders = map[string]struct{}{
		"Www-Authenticate":    {},
		"Authorization":       {},
		"Proxy-Authorization": {},
	}
)

// XMLCharDecoder is a helper type that takes a stream of bytes
type XMLCharDecoder func(charset string, input io.Reader) (io.Reader, error)

func addRedirectFunctionality(client *http.Client, ro *RequestOptions) {

}
