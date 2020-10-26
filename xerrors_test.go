package yerrors

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrap(t *testing.T) {
	err := errors.New("root err")
	werr := Errorf("context provided: %w", err)
	merr := Wrap(werr)
	m2err := Wrap(merr)
	terr := Errorf("top context: %w", m2err)

	require.True(t, Is(werr, err))
	require.True(t, Is(merr, err))
	require.True(t, Is(m2err, err))
	require.True(t, Is(terr, err))
	require.Equal(t, strings.TrimSpace(`
top context:
    github.com/yext/yerrors.TestWrap
        xerrors_test.go:18
  - github.com/yext/yerrors.TestWrap
        xerrors_test.go:17
  - github.com/yext/yerrors.TestWrap
        xerrors_test.go:16
  - context provided:
    github.com/yext/yerrors.TestWrap
        xerrors_test.go:15
  - root err
`), regexp.MustCompile(`/.*/yerrors/`).
		ReplaceAllString(fmt.Sprintf("%+v", terr), ""))
}

func TestMask(t *testing.T) {
	err := errors.New("root err")
	werr := Errorf("context provided: %w", err)
	merr := Mask(werr)
	m2err := Mask(merr)
	terr := Errorf("top context: %w", m2err)

	require.True(t, Is(werr, err))
	require.False(t, Is(merr, err))
	require.False(t, Is(m2err, err))
	require.False(t, Is(terr, err))
	require.Equal(t, strings.TrimSpace(`
top context:
    github.com/yext/yerrors.TestMask
        xerrors_test.go:45
  - github.com/yext/yerrors.TestMask
        xerrors_test.go:44
  - github.com/yext/yerrors.TestMask
        xerrors_test.go:43
  - context provided:
    github.com/yext/yerrors.TestMask
        xerrors_test.go:42
  - root err
`), regexp.MustCompile(`/.*/yerrors/`).
		ReplaceAllString(fmt.Sprintf("%+v", terr), ""))
}

func TestWrap_NoStack(t *testing.T) {
	err := errors.New("root err")
	wrapped := Wrap(err)
	require.Equal(t, "root err", fmt.Sprintf("%v", wrapped))
	require.Equal(t, "root err", wrapped.Error())

	wrapped2 := Wrap(wrapped)
	require.Equal(t, "root err", fmt.Sprintf("%v", wrapped2))
	require.Equal(t, "root err", wrapped2.Error())

	wrapped3 := Errorf("context: %w", wrapped2)
	require.Equal(t, "context: root err", fmt.Sprintf("%v", wrapped3))
	require.Equal(t, "context: root err", wrapped3.Error())

	wrapped4 := Wrap(wrapped3)
	require.Equal(t, "context: root err", fmt.Sprintf("%v", wrapped4))
	require.Equal(t, "context: root err", wrapped4.Error())
}

func TestMask_NoStack(t *testing.T) {
	err := errors.New("root err")
	masked := Mask(err)
	require.Equal(t, "root err", fmt.Sprintf("%v", masked))
	require.Equal(t, "root err", masked.Error())

	masked2 := Mask(masked)
	require.Equal(t, "root err", fmt.Sprintf("%v", masked2))
	require.Equal(t, "root err", masked2.Error())

	masked3 := Errorf("context: %w", masked2)
	require.Equal(t, "context: root err", fmt.Sprintf("%v", masked3))
	require.Equal(t, "context: root err", masked3.Error())

	masked4 := Mask(masked3)
	require.Equal(t, "context: root err", fmt.Sprintf("%v", masked4))
	require.Equal(t, "context: root err", masked4.Error())
}
