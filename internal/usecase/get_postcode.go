package usecase

import (
	"errors"
	"fmt"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/philgresh/postcode-timezone/internal/present"
	"github.com/philgresh/postcode-timezone/internal/repo"
)

func GetPostcode(postcodeArg string) (*api.Postcode, error) {
	if postcodeArg == "" {
		return nil, getPostcodeError(errors.New("postcode required"))
	}

	modelPostcode, err := repo.GetPostcode(postcodeArg)
	if err != nil {
		return nil, getPostcodeError(err)
	}

	return present.ModelPostcodeToPostcode(modelPostcode)
}

func getPostcodeError(e error) error {
	return fmt.Errorf("Usecase.GetPostcode: unable to get postcode: %w", e)
}
