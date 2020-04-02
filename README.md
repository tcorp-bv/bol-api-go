# Bol-api-go
Generated golang API for [bol.com v3](https://api.bol.com/retailer/public/redoc/v3) using the [swagger spec](https://api.bol.com/retailer/public/apispec/v3).

We do not hold copyright over the API specification and [types.json](types.json) is from [bol.com](https://api.bol.com/retailer/public/apispec/v3).

## Using the API
```go
// Setup the authentication
bolApi, err := api.New(&auth.EnvironmentProvider{ClientIdKey: "CLIENT_ID", ClientSecretKey: "CLIENT_SECRET"})
if err != nil {
	//handle error
}
// Get the client (you should do this once)
client := bolApi.GetClient()
	
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

## Todo
* Rate-limit management

## Security notes
As this project uses go-swagger, it has external dependencies. In production you should audit go.sum to make sure that these are not malicious.