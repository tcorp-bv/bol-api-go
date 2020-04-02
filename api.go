package bolapi

import (
	"context"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tcorp-bv/bol-api-go/auth"
	"github.com/tcorp-bv/bol-api-go/client"
)

const (
	defaultHost = "api.bol.com"
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

	return &bolApi{transport: transport}, nil
}

func New(provider auth.CredentialProvider) (BolApi, error) {
	return NewWithHost(provider, defaultHost)
}
