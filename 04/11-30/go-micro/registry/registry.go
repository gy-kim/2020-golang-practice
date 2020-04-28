package registry

import "errors"

var (
	DefaultRegistry = NewRegistry()

	// ErrNotFound not found error when GetService is Called.
	ErrNotFound = errors.New("service not found")

	// ErrWatcherStopped watcher stopped error when watcher is stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// Registry provices an interface for service discovery
// and an abstaction over varing implementations
type Registry interface {
	Init(...Option) error
	Options() Options
	Register(*Service, ...RegisterOption) error
	Deregister(*Service, ...DeregisterOption) error
	GetService(string, ...GetOption) ([]*Service, error)
	ListServices(...ListOption) ([]*Service, error)
	Watch(...WatchOption) (Watcher, error)
	String() string
}

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`
	Nodes     []*Node           `json:"nodes"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json::"values"`
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)

type DeregisterOption func(*DeregisterOptions)

type GetOption func(*GetOptions)

type ListOption func(*ListOptions)

// Register a service node. Additionally supply options such as TTL.
func Register(s *Service, opts ...RegisterOption) error {
	return DefaultRegistry.Register(s, opts...)
}

// Deregister a service node
func Deregister(s *Service) error {
	return DefaultRegistry.Deregister(s)
}

// GetService retrive a service. A slice is returned since we separate Namne/Version.
func GetService(name string) ([]*Service, error) {
	return DefaultRegistry.GetService(name)
}

// ListServices list the services. Only returns service names
func ListServices() ([]*Service, error) {
	return DefaultRegistry.ListServices()
}
