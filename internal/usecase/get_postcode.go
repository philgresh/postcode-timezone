package usecase

import (
	"context"
	"fmt"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/philgresh/postcode-timezone/internal/present"
	"github.com/philgresh/postcode-timezone/internal/repo"
)

func GetPostcode(ctx context.Context, country api.Country, postcodeArg string) (*api.Postcode, error) {
	if postcodeArg == "" {
		return nil, getPostcodeError("postcode required")
	}

	if country == api.DoNotUse {
		return nil, getPostcodeError("country required")
	}

	if country != api.US {
		return nil, getPostcodeError("unsupported country")
	}

	modelPostcode, err := repo.GetPostcode(ctx, postcodeArg)
	if err != nil {
		return nil, getPostcodeError(err.Error())
	}

	return present.ModelPostcodeToPostcode(modelPostcode)
}

func getPostcodeError(s string) error {
	return fmt.Errorf("unable to get postcode : %s", s)
}
