package bolapi

/*
	RETRY AFTER LOGIC: Because the http transport is hidden from the client we have to do something like the following to ge the retry-after header:

		res, err := client.Shipments.GetShipments(&shipments.GetShipmentsParams{Context: context.Background()})

		if err != nil {
			apiError, ok := err.(*runtime.APIError)
			if !ok {
				panic("HELP!")
			}

			rr := reflect.ValueOf(apiError.Response).FieldByName("resp").Elem()
			res := reflect.NewAt(rr.Type(), unsafe.Pointer(rr.UnsafeAddr())).Elem().Interface().(http.Response)

			fmt.Printf("Retry after: %s\n", res.Header.Get("Retry-After"))
		}
		fmt.Printf("%T\n", res)
*/
