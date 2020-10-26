package yerrors

import (
	"golang.org/x/xerrors"
)

// Create aliases for xerrors functionality, with the idea that typical
// application code uses only this package.

type Wrapper = xerrors.Wrapper

var (
	New    = xerrors.New
	Is     = xerrors.Is
	As     = xerrors.As
	Unwrap = xerrors.Unwrap
	Opaque = xerrors.Opaque
)

// Wrap returns the given error transparently wrapped, without additional context.
// As a special case, a nil argument results in a nil return.
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return &wrapError{"", err, xerrors.Caller(1)}
}

// WrapFrame is like Wrap, except uses the location a given number of stack frames up.
// WrapFrame(err, 0) is equivalent to Wrap(err).
func WrapFrame(err error, caller int) error {
	if err == nil {
		return nil
	}
	return &wrapError{"", err, xerrors.Caller(1 + caller)}
}

// Mask returns the given error opaquely wrapped, without additional context.
// As a special case, a nil argument results in a nil return.
func Mask(err error) error {
	if err == nil {
		return nil
	}
	return &noWrapError{"", err, xerrors.Caller(1)}
}
