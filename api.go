package bolapi

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tcorp-bv/bol-api-go/auth"
	"github.com/tcorp-bv/bol-api-go/backoff"
	"github.com/tcorp-bv/bol-api-go/client"
	"net/http"
	"time"
)

const (
	defaultHost  = "api.bol.com"
	startBackoff = 1 * time.Second
	maxBackoff   = 20 * time.Second
	maxTries     = 50
)

var (
	doNotRetryCodes = map[int]bool{200: true, 400: true}
)

// The core interface to retrieve the bol.com v3 client
type BolApi interface {
	GetClient() *client.V3
}

type bolAPI struct {
	authenticator auth.Authenticator
	Transport     *httptransport.Runtime
}

func (api *bolAPI) GetClient() *client.V3 {
	return client.New(api.Transport, strfmt.Default)
}

type middleware struct {
	maxTries  uint
	backoff   backoff.Backoff
	transport http.RoundTripper
	verbose   bool
}

// Middleware to add retrying an
func (m middleware) RoundTrip(req *http.Request) (*http.Response, error) {
	var tries uint = 0
	var res *http.Response
	for tries <= m.maxTries {
		res, err := m.transport.RoundTrip(req)
		if err != nil || res == nil || doNotRetryCodes[res.StatusCode] == true {
			return res, err
		}
		t := m.backoff.Get(tries)
		if m.verbose {
			fmt.Printf("Sleeping %v with code %v\n.", t, res.StatusCode)
		}
		time.Sleep(t)
		tries++
	}
	return res, nil
}

// Create a new API client with the host (default is "api.bol.com")
func NewWithHost(provider auth.CredentialProvider, host string) (BolApi, error) {
	return newBolAPI(provider, host)
}

func NewWithTransport(provider auth.CredentialProvider, host string, transportProvider func(tripper http.RoundTripper) http.RoundTripper) (BolApi, error) {
	// create the original BolAPI
	bolAPI, err := newBolAPI(provider, host)
	if err != nil {
		return nil, err
	}
	// wrap the transport
	bolAPI.Transport.Transport = transportProvider(bolAPI.Transport.Transport)
	return bolAPI, nil
}

func newBolAPI(provider auth.CredentialProvider, host string) (*bolAPI, error) {
	if host == "" {
		host = defaultHost
	}
	bolAuth, err := auth.New(provider)
	if err != nil {
		return nil, err
	}
	// Set up he transport to use the client with oauth
	client := bolAuth.Client(context.Background())
	client.Transport = middleware{
		maxTries:  maxTries,
		backoff:   backoff.NewExponentialBackoff(startBackoff, maxBackoff),
		transport: client.Transport,
	}
	transport := httptransport.NewWithClient(host, "", nil, client)
	// Map the consumers and producers to json
	transport.Consumers["application/vnd.retailer.v3+json"] = runtime.JSONConsumer()
	transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
	transport.Producers["application/vnd.retailer.v3+json"] = runtime.JSONProducer()
	return &bolAPI{Transport: transport}, nil
}

func New(provider auth.CredentialProvider) (BolApi, error) {
	return NewWithHost(provider, defaultHost)
}
