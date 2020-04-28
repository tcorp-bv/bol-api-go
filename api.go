package bolapi

import (
	"context"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tcorp-bv/bol-api-go/auth"
	"github.com/tcorp-bv/bol-api-go/client"
	"github.com/tcorp-bv/bol-api-go/http/backoff"
	"github.com/tcorp-bv/bol-api-go/http/middleware"
	"net/http"
	"time"
)

const (
	defaultHost  = "api.bol.com"
	startBackoff = 1 * time.Second
	maxBackoff   = 20 * time.Second
	maxTries     = 10
)

var (
	// Constant with error codes for which API requests should be retried.
	//Note that 429 will be retried after the period indicated in the header, the rest will retry according to the backoff.
	retryCodes = map[int]bool{
		429: true, // rate limit reached
		500: true, // internal server error
		504: true, // gateway timeout error
		408: true, // request timeout error
	}
)

// The core interface to retrieve the bol.com v3 client.
// The client contains middleware to retry on rate limits and backoff on server timeouts.
// Requests can thus sometimes take multiple minutes to complete.
type BolApi interface {
	// Get the bol.com http API client.
	GetClient() *client.V3
}

type bolAPI struct {
	authenticator auth.Authenticator
	Transport     *httptransport.Runtime
}

func (api *bolAPI) GetClient() *client.V3 {
	return client.New(api.Transport, strfmt.Default)
}

// Create a new API client with the host (default is "api.bol.com")
func NewWithHost(provider auth.CredentialProvider, host string, basepath string) (BolApi, error) {
	return newBolAPI(provider, host, basepath, nil)
}

func NewWithTransport(provider auth.CredentialProvider, host string, basepath string, transportProvider func(tripper http.RoundTripper) http.RoundTripper) (BolApi, error) {
	// create the original BolAPI
	bolAPI, err := newBolAPI(provider, host, basepath, transportProvider)
	if err != nil {
		return nil, err
	}
	// wrap the transport
	bolAPI.Transport.Transport = transportProvider(bolAPI.Transport.Transport)
	return bolAPI, nil
}

func newBolAPI(provider auth.CredentialProvider, host string, basepath string, transportProvider func(tripper http.RoundTripper) http.RoundTripper) (*bolAPI, error) {
	if host == "" {
		host = defaultHost
	}
	bolAuth, err := auth.New(provider)
	if err != nil {
		return nil, err
	}
	// Gets an http client with the oauth automatically added as middleware
	client := bolAuth.Client(context.Background())
	// Add the retrying middleware
	client.Transport = middleware.Middleware{
		MaxTries:  maxTries,
		Backoff:   backoff.NewExponentialBackoff(startBackoff, maxBackoff),
		Transport: client.Transport,
	}
	// Add the client middleware
	if transportProvider != nil {
		client.Transport = transportProvider(client.Transport)
	}
	// Sets up the runtime
	transport := httptransport.NewWithClient(host, basepath, nil, client)
	// Map the consumers and producers to json
	transport.Consumers["application/vnd.retailer.v3+json"] = runtime.JSONConsumer()
	transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
	transport.Producers["application/vnd.retailer.v3+json"] = runtime.JSONProducer()
	return &bolAPI{Transport: transport}, nil
}

func New(provider auth.CredentialProvider) (BolApi, error) {
	return NewWithHost(provider, defaultHost, "")
}
