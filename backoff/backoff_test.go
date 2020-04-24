package backoff

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

func TestExponentialBackoff(t *testing.T) {
	is := is.New(t)
	s := 2 * time.Second
	max := 20 * time.Second
	bo := NewExponentialBackoff(s, max)
	is.Equal(bo.Get(0), s)
	is.Equal(bo.Get(1), s*2)
	is.Equal(bo.Get(2), s*4)
	is.Equal(bo.Get(3), s*8)
	is.Equal(bo.Get(4), max)
	is.Equal(bo.Get(5), max)
	is.Equal(bo.Get(10000), max)
}
