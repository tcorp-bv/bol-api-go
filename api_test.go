//+build manual_test

// THIS IS A MANUAL TEST, NOT AN INTEGRATION/UNIT TEST

package bolapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/tcorp-bv/bol-api-go/auth"
	"github.com/tcorp-bv/bol-api-go/client/shipments"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	demoURL      = "api.bol.com"
	demoBasePath = "/retailer-demo/"
)

type URLRewriter struct {
	Proxied http.RoundTripper
	Verbose bool
}

func (lrt URLRewriter) RoundTrip(req *http.Request) (res *http.Response, err error) {
	println("rewriting")
	if len(req.URL.Path) < 15 {
		return lrt.Proxied.RoundTrip(req)
	}
	if req.URL.Path[:15] == "/retailer-demo/" { //fix demo url
		req.URL.Path = strings.Replace(req.URL.Path, "/retailer-demo/retailer/", "/retailer-demo/", 1)
	}
	println(req.URL.String())
	// Send the request, get the response (or the error)
	res, err = lrt.Proxied.RoundTrip(req)
	//print the response
	if lrt.Verbose {
		buf := bytes.Buffer{}
		buf.ReadFrom(res.Body)
		println(buf.String())
	}
	return res, err
}
func TestDemo(t *testing.T) {
	provider := auth.EnvironmentProvider{ClientIdKey: "CLIENT_ID", ClientSecretKey: "CLIENT_SECRET"}

	bolAPI, err := NewWithTransport(&provider, demoURL, demoBasePath, func(rt http.RoundTripper) http.RoundTripper {
		return URLRewriter{Proxied: rt, Verbose: false}
	})
	if err != nil {
		t.Fatal(err)
	}
	client := bolAPI.GetClient()
	for {
		time.Sleep(1 * time.Second)
		fbb := "FBB"
		res, err := client.Shipments.GetShipments(&shipments.GetShipmentsParams{FulfilmentMethod: &fbb, Context: context.Background()})
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Printf("%+v\n", res.Payload.Shipments)
	}
}
func TestProd(t *testing.T) {
	provider := auth.EnvironmentProvider{ClientIdKey: "CLIENT_ID", ClientSecretKey: "CLIENT_SECRET"}

	bolAPI, err := New(&provider)
	if err != nil {
		t.Fatal(err)
	}
	client := bolAPI.GetClient()
	for {
		time.Sleep(1 * time.Second)
		fbb := "FBB"
		res, err := client.Shipments.GetShipments(&shipments.GetShipmentsParams{FulfilmentMethod: &fbb, Context: context.Background()})
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Printf("%+v\n", res.Payload.Shipments)
	}
}
