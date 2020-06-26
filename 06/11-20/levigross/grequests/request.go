package grequests

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"golang.org/x/net/publicsuffix"
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

	InsecureSkipVerify bool

	DisableCompression bool

	UserAgent string

	Host string

	Auth []string

	IsAjax bool

	Cookies []*http.Cookie

	UseCookieJar bool

	Proxies map[string]*url.URL

	TLSHandshakeTimeout time.Duration

	DialTimeout time.Duration

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

// DoRegularRequest adds generic test functionality
// func DoRegularRequest(requestVerb, url string, ro *RequestOptions) (*Response, error) {
// 	return buildResponse(buildRequest(requestVerb, url, ro, nil))
// }

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// func buildRequest(httpMethod, url string, ro *RequestOptions, httpClient *http.Client) (*http.Response, error) {
// 	if ro == nil {
// 		ro = &RequestOptions{}
// 	}

// 	if ro.CookieJar != nil {
// 		ro.UseCookieJar = true
// 	}

// 	if httpClient == nil {
// 		httpClient = BuildHTTPClient(*ro)
// 	}

// 	var err error // we don't want to shadow url so we won't use :=
// 	switch {
// 	case len(ro.Params) != 0:
// 		if url, err = buildURLParams(url, ro.Params); err != nil {
// 			return nil, err
// 		}
// 	case ro.QueryStruct != nil:
// 		if url, err = buildURLStruct(url, ro.QueryStruct); err != nil {
// 			return nil, err
// 		}
// 	}
// }

func encodePostValues(postValues map[string]string) string {
	urlValues := &url.Values{}

	for key, value := range postValues {
		urlValues.Set(key, value)
	}

	return urlValues.Encode()
}

func (ro RequestOptions) proxySettings(req *http.Request) (*url.URL, error) {
	if len(ro.Proxies) == 0 {
		return http.ProxyFromEnvironment(req)
	}

	if _, ok := ro.Proxies[req.URL.Scheme]; ok {
		return ro.Proxies[req.URL.Scheme], nil
	}

	return http.ProxyFromEnvironment(req)
}

func (ro RequestOptions) dontUseDefaultClient() bool {
	switch {
	case ro.InsecureSkipVerify == true:
	case ro.DisableCompression == true:
	case len(ro.Proxies) != 0:
	case ro.TLSHandshakeTimeout != 0:
	case ro.DialTimeout != 0:
	case ro.DialKeepAlive != 0:
	case len(ro.Cookies) != 0:
	case ro.UseCookieJar != false:
	case ro.RequestTimeout != 0:
	case ro.LocalAddr != nil:
	default:
		return false
	}
	return true
}

// BuildHTTPClient is a function what will return a custom HTTP client case on the request options provided
// the check is in UseDefaultClient
func BuildHTTPClient(ro RequestOptions) *http.Client {
	if ro.HTTPClient != nil {
		return ro.HTTPClient
	}

	// Does the user want to change the defaults
	if !ro.dontUseDefaultClient() {
		return http.DefaultClient
	}

	// Using the user config for tls timeout or default
	if ro.TLSHandshakeTimeout == 0 {
		ro.TLSHandshakeTimeout = tlsHandshakeTimeout
	}

	// Using the user config for tls timeout or default
	if ro.DialTimeout == 0 {
		ro.DialTimeout = dialTimeout
	}

	// Using the user config for dial keep alive or default
	if ro.DialKeepAlive == 0 {
		ro.DialKeepAlive = dialKeepAlive
	}

	if ro.RequestTimeout == 0 {
		ro.RequestTimeout = requestTimeout
	}

	var cookieJar http.CookieJar

	if ro.UseCookieJar {
		if ro.CookieJar != nil {
			cookieJar = ro.CookieJar
		} else {
			cookieJar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		}
	}

	return &http.Client{
		Jar:       cookieJar,
		Transport: createHTTPTransport(ro),
		Timeout:   ro.RequestTimeout,
	}
}

func createHTTPTransport(ro RequestOptions) *http.Transport {
	ourHTTPTransport := &http.Transport{
		// These are borrowed from the default transporter
		Proxy: ro.proxySettings,
		Dial: (&net.Dialer{
			Timeout:   ro.DialTimeout,
			KeepAlive: ro.DialKeepAlive,
			LocalAddr: ro.LocalAddr,
		}).Dial,

		// Here comes the user settings
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: ro.InsecureSkipVerify},
		DisableCompression: ro.DisableCompression,
	}
	EnsureTransporterFinalized(ourHTTPTransport)
	return ourHTTPTransport
}

func buildURLParams(userURL string, params map[string]string) (string, error) {
	parseURL, err := url.Parse(userURL)

	if err != nil {
		return "", err
	}

	parseQuery, err := url.ParseQuery(parseURL.RawQuery)

	if err != nil {
		return "", nil
	}

	for key, value := range params {
		parseQuery.Set(key, value)
	}

	return addQueryParams(parseURL, parseQuery), nil
}

func addHTTPHeaders(ro *RequestOptions, req *http.Request) {
	for key, value := range ro.Headers {
		req.Header.Set(key, value)
	}

	if ro.UserAgent != "" {
		req.Header.Set("User-Agent", ro.UserAgent)
	} else {
		req.Header.Set("User-Agent", localUserAgent)
	}

	if ro.Host != "" {
		req.Host = ro.Host
	}

	if ro.Auth != nil {
		req.SetBasicAuth(ro.Auth[0], ro.Auth[1])
	}

	if ro.IsAjax == true {
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
}

func addCookie(ro *RequestOptions, req *http.Request) {
	for _, c := range ro.Cookies {
		req.AddCookie(c)
	}
}

func addQueryParams(parseURL *url.URL, parseQuery url.Values) string {
	return strings.Join([]string{strings.Replace(parseURL.String(), "?"+parseURL.RawQuery, "", -1), parseQuery.Encode()}, "?")
}

func buildURLStruct(userURL string, URLStruct interface{}) (string, error) {
	parseURL, err := url.Parse(userURL)

	if err != nil {
		return "", err
	}

	parseQuery, err := url.ParseQuery(parseURL.RawQuery)

	if err != nil {
		return "", err
	}

	queryStruct, err := query.Values(URLStruct)
	if err != nil {
		return "", err
	}

	for key, value := range queryStruct {
		for _, v := range value {
			parseQuery.Add(key, v)
		}
	}

	return addQueryParams(parseURL, parseQuery), nil
}
