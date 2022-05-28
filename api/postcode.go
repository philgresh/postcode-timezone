package api

import "googlemaps.github.io/maps"

type Postcode struct {
	City     string       `json:"city"`
	State    string       `json:"state"`
	Postcode string       `json:"postcode"`
	Location *maps.LatLng `json:"location"` // The centroid of the postcode
}
