package errors

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFormatLocation(t *testing.T) {
	err := NewLoc("inner")

	assert.Equal(t, "inner", fmt.Sprintf("%v", err))
	assert.Equal(t, "inner (fmt_test.go:12)", fmt.Sprintf("%+v", err))
	assert.Equal(t, "inner at github.com/nikandfor/errors/fmt_test.go:12", fmt.Sprintf("% +v", err))

	// more
	err = WrapLoc(
		WrapLoc(os.ErrNotExist, "middle"),
		"global")

	assert.Equal(t, "global: middle: file does not exist", fmt.Sprintf("%v", err))
	assert.Equal(t, "global (fmt_test.go:19): middle (fmt_test.go:20): file does not exist", fmt.Sprintf("%+v", err))
	assert.Equal(t, `global at github.com/nikandfor/errors/fmt_test.go:19
middle at github.com/nikandfor/errors/fmt_test.go:20
file does not exist`, fmt.Sprintf("% +v", err))

	// one more
	err = WrapLoc(
		WrapLoc(
			NewLoc("inner"),
			"middle"),
		"global")

	assert.Equal(t, "global: middle: inner", fmt.Sprintf("%v", err))
	assert.Equal(t, "global (fmt_test.go:30): middle (fmt_test.go:31): inner (fmt_test.go:32)", fmt.Sprintf("%+v", err))
	assert.Equal(t, `global at github.com/nikandfor/errors/fmt_test.go:30
middle at github.com/nikandfor/errors/fmt_test.go:31
inner at github.com/nikandfor/errors/fmt_test.go:32`, fmt.Sprintf("% +v", err))

	// with no messages
	err = WrapLoc(
		WrapLoc(
			NewLoc(""),
			""),
		"")

	assert.Equal(t, "(no message)", fmt.Sprintf("%v", err))
	assert.Equal(t, "(no message) (fmt_test.go:43): (no message) (fmt_test.go:44): (no message) (fmt_test.go:45)", fmt.Sprintf("%+v", err))
	assert.Equal(t, `(no message) at github.com/nikandfor/errors/fmt_test.go:43
(no message) at github.com/nikandfor/errors/fmt_test.go:44
(no message) at github.com/nikandfor/errors/fmt_test.go:45`, fmt.Sprintf("% +v", err))
}

func TestErrorFormat(t *testing.T) {
	err := New("inner")

	assert.Equal(t, "inner", fmt.Sprintf("%v", err))
	assert.Equal(t, "inner", fmt.Sprintf("%+v", err))
	assert.Equal(t, "inner", fmt.Sprintf("% +v", err))

	// more
	err = Wrap(Wrap(os.ErrNotExist, "middle"), "global")

	assert.Equal(t, "global: middle: file does not exist", fmt.Sprintf("%v", err))
	assert.Equal(t, "global: middle: file does not exist", fmt.Sprintf("%+v", err))
	assert.Equal(t, "global: middle: file does not exist", fmt.Sprintf("% +v", err))

	// one more
	err = Wrap(Wrap(New("inner"), "middle"), "global")

	assert.Equal(t, "global: middle: inner", fmt.Sprintf("%v", err))
	assert.Equal(t, "global: middle: inner", fmt.Sprintf("%+v", err))
	assert.Equal(t, "global: middle: inner", fmt.Sprintf("% +v", err))
}
