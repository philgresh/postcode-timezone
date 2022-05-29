package usecase

import (
	"context"
	"testing"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/stretchr/testify/require"
	"googlemaps.github.io/maps"
)

func TestGetPostcode(t *testing.T) {
	testcases := []struct {
		desc             string
		countryArg       api.Country
		postcodeArg      string
		expectedPostcode *api.Postcode
		expectedErr      string
	}{
		{
			desc:        "returns an error when a country arg is not provided",
			postcodeArg: "94108",
			expectedErr: "unable to get postcode: supported country required",
		},
		{
			desc:        "returns an error when a postcode arg is not provided",
			countryArg:  api.US,
			expectedErr: "unable to get postcode: postcode required",
		},
		{
			desc:        "returns an API postcode",
			countryArg:  api.US,
			postcodeArg: "94108",
			expectedPostcode: &api.Postcode{
				City: "San Francisco",
				Location: &maps.LatLng{
					Lat: 37.7929,
					Lng: -122.4079,
				},
				Postcode: "94108",
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		ctx := context.Background()

		t.Run(tc.desc, func(t *testing.T) {
			pc, err := GetPostcode(ctx, tc.countryArg, tc.postcodeArg)
			if tc.expectedErr != "" {
				require.Equal(t, tc.expectedErr, err.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, tc.expectedPostcode, pc)
		})
	}
}
