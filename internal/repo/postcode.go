package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3" // Initialize go-sqlite3 library
	"github.com/philgresh/postcode-timezone/internal/model"
)

const (
	file           = "../data/db.sqlite3"
	commonErrorStr = "unable to get postcode from DB"
	getPostcode    = `
		SELECT z.id,
			z.code,
			z.state_id,
			z.accuracy,
			z.area_code,
			z.city,
			z.lat,
			z.lon,
			states.abbr AS state_abbr,
			states.name AS state_name
		FROM zipcodes AS z
			INNER JOIN
			states ON z.state_id = states.id
		WHERE z.code = ?;
	`
)

// GetPostcode attempts to get the given postcode details from the DB.
func GetPostcode(ctx context.Context, postcodeArg string) (*model.Postcode, error) {
	if postcodeArg == "" {
		return nil, fmt.Errorf("%s, postcode arg is required", commonErrorStr)
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("%s, database error: %w", commonErrorStr, err)
	}

	rows, err := db.QueryContext(ctx, getPostcode, postcodeArg)
	if err != nil {
		return nil, fmt.Errorf("%s, query row error: %w", commonErrorStr, err)
	}

	var postcode model.Postcode
	if err = scan.Row(&postcode, rows); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s, postcode does not exist", commonErrorStr)
		}

		return nil, fmt.Errorf("%s, query row scan error: %w", commonErrorStr, err)
	}

	return &postcode, nil
}
