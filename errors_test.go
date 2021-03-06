package errors

import (
	"os"
	"testing"

	"github.com/nikandfor/loc"
	"github.com/stretchr/testify/assert"
)

type testWrapper struct {
	err error
}

func TestWrapperError(t *testing.T) {
	assert.EqualError(t, New("qwe"), "qwe")
	assert.EqualError(t, NewNoLoc("qwe %v %v", 1, 2), "qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, WrapNoLoc(New("qwe %v %v", 1, 2), "context %v %v", "a", "b"), "context a b: qwe 1 2")
	assert.EqualError(t, Wrap(New("qwe %v %v", 1, 2), "context").(interface{ Unwrap() error }).Unwrap(), "qwe 1 2")
	assert.EqualError(t, Unwrap(Wrap(New("qwe %v %v", 1, 2), "context")), "qwe 1 2")
}

//nolint:goerr113
func TestUnwrap(t *testing.T) {
	mid := Wrap(os.ErrNotExist, "middle")
	err := Wrap(mid, "global")

	assert.True(t, os.ErrNotExist == Unwrap(mid))

	assert.True(t, mid == Unwrap(err))

	assert.True(t, mid == Unwrap(testWrapper{mid}))

	assert.Nil(t, Unwrap(os.ErrNotExist))
}

func TestWrapNil(t *testing.T) {
	var err error

	assert.Nil(t, Wrap(err, "qwe"))
	assert.Nil(t, WrapNoLoc(err, "qwe"))
	assert.Nil(t, WrapDepth(err, 0, "qwe"))
	assert.Nil(t, WrapLoc(err, loc.Funcentry(0), "qwe"))
}

func (w testWrapper) Error() string { return "none" }

func (w testWrapper) Unwrap() error { return w.err }
