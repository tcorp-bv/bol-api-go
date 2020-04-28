package middleware

import (
	"fmt"
	"github.com/matryer/is"
	"github.com/tcorp-bv/bol-api-go/http/backoff"
	"net/http"
	"testing"
	"time"
)

const (
	maxTries = 4
)

var (
	req                         = http.Request{}
	successRes                  = http.Response{StatusCode: 200}
	retryResWith2sHeader        = http.Response{StatusCode: 429, Header: http.Header{"Retry-After": []string{"2"}}}
	retryResWithMalformedHeader = http.Response{StatusCode: 429, Header: http.Header{"Retry-After": []string{"2s"}}}
	retryResWithoutHeader       = http.Response{StatusCode: 429, Header: http.Header{}}
	retryResErrWithRetry        = http.Response{StatusCode: 500, Header: http.Header{}}
	retryResErrWithoutRetry     = http.Response{StatusCode: 404, Header: http.Header{}}
)

func TestMiddleware_success(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0)
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			return &successRes, nil
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, 1)
	is.Equal(slept, int64(0))
	is.Equal(err, nil)
	is.Equal(*res, successRes)
}
func TestMiddleware_withRetry(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0) //ms
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			return &retryResErrWithRetry, nil
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, maxTries)
	is.Equal(slept, int64(600))
	is.Equal(err, nil)
	is.Equal(*res, retryResErrWithRetry)
}
func TestMiddleware_withRetryOnce(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0) //ms
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			if calls == 2 {
				return &successRes, nil
			}
			return &retryResErrWithRetry, nil
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, 2)
	is.Equal(slept, int64(100))
	is.Equal(err, nil)
	is.Equal(*res, successRes)
}
func TestMiddleware_withRetryRatelimit(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0) //ms
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			if calls == 3 {
				return &successRes, nil
			}
			return &retryResWith2sHeader, nil
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, 3)
	is.Equal(slept, int64(4000))
	is.Equal(err, nil)
	is.Equal(*res, successRes)
}
func TestMiddleware_withRetryMalformedRatelimit(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0) //ms
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			if calls == 3 {
				return &successRes, nil
			}
			return &retryResWithMalformedHeader, nil
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, 3)
	is.Equal(slept, int64(300))
	is.Equal(err, nil)
	is.Equal(*res, successRes)
}

func TestMiddleware_WithoutRateLimitHeader(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0) //ms
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			if calls == 3 {
				return &successRes, nil
			}
			return &retryResWithoutHeader, nil
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, 3)
	is.Equal(slept, int64(300))
	is.Equal(err, nil)
	is.Equal(*res, successRes)
}

func TestMiddleware_WithTransportError(t *testing.T) {
	is := is.New(t)
	calls := 0
	slept := int64(0) //ms
	c := config{
		transport: func(r *http.Request) (*http.Response, error) {
			calls++
			return nil, fmt.Errorf("mock transport error")
		},
		sleeper: func(duration time.Duration) {
			slept += duration.Milliseconds()
		},
	}
	res, err := c.getTestMiddleware().RoundTrip(&req)

	is.Equal(calls, 1)
	is.Equal(slept, int64(0))
	is.True(err != nil)
	is.Equal(res, nil)
}

type transportWrapper struct {
	transport func(r *http.Request) (*http.Response, error)
}

func (t transportWrapper) RoundTrip(request *http.Request) (*http.Response, error) {
	return t.transport(request)
}

type config struct {
	transport func(r *http.Request) (*http.Response, error)
	sleeper   func(duration time.Duration)
}

func (c *config) getTestMiddleware() Middleware {
	// Should backoff as 100ms, 200ms, 300ms, 300ms...
	testBackoff := backoff.NewExponentialBackoff(100*time.Millisecond, 300*time.Millisecond)
	// Retry on 429 (according to the Retry-After header)
	m := Middleware{
		RetryCodes: map[int]bool{429: true, 500: true},
		MaxTries:   maxTries,
		Backoff:    testBackoff,
		Transport:  transportWrapper{transport: c.transport},
		sleeper:    c.sleeper,
	}
	return m
}
