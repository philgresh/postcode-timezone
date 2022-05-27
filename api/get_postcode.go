package api

import (
	"fmt"

	"github.com/philgresh/zipcode-timezone/present"
)


func GetPostcode(country, postcode string) (*present.APIPostcode, error) {
	if postcode == "" {
		return nil, fmt.Errorf("unable to get postcode details, postcode required")
	}



	return nil, nil
}