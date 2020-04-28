//+build integration

package bolapi

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/matryer/is"
	"github.com/tcorp-bv/bol-api-go/client/shipments"
	"github.com/tcorp-bv/bol-api-go/models"
	"testing"
	"time"
)

func TestGetShipments(t *testing.T) {
	is := is.New(t)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Minute)
	api := *GetTestAPI(t, false, false)
	// Get the client (you should do this once)
	client := api.GetClient()

	// Use the API
	fMethod := "FBB"
	req := &shipments.GetShipmentsParams{Context: ctx, FulfilmentMethod: &fMethod}
	res, err := client.Shipments.GetShipments(req)
	is.NoErr(err)

	expectedDt, err := strfmt.ParseDateTime("2019-08-06T16:27:41+02:00")
	is.NoErr(err)
	expected := []*models.ReducedShipment{
		{
			ShipmentID:   953992381,
			ShipmentDate: expectedDt,
			ShipmentItems: []*models.ReducedShipmentItem{
				{
					OrderItemID: "2070686066",
					OrderID:     "1045972070",
				},
			},
			Transport: &models.ReducedTransport{
				TransportID: 309119453,
			},
		},
	}
	is.Equal(res.Payload.Shipments, expected)
}

func TestGetShipment(t *testing.T) {
	is := is.New(t)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Minute)
	api := *GetTestAPI(t, false, false)
	// Get the client (you should do this once)
	client := api.GetClient()

	// Use the API
	req := &shipments.GetShipmentParams{Context: ctx, ShipmentID: 914587795}
	res, err := client.Shipments.GetShipment(req)
	is.NoErr(err)

	shipmentDate, err := strfmt.ParseDateTime("2018-04-20T19:20:11+02:00")
	is.NoErr(err)
	orderDate, err := strfmt.ParseDateTime("2018-04-20T16:13:31+02:00")
	is.NoErr(err)
	latestDeliveryDate, err := strfmt.ParseDateTime("2018-04-21T00:00:00+02:00")
	is.NoErr(err)
	expected := models.Shipment{
		ShipmentID:        914587795,
		PickUpPoint:       true,
		ShipmentDate:      shipmentDate,
		ShipmentReference: "Shipment1",
		ShipmentItems: []*models.ShipmentItem{
			{
				OrderItemID:        "6107434013",
				OrderID:            "7616222250",
				OrderDate:          orderDate,
				LatestDeliveryDate: latestDeliveryDate,
				Ean:                "8712626055150",
				Title:              "Star Wars - Original Trilogy",
				Quantity:           3,
				OfferPrice:         104.97,
				OfferCondition:     "NEW",
				OfferReference:     "Test1",
				FulfilmentMethod:   "FBR",
			},
		},
		Transport: &models.Transport{
			TransportID:     358612589,
			TransporterCode: "TNT",
			TrackAndTrace:   "3SAOLD1234567",
		},
		CustomerDetails: &models.CustomerDetails{
			PickUpPointName:     "Albert Heijn: UTRECHT",
			SalutationCode:      "01",
			FirstName:           "Anakin",
			Surname:             "Skywalker",
			StreetName:          "Tatooinestraat",
			HouseNumber:         "100",
			HouseNumberExtended: "B",
			ZipCode:             "3528BJ",
			City:                "Utrecht",
			CountryCode:         "NL",
			Email:               "25whxgzlkmvotjhskwf5x27wlrldny@verkopen.test2.bol.com",
			DeliveryPhoneNumber: "0612345678",
		},
		BillingDetails: &models.CustomerDetails{
			SalutationCode:          "01",
			FirstName:               "Anakin",
			Surname:                 "Skywalker",
			StreetName:              "Tatooinestraat",
			HouseNumber:             "100",
			HouseNumberExtended:     "B",
			AddressSupplement:       "A",
			ExtraAddressInformation: "Extra informatie",
			ZipCode:                 "3528 BJ",
			City:                    "UTRECHT",
			CountryCode:             "NL",
			Email:                   "25whxgzlkmvotjhskwf5x27wlrldny@verkopen.test2.bol.com",
			Company:                 "Pieter Post",
			VatNumber:               "NL123456789B01",
			ChamberOfCommerceNumber: "99887766",
			OrderReference:          "Mijn order ref",
		},
	}
	is.Equal(*res.Payload, expected)
}
