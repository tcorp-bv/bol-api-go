package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

const (
	defaultAuthUrl = "https://login.bol.com"
	tokenEndpoint  = "token"
)

// Gets the default token endpoint (https://login.bol.com/token)
func defaultTokenEndpoint() string {
	return fmt.Sprintf("%s/%s", defaultAuthUrl, tokenEndpoint)
}

// The authenticator manages the oauth authentication to the bol API
type Authenticator interface {
	// Returns the bearer token or an error
	Token() (*oauth2.Token, error)
	Client(ctx context.Context) *http.Client
}

// Creates a new instance of the Authenticator using the credentials from the provider.
func New(provider CredentialProvider) (Authenticator, error) {
	config, err := getConfig(provider)
	if err != nil {
		return nil, err
	}
	return &authenticator{config: config}, nil
}

// Authenticator implementation using golang's oauth2 client.
type authenticator struct {
	config clientcredentials.Config
}

// Fetch Token from the golang oauth2 client
func (a *authenticator) Token() (*oauth2.Token, error) {
	return a.config.Token(context.Background())
}
func (a *authenticator) Client(ctx context.Context) *http.Client {
	return a.config.Client(ctx)
}

// Get config for the golang oauth2 client.
func getConfig(provider CredentialProvider) (clientcredentials.Config, error) {
	// Get the credentials from the provider
	clientId, err := provider.ClientId()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	clientSecret, err := provider.ClientSecret()
	if err != nil {
		return clientcredentials.Config{}, err
	}

	// Setup the config
	return clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     defaultTokenEndpoint(),
	}, nil
}
