package present

import (
	"database/sql"
	"testing"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/philgresh/postcode-timezone/internal/model"
	"github.com/stretchr/testify/require"
	"googlemaps.github.io/maps"
)

func TestGetPostcode(t *testing.T) {
	testcases := []struct {
		desc             string
		modelPostcode    *model.Postcode
		expectedPostcode *api.Postcode
		expectedErr      string
	}{
		{
			desc:        "returns an error if the model is not provided",
			expectedErr: "ModelPostcodeToPostcode: unable to convert model postcode to api postcode, model struct is required",
		},
		{
			desc: "successfully converts a model Postcode to an Postcode",
			modelPostcode: &model.Postcode{
				ID:        4251,
				Code:      stringToNullString("94108"),
				StateID:   5,
				City:      stringToNullString("San Francisco"),
				Lat:       37.7929,
				Lon:       -122.4079,
				Accuracy:  4,
				StateAbbr: stringToNullString("CA"),
				StateName: stringToNullString("California"),
			},
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
			Postcode, err := ModelPostcodeToPostcode(tc.modelPostcode)
			if tc.expectedErr != "" {
				require.Equal(t, tc.expectedErr, err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tc.expectedPostcode, Postcode)
			}
		})
	}
}

func stringToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
