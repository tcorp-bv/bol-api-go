package bolapi

import (
	"context"
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

// The core interface to retrieve the bol.com v3 client
type BolApi interface {
	GetClient() *client.V3
}

type bolApi struct {
	authenticator auth.Authenticator
	transport     *httptransport.Runtime
}

func (api *bolApi) GetClient() *client.V3 {
	return client.New(api.transport, strfmt.Default)
}

type middleware struct {
	maxTries  uint
	backoff   backoff.Backoff
	transport http.RoundTripper
}

// Middleware to add retrying an
func (m middleware) RoundTrip(req *http.Request) (*http.Response, error) {
	var tries uint = 0
	var err error
	for tries <= m.maxTries {
		res, err := m.transport.RoundTrip(req)
		if err == nil {
			return res, nil
		}
		time.Sleep(m.backoff.Get(tries))
		tries++
	}
	return nil, err
}

// Create a new API client with the host (default is "api.bol.com")
func NewWithHost(provider auth.CredentialProvider, host string) (BolApi, error) {
	if host == "" {
		host = defaultHost
	}

	bolAuth, err := auth.New(provider)
	if err != nil {
		return nil, err
	}
	// Set up he transport to use the client with oauth
	transport := httptransport.NewWithClient(host, "", nil, bolAuth.Client(context.Background()))
	// Map the consumers and producers to json
	transport.Consumers["application/vnd.retailer.v3+json"] = runtime.JSONConsumer()
	transport.Consumers["application/problem+json"] = runtime.JSONConsumer()
	transport.Producers["application/vnd.retailer.v3+json"] = runtime.JSONProducer()
	transport.Transport = middleware{
		maxTries:  maxTries,
		backoff:   backoff.NewExponentialBackoff(startBackoff, maxBackoff),
		transport: transport.Transport,
	}
	return &bolApi{transport: transport}, nil
}

func New(provider auth.CredentialProvider) (BolApi, error) {
	return NewWithHost(provider, defaultHost)
}
