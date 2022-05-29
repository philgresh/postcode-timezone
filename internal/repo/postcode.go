package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/blockloop/scan"
	_ "github.com/mattn/go-sqlite3" // Initialize go-sqlite3 library
	"github.com/philgresh/postcode-timezone/internal/model"
)

const (
	file = "../data/db.sqlite3"
)

// GetPostcode attempts to get the given postcode details from the DB.
func GetPostcode(postcodeArg string) (*model.Postcode, error) {
	if postcodeArg == "" {
		return nil, getPostcodeError(errors.New("postcode arg is required"))
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, getPostcodeError(fmt.Errorf("database error: %w", err))
	}

	rows, err := db.Query(getPostcode, postcodeArg)
	if err != nil {
		return nil, getPostcodeError(fmt.Errorf("query row error: %w", err))
	}

	var postcode model.Postcode
	if err = scan.Row(&postcode, rows); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, getPostcodeError(fmt.Errorf("postcode '%s' does not exist", postcodeArg))
		}

		return nil, getPostcodeError(fmt.Errorf("query row scan error: %w", err))
	}

	return &postcode, nil
}

func getPostcodeError(e error) error {
	return fmt.Errorf("Repo.GetPostcode: unable to get postcode from DB: %w", e)
}
