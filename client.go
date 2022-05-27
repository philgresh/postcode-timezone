package zipcode_timezone

import (
	"fmt"

	"github.com/philgresh/zipcode-timezone/api"
	"github.com/philgresh/zipcode-timezone/present"
)

type Client struct {
	country string
}

// ClientOption is the type of constructor options for NewClient(...).
type ClientArgs struct {
	country string
}

// NewClient constructs a new Client which can make requests to the zipcode-timezone DB
func NewClient(args *ClientArgs) (*Client, error) {
	if args.GetCountry() == "" {
		return nil, fmt.Errorf("Client cannot be initialized, country required")
	}

	return &Client{
		country: args.country,
	}, nil
}

func (c *Client) GetPostcode(postcode string) (*present.APIPostcode, error) {
	Postcode, err := api.GetPostcode(c.country, postcode)
	if err != nil {
		return nil, fmt.Errorf("unable to get postcode details, %s", err)
	}
	return Postcode, nil
}

func (args *ClientArgs) GetCountry() string {
	if args == nil {
		return ""
	}
	return args.country
}