package grequests

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// RequestOptions is the location that of where the data
type RequestOptions struct {

	// Data is a mpa of key values that will eventually convert into the
	// body if a POST request.
	Data map[string]string

	// Params is a map
	Params map[string]string

	// QueryStruct is a struct that encapsulates a set of
	QueryStruct interface{}

	// Files is where you can include files to upload.
	Files []FileUpload

	JSON interface{}

	XML interface{}

	Headers map[string]string

	DisableCompression bool

	UserAgent string

	Host string

	Auth []string

	IsAjax bool

	Cookies []*http.Cookie

	UseCookieJar bool

	Proxies map[string]*url.URL

	TLSHandshakeTimeout time.Duration

	DialKeepAlive time.Duration

	RequestTimeout time.Duration

	HTTPClient *http.Client

	SensitiveHTTPHeaders map[string]struct{}

	RedirectLimit int

	RequestBody io.Reader

	CookieJar http.CookieJar

	BeforeRequest func(req *http.Request) error

	LocalAddr *net.TCPAddr
}
