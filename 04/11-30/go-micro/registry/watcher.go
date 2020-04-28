package registry

import "time"

// Watcher is an interface that returns updates about services within the registry.
type Watcher interface {
	Next() (*Result, error)
	Stop()
}

// Result is returned by a call to Next on watcher.
// Actions can be crate, update, delete
type Result struct {
	Action  string
	Service *Service
}

// EventType defines registry event type
type EventType int

const (
	// Create is emiitted when a new service is registered
	Create EventType = iota
	// Delete is emitted when an existing service is deregisterd
	Delete
	// Update is emitted when an exiting services is updated.
	Update
)

func (t EventType) String() string {
	switch t {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

// Event is registry event
type Event struct {
	Id        string
	Type      EventType
	Timestamp time.Time
	Service   *Service
}
