# Requires the BOL_CLIENT_ID and BOL_CLIENT_SECRET environment variables to be set
# Executes tests that run against the actual bol.com retailer API
integration-test:
	go test ./... -tags=integration