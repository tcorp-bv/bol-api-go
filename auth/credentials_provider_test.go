package auth

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Make sure that the file provider can actually read the two files
func TestFileProvider(t *testing.T) {
	clientId := "oRNWbHFXtAECmhnZmEndcjLIaSKbRMVE"
	clientSecret := "aQHPOnmYkPZNgeRziPnQyyOJYytUbcFBVJBvbMKoDdpPqaZbaOiLUTWzPAkpPsZFZbJHrcoltdgpZolyNcgvvBaKcmkqFjucFzXhDONTsPAtHHyccQlLUZpkOuywMiOycDWcCySFsgpDiyGnCWCZJkNTtVdPxbSUTWVIFQiUxaPDYDXRQAVVTbSVZArAZkaLDLOoOvPzxSdhnkkJWzlQDkqsXNKfAIgAldrmyfROSyCGMCfvzdQdUQEaYZTPEoA"

	// Setup temp dir
	dirName, err := ioutil.TempDir(os.TempDir(), "test-")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		_ = os.Remove(dirName)
	}()
	// Setup the files
	ioutil.WriteFile(filepath.Join(dirName, "clientId"), []byte(clientId), 0644)
	ioutil.WriteFile(filepath.Join(dirName, "clientSecret"), []byte(clientSecret), 0644)

	fp := FileProvider{Dir: dirName, ClientIdFile: "clientId", ClientSecretFile: "clientSecret"}
	testProvider(t, fp, clientId, clientSecret)
}

func testProvider(t *testing.T, provider FileProvider, clientId string, clientSecret string) {
	if cId, err := provider.ClientId(); err != nil || cId != clientId {
		t.Error("File provider does not read clientId file properly. ", err)
	}
	if cSecret, err := provider.ClientId(); err != nil || cSecret != clientId {
		t.Error("File provider does not read clientId file properly. ", err)
	}
}
