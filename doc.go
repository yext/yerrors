// Package yerrors wraps golang.org/x/xerrors to support a no-context Wrapper.
//
// The party line is that error traces should include meaningful context at each
// step. If there's no meaningful context to add, then omitting the stack frame
// is ok. This policy leads to concise stack traces, especially compared to
// enterprise development in Java.
//
// However, the ability to record a stack trace without providing context is
// useful in some places, and the community-favorite `pkg/errors` supported
// that. The official solution (xerrors) does not.
//
// This package provides a no-context Wrapper that is compatible with xerrors.
//
// Summary
//
//   - yerrors.Wrap will add a stack frame to the error without modifying the
//     error's message, printed when formatted with "%+v", but not with "%v".
//
//   - yerrors.Mask does the same thing, while preventing `As` or `Is` from
//     introspecting the error it wraps.
//
//   - The above functionality depends on using yerrors.Errorf instead of
//     xerrors.Errorf, so we also supply aliases to all of the commonly-used
//     things within the xerrors package.
//
// This package should be used to replace all xerrors usage, due to the last
// point above.
//
// Example
//
// Taking this error as an example:
//
//   err := yerrors.Errorf("some context: %w",
//     yerrors.Wrap(
//       yerrors.New("an error")))
//
// It would include all 3 invocations in the stack trace, e.g. when printed like
// this:
//
//   fmt.Printf("%+v", err)
//
// It would include only the two pieces of content in the error message:
//
//   err.Error() == fmt.Sprintf("%v", err)
//
//   // Output: "some context: an error"
//
// See xerrors_test.go for a complete example.
//
package yerrors
