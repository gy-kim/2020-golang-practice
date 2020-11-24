package websocket

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"
)

// ErrBadHandshake is returned when the server response to opening handshake is invalid
var ErrBadHandshake = errors.New("websocket: bad handshake")

var errInvalidCompression = errors.New("websocket: invalid compression negotiation")

// A Dialer contains options for connecting to Websocket server.
type Dialer struct {
	// NetDial specifies the dial function for creating TCP connections. If
	// NetDial is nil, net.Dial is used.
	NetDial func(network, addr string) (net.Conn, error)

	// NetDialContext specifies the dial function for creating TCP connections. If
	// NetDialContext is nil, net.DialContext is used.
	NetDialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	Proxy func(*http.Request) (*url.URL, error)

	TLSClientConfig *tls.Config

	HandshakeTimeout time.Duration

	ReadBifferSize, WriteBufferSize int

	WriteBufferPool BufferPool

	Subprotocols []string

	EnableCompression bool

	Jar http.CookieJar
}
