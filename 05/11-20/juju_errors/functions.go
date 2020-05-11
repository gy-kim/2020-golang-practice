package errors

// Cause returns the cause of the given error.
// This will be either the original error, or the result of a Wrap a Wrap or Mask call.
func Cause(err error) error {
	var diag error
	if err, ok := err.(causer); ok {
		diag = err.Cause()
	}
	if diag != nil {
		return diag
	}
	return err
}

type causer interface {
	Cause() error
}

type wrapper interface {
	// Message returns the top level error message,
	// not including the message from the Previous error.
	Message() string

	// Underlying returns the Previous error, or nil
	// if there
	Underlying() error
}

type locationer interface {
	Location() (string, int)
}

var (
	_ wrapper    = (*Err)(nil)
	_ locationer = (*Err)(nil)
	_ causer     = (*Err)(nil)
)
