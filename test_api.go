package bolapi

import (
	"bytes"
	"fmt"
	"github.com/tcorp-bv/bol-api-go/auth"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	demoURL      = "api.bol.com"
	demoBasePath = "/retailer-demo/"
)

func GetTestAPI(t *testing.T, verbose bool, production bool) *BolApi {
	provider := auth.EnvironmentProvider{ClientIdKey: "BOL_CLIENT_ID", ClientSecretKey: "BOL_CLIENT_SECRET"}
	path := demoBasePath
	if production {
		path = ""
	}
	bolAPI, err := NewWithTransport(&provider, demoURL, path, func(t http.RoundTripper) http.RoundTripper {
		return URLRewriter{Proxied: t, Verbose: verbose, Production: production}
	})
	if err != nil {
		t.Fatal(err)
	}
	return &bolAPI
}

// Little bit of to connect to the demo environment
type URLRewriter struct {
	Proxied    http.RoundTripper
	Production bool
	Verbose    bool
}

func (lrt URLRewriter) RoundTrip(req *http.Request) (res *http.Response, err error) {
	if !lrt.Production && req.URL.Path[:15] == "/retailer-demo/" { //fix demo url
		req.URL.Path = strings.Replace(req.URL.Path, "/retailer-demo/retailer/", "/retailer-demo/", 1)
	}
	// Send the request, get the response (or the error)
	res, err = lrt.Proxied.RoundTrip(req)
	//print the response
	if lrt.Verbose {
		buf := bytes.Buffer{}
		buf.ReadFrom(res.Body)
		res.Body.Close()
		res.Body = ioutil.NopCloser(bytes.NewBufferString(buf.String()))
		fmt.Printf("%s (%d): %s\n", req.URL.String(), res.StatusCode, buf.String())
	}
	return res, err
}
