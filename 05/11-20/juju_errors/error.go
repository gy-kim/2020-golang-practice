package errors

import (
	"fmt"
	"reflect"
	"runtime"
)

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

func (e *Err) Error() string {
	// We want to walk up the stack of errors showiwng the annotations
	// as long as the cause is the same.
	err := e.previous
	if !sameError(Cause(err), e.cause) && e.cause != nil {
		err = e.cause
	}
	switch {
	case err == nil:
		return e.message
	case e.message == "":
		return err.Error()
	}
	return fmt.Sprintf("%s: %v", e.message, err)
}

// Format implements fmt.Formatter
func (e *Err) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			fmt.Fprintf(s, "%s", errorStack(e))
			return
		case s.Flag('#'):
			// avoid infinite recursion by wrapping e into a type
			// that doesn't implement Formatter.
			fmt.Fprintf(s, "%#v", (*unformatter)(e))
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s", e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	default:
		fmt.Fprintf(s, "%%!%c(%T=%s)", verb, e, e.Error())
	}
}

type unformatter Err

func (unformatter) Format() {}

// SetLocation records the source location of the error at callDepth stack
// frames above the call.
func (e *Err) SetLocation(callDepth int) {
	_, file, line, _ := runtime.Caller(callDepth + 1)
	e.file = trimSourcePath(file)
	e.line = line
}

// StackTrace returns one string for each location recorded in the stack of errors.
// The first value is the originating eror, with a line for each
// other annotation or tracing of the error.
func (e *Err) StackTrace() []string {
	return errorStack(e)
}

// Ideally we'd have a way to check identity, but deep equals will do.
func sameError(e1, e2 error) bool {
	return reflect.DeepEqual(e1, e2)
}
