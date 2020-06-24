package grequests

import (
	"errors"
	"io"
	"net/http"
	"runtime"
	"time"
)

const (
	localUserAgent = "GRequests/0.10"

	// Default value for net.Dialer Timeout
	dialTimeout = 30 * time.Second

	// Default value for net.Dialer KeepAlive
	dialKeepAlive = 30 * time.Second

	// Default value for http.Transport TLSHandshakeTimeout
	tlsHandshakeTimeout = 10 * time.Second

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
	if client.CheckRedirect != nil {
		return
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if ro.RedirectLimit < 0 {
			return http.ErrUseLastResponse
		}

		if ro.RedirectLimit == 0 {
			ro.RedirectLimit = RedirectLimit
		}

		if len(via) >= ro.RedirectLimit {
			return ErrRedirectLimitExceed
		}

		if ro.SensitiveHTTPHeaders == nil {
			ro.SensitiveHTTPHeaders = SensitiveHTTPHeaders
		}

		for k, vv := range via[0].Header {
			if _, found := ro.SensitiveHTTPHeaders[k]; found {
				continue
			}
			for _, v := range vv {
				req.Header.Add(k, v)
			}
		}

		return nil
	}
}

// EnsureTransporterFinalized will ensure that when the HTTP clinet is GCed
// the runtime will close the idel connections
func EnsureTransporterFinalized(httpTransport *http.Transport) {
	runtime.SetFinalizer(&httpTransport, func(transportInt **http.Transport) {
		(*transportInt).CloseIdleConnections()
	})
}
