package errors

import "fmt"

// wrap is a helper to construct an *wapper
func wrap(err error, format, suffix string, args ...interface{}) Err {
	newErr := Err{
		message:  fmt.Sprintf(format+suffix, args...),
		previous: err,
	}
	newErr.SetLocation(2)
	return newErr
}

// timeout represents an error on timeout
type timeout struct {
	Err
}

// Timeoutf returns an error which satisfies IsTimeout().
func Timeoutf(format string, args ...interface{}) error {
	return &timeout{wrap(nil, format, " timeout", args...)}
}

// NewTimeout returns an error which wraps err that satisfies IsTimeout().
func NewTimeout(err error, msg string) error {
	return &timeout{wrap(err, msg, "")}
}

// IsTimeout reports whether err was created with Timeout() or NewTimeout().
func IsTimeout(err error) bool {
	err = Cause(err)
	_, ok := err.(*timeout)
	return ok
}

// notFound represents an error when something has not been found.
type notFound struct {
	Err
}

// NotFoundf returns an error which satisfies IsNotFound().
func NotFoundf(format string, args ...interface{}) error {
	return &notFound{wrap(nil, format, " not found", args...)}
}

// NewNotFound returns an error was created with NotFoundf() or NewNotFound().
func NewNotFound(err error, msg string) error {
	return &notFound{wrap(err, msg, "")}
}

// IsNotFound reports whether err was created with NotFoundf() or NewNotfound().
func IsNotFound(err error) bool {
	err = Cause(err)
	_, ok := err.(*notFound)
	return ok
}

// userNotFound represents an error when an inexistent user is looked up.
type userNotFound struct {
	Err
}

// UserNotFoundf returns an error which satisfies IsUserNotFound().
func UserNotFoundf(format string, args ...interface{}) error {
	return &userNotFound{wrap(nil, format, " user not found", args...)}
}

// NewUserNotFound returns an error which wraps err and satisfies IsUserNotFound().
func NewUserNotFound(err error, msg string) error {
	return &userNotFound{wrap(err, msg, "")}
}

// IsUserNotFound reports whether err was created with userNotFoundf() or NewUserNotFound().
func IsUserNotFound(err error) bool {
	err = Cause(err)
	_, ok := err.(*userNotFound)
	return ok
}

// unauthorized represents an error when an operation is unauthorized.
type unauthorized struct {
	Err
}

// Unauthorizedf returns an error which satisfies IsUnauthorized().
func Unauthorizedf(format string, args ...interface{}) error {
	return &unauthorized{wrap(nil, format, "", args...)}
}

// NewUnauthorized returns an error which wraps err and satisfies IsUnauthorized()
func NewUnauthorized(err error, msg string) error {
	return &unauthorized{wrap(err, msg, "")}
}

// IsUnauthorized reports whether err was created with Unauthorizedf() or NewUnauthorized()
func IsUnauthorized(err error) bool {
	err = Cause(err)
	_, ok := err.(*unauthorized)
	return ok
}

// notImplemented represents an error when something is not implements.
type notImplemented struct {
	Err
}

// NotImplementedf returns an error which wraps err and satisfies IsNotImplemented().
func NotImplementedf(format string, args ...interface{}) error {
	return &notImplemented{wrap(nil, format, " not implemented", args...)}
}

// NewNotImplemented returns an error which wraps err and satisfies isNotImplemented().
func NewNotImplemented(err error, msg string) error {
	return &notImplemented{wrap(err, msg, "")}
}

// IsNotImplemented reports whether err was created with NotImplementedf() or NewNotImplemented().
func IsNotImplemented(err error) bool {
	err = Cause(err)
	_, ok := err.(*notImplemented)
	return ok
}
