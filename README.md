# Bol-api-go
Generated golang API for [bol.com v3](https://api.bol.com/retailer/public/redoc/v3) using the [swagger spec](https://api.bol.com/retailer/public/apispec/v3).

We do not hold copyright over the API specification and [types.json](types.json) is from [bol.com](https://api.bol.com/retailer/public/apispec/v3).

## Using the API
Add `github.com/tcorp-bv/bol-api-go v1.6.0` to go.mod.

```go
import (
    bolapi "github.com/tcorp-bv/bol-api-go"
)
// Setup the authentication
api, err := bolapi.New(&auth.EnvironmentProvider{ClientIdKey: "CLIENT_ID", ClientSecretKey: "CLIENT_SECRET"})
if err != nil {
	//handle error
}
// Get the client (you should do this once)
client := api.GetClient()
	
// Use the API
res, err := client.Shipments.GetShipments(&shipments.GetShipmentsParams{Context:context.Background()})
if err != nil {
	// handle error
}
for _, s := range res.Payload.Shipments {
	println(s.ShipmentID)
}
```


## Regenerating the API
[Go-swagger](https://github.com/go-swagger/go-swagger) was used for the client generation. Make sure it is accessible as `swagger`.
```shell script
swagger generate client -f types.json
```

In the current version we had to change `get-process-status` to `get-process-status-list` due to the duplicate.

## Notes on rate limiting
Rate limits of the bol.com api are shared between all your clientIds and are extremely low.
Because of this, we recommend having a single service that consumes your bol.com api and indexes all resources.

This library is opinionated in the way it handles rate limiting.
Requests that fail due to rate limiting will be tried again a certain amount of time.
Because of this, requests could take a very long time. 

The default behavior is to retry 10 times.