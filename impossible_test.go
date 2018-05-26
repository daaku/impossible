package impossible

import (
	"errors"
	"testing"

	"github.com/daaku/ensure"
)

func TestNilError(t *testing.T) {
	Error(nil)
}

func TestError(t *testing.T) {
	givenErr := errors.New("foo")
	defer ensure.PanicDeepEqual(t, givenErr)
	Error(givenErr)
}
