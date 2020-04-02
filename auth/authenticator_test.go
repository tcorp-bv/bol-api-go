package auth

import "testing"

const (
	testClientId     = "oRNWbHFXtAECmhnZmEndcjLIaSKbRMVE"
	testClientSecret = "aQHPOnmYkPZNgeRziPnQyyOJYytUbcFBVJBvbMKoDdpPqaZbaOiLUTWzPAkpPsZFZbJHrcoltdgpZolyNcgvvBaKcmkqFjucFzXhDONTsPAtHHyccQlLUZpkOuywMiOycDWcCySFsgpDiyGnCWCZJkNTtVdPxbSUTWVIFQiUxaPDYDXRQAVVTbSVZArAZkaLDLOoOvPzxSdhnkkJWzlQDkqsXNKfAIgAldrmyfROSyCGMCfvzdQdUQEaYZTPEoA"
)

func TestDefaultEndpoint(t *testing.T) {
	if defaultTokenEndpoint() != "https://login.bol.com/token" {
		t.Errorf("Expected https://login.bol.com, got %s as token endpoint", defaultTokenEndpoint())
	}
}

func TestGetConfig(t *testing.T) {
	c, err := getConfig(&BasicProvider{ClientIdVal: testClientId, ClientSecretVal: testClientSecret})
	if err != nil {
		t.Error(err)
	}
	if c.ClientID != testClientId || c.ClientSecret != testClientSecret || c.TokenURL != defaultTokenEndpoint() {
		t.Error("Client credentials not properly parsed")
	}
}
