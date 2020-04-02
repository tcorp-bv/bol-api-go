package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	defaultDir              = "/etc/secrets"
	defaultClientIdFile     = "clientId"
	defaultClientSecretFile = "clientSecret"
)

// The credentials to communicate with the bol API.
type CredentialProvider interface {
	// oauth ID for bol.com v3 api
	ClientId() (string, error)
	// oauth secret for bol.com v3 api
	ClientSecret() (string, error)
}

/*
	By default we use a FileProvider. This allows us to mount a kubernetes secret into the container.
	The mounted volume will create two files containing the clientId and clientSecret
	The FileProvider is responsible for reading them.
*/
// CredentialProvider implementation for a dir containing two files, each containing the secret as a string.
type FileProvider struct {
	// The directory with the files, this is usually the place where you mount the secret volume
	Dir string // Should equal /etc/secrets

	// The filename within Dir containing the ClientId
	ClientIdFile string // Should equal "clientId"

	// The filename within Dir containing the ClientSecret
	ClientSecretFile string // Should equal "clientSecret"
}

func NewFileProvider() CredentialProvider {
	return &FileProvider{Dir: defaultDir, ClientIdFile: defaultClientIdFile, ClientSecretFile: defaultClientSecretFile}
}

func (p *FileProvider) ClientId() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.ClientIdFile))
	return string(data), err
}

func (p *FileProvider) ClientSecret() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(p.Dir, p.ClientSecretFile))
	return string(data), err
}

// Simple raw value credential provider
type BasicProvider struct {
	ClientIdVal     string
	ClientSecretVal string
}

func (p *BasicProvider) ClientId() (string, error) {
	return p.ClientIdVal, nil
}

func (p *BasicProvider) ClientSecret() (string, error) {
	return p.ClientSecretVal, nil
}

// Fetches the credentials from the environment
type EnvironmentProvider struct {
	ClientIdKey     string
	ClientSecretKey string
}

func (p *EnvironmentProvider) ClientId() (string, error) {
	clientId, present := os.LookupEnv(p.ClientIdKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.ClientIdKey)
	}
	return clientId, nil
}

func (p *EnvironmentProvider) ClientSecret() (string, error) {
	clientSecret, present := os.LookupEnv(p.ClientSecretKey)
	if !present {
		return "", fmt.Errorf("environment variable '%s' not found", p.ClientSecretKey)
	}
	return clientSecret, nil
}
