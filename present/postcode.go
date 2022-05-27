package present

import (
	"fmt"

	"github.com/philgresh/zipcode-timezone/internal/model"
	"googlemaps.github.io/maps"
)

type APIPostcode struct {
	City     string      `json:"city"`
	State    string      `json:"state"`
	Postcode string      `json:"postcode"`
	Location maps.LatLng `json:"location"` // The centroid of the postcode
}

func ModelPostcodeToAPIPostcode(m *model.Postcode) (*APIPostcode, error) {
	if m == nil {
		return nil, fmt.Errorf("unable to convert model postcode details to api postcode details, model object is required")
	}

	Postcode := &APIPostcode{
		City: m.GetCity(),
		Location: maps.LatLng{
			Lat: m.GetLat(),
			Lng: m.GetLon(),
		},
		Postcode: m.GetCode(),
	}

	return Postcode, nil
}
