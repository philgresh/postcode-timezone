package repo

import (
	"database/sql"
	"testing"

	"github.com/philgresh/postcode-timezone/internal/model"
	"github.com/stretchr/testify/require"
)

func TestGetPostcode(t *testing.T) {
	testcases := []struct {
		desc             string
		postcodeArg      string
		expectedErr      string
		expectedPostcode *model.Postcode
	}{
		{
			desc:        "returns an error if no postcode arg is provided",
			expectedErr: "Repo.GetPostcode: unable to get postcode from DB: postcode arg is required",
		},
		{
			desc:        "returns an error if the postcode does not exist",
			expectedErr: "Repo.GetPostcode: unable to get postcode from DB: postcode '00000' does not exist",
			postcodeArg: "00000",
		},
		{
			desc:        "returns a postcode if it exists",
			postcodeArg: "94108",
			expectedPostcode: &model.Postcode{
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

func stringToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
