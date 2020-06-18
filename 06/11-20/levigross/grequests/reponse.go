package grequests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// Response is what is returned to a user when they fire off a request
type Response struct {

	// Ok is a boolean flag that validates that the server returned a 2xx code
	Ok bool

	// This is the Go error flag - if something went wrong within the request, this flag will be set.
	Error error

	// We want to abstract (at leat at the moment) the Go http.Response object away. So we are going to make use of it
	// internal but not give the user access
	RawResponse *http.Response

	// StatusCode is the HTTP status code returned by the HTTP Response. Taken from resp.StatusCode
	StatusCode int

	// Header is a net/http/Header structure
	Header http.Header

	internalByteBuffer *bytes.Buffer
}

func buildResponse(resp *http.Response, err error) (*Response, error) {
	if err != nil {
		return &Response{Error: err}, err
	}

	goodResp := &Response{
		// If your coee is within 2xx range - the response is considered `Ok`
		Ok:                 resp.StatusCode >= 200 && resp.StatusCode < 300,
		Error:              nil,
		RawResponse:        resp,
		StatusCode:         resp.StatusCode,
		Header:             resp.Header,
		internalByteBuffer: bytes.NewBuffer([]byte{}),
	}

	return goodResp, nil
}

// Read is part of our ability to support io.ReadCloser if someone wants to make use of the raw body
func (r *Response) Read(p []byte) (n int, err error) {
	if r.Error != nil {
		return -1, r.Error
	}

	return r.RawResponse.Body.Read(p)
}

// Close is part of our ability to support io.ReadCloser if someone want to make use of the raw body
func (r *Response) Close() error {
	if r.Error != nil {
		return r.Error
	}

	io.Copy(ioutil.Discard, r)

	return r.RawResponse.Body.Close()
}
