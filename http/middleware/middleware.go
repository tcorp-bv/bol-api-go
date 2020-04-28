package middleware

import (
	"fmt"
	"github.com/tcorp-bv/bol-api-go/http/backoff"
	"net/http"
	"strconv"
	"time"
)

// Wraps Transport with retry logic as an http.RoundTripper.
type Middleware struct {
	RetryCodes map[int]bool
	MaxTries   uint
	Backoff    backoff.Backoff
	Transport  http.RoundTripper
	Verbose    bool
	// Testing parameter used to sleep after each try according to the backoff, defaults to time.Sleep
	sleeper func(duration time.Duration)
}

// http.RoundTripper implementation
func (m Middleware) RoundTrip(req *http.Request) (*http.Response, error) {
	// Setup the sleep method for test
	sleep := m.sleeper
	if sleep == nil {
		sleep = time.Sleep
	}
	var res *http.Response
	var err error
	// Retry logic: Retry up to maxTries, wait exponentially long (using the backoff). If status code is 429, wait according to the Retry-After header instead.
	for try := uint(0); try < m.MaxTries; try++ {
		res, err = m.Transport.RoundTrip(req)
		t := m.Backoff.Get(try) // time to wait for next retry
		// Do not try again if there was an error with the transport or the status code is not in the retryCodes
		if err != nil || res == nil || m.RetryCodes[res.StatusCode] == false {
			return res, err
		}
		// In case the api is rate-limited: t is updated to the Retry-After header
		if res.StatusCode == 429 {
			retryAfter := getRetryAfter(res.Header.Get("Retry-After"))
			if retryAfter != 0 {
				t = time.Duration(retryAfter) * time.Second
			}
		}
		// This is for test mainly
		if m.Verbose {
			fmt.Printf("Sleeping %v with code %v\n.", t, res.StatusCode)
		}
		// Wait and try again
		if try != m.MaxTries-1 {
			sleep(t) // TODO: FAIL early if the context deadline is exceeded
		}
	}
	return res, err
}

// returns 0 in case the header is empty or malformed, otherwise converts the string to an integer.
func getRetryAfter(headerValue string) int {
	// Note that due to the way atoi works we could simply return the integer value as it defaults to zero.
	if headerValue == "" {
		return 0
	}
	i, err := strconv.Atoi(headerValue)
	if err != nil {
		return 0
	}
	return i
}
