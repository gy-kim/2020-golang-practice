package errors

import "fmt"

// Err holds a description of an error along with informatoin about
// where the error was created
//
// It may be enbedded in custom error types t add extra imformation that
// this errors package can understand.
type Err struct {
	// message hold an annotation of the error
	message string

	// cause hold the cause if the error as returned
	// by the Cause method.
	cause error

	// previous holds the previous error in the error stack, if any.
	previous error

	// // file and line hold the source code location where the error was created.
	file string
	line int
}

// NewErr is used to return an Err for the purpose of embedding in other structures.
func NewErr(format string, args ...interface{}) Err {
	return Err{
		message: fmt.Sprintf(format, args...),
	}
}

// NewErrWithCause is used to return an Err with cause by other error for the purpose of embedding in other structures.
func NewErrWithCause(other error, format string, args ...interface{}) Err {
	return Err{
		message:  fmt.Sprintf(format, args...),
		cause:    Cause(other),
		previous: other,
	}
}

// Location is the file and line of where the error was most recently
// created or annotated.
func (e *Err) Location() (filename string, line int) {
	return e.file, e.line
}

// Underlying returns the previous error in the error stack, if any.
func (e *Err) Underlying() error {
	return e.previous
}

// Cause returns the most recent error in the error stack that
// meets one of these criteria: the original error that was raised; the new
// error that was passed into the Wrap function;
func (e *Err) Cause() error {
	return e.cause
}

// Message returns the message stored with the most recent location.
func (e *Err) Message() string {
	return e.message
}
