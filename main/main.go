package main

import (
	"context"
	bolapi "github.com/tcorp-bv/bol-api-go"
	"github.com/tcorp-bv/bol-api-go/auth"
	"github.com/tcorp-bv/bol-api-go/client/shipments"
)

func main() {
	api, err := bolapi.New(&auth.EnvironmentProvider{ClientIdKey: "CLIENT_ID", ClientSecretKey: "CLIENT_SECRET"})
	if err != nil {
		panic(err)
	}
	// Get the client (you should do this once)
	client := api.GetClient()

	// Use the API
	res, err := client.Shipments.GetShipments(&shipments.GetShipmentsParams{Context: context.Background()})
	if err != nil {
		panic(err)
	}
	for _, s := range res.Payload.Shipments {
		println(s.ShipmentID)
	}
}
