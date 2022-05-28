package present

import (
	"fmt"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/philgresh/postcode-timezone/internal/model"
	"googlemaps.github.io/maps"
)

func ModelPostcodeToPostcode(m *model.Postcode) (*api.Postcode, error) {
	if m == nil {
		return nil, fmt.Errorf("unable to convert model postcode details to api postcode details, model object is required")
	}

	Postcode := &api.Postcode{
		City: m.GetCity(),
		Location: &maps.LatLng{
			Lat: m.GetLat(),
			Lng: m.GetLon(),
		},
		Postcode: m.GetCode(),
	}

	return Postcode, nil
}
