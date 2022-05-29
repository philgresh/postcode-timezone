package usecase

import (
	"testing"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/stretchr/testify/require"
	"googlemaps.github.io/maps"
)

func TestGetPostcode(t *testing.T) {
	testcases := []struct {
		desc             string
		postcodeArg      string
		expectedPostcode *api.Postcode
		expectedErr      string
	}{
		{
			desc:        "returns an error when a postcode arg is not provided",
			expectedErr: "Usecase.GetPostcode: unable to get postcode: postcode required",
		},
		{
			desc:        "returns an API postcode",
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

		t.Run(tc.desc, func(t *testing.T) {
			pc, err := GetPostcode(tc.postcodeArg)
			if tc.expectedErr != "" {
				require.Equal(t, tc.expectedErr, err.Error())
			} else {
				require.Nil(t, err)
			}
			require.Equal(t, tc.expectedPostcode, pc)
		})
	}
}
