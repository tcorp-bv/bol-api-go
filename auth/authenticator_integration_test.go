//+build integration

package auth

import (
	"testing"
)

func TestOauth(t *testing.T) {
	provider := EnvironmentProvider{
		ClientIdKey:     "BOL_CLIENT_ID",
		ClientSecretKey: "BOL_CLIENT_SECRET",
	}
	bolAuth, err := New(&provider)
	if err != nil {
		t.Fatal(err)
	}
	token, err := bolAuth.Token()
	if err != nil {
		t.Fatal(err)
	}
	if token == nil {
		t.Fatalf("received nil token from auth endpoint")
	}
	if !token.Valid() {
		t.Fatalf("received invalid token from auth endpoint")
	}
	if token.AccessToken == "" {
		t.Fatalf("received '' token from token endpoint")
	}
}
