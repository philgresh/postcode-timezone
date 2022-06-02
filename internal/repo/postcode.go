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
	db = "../data/db.sqlite3"
)

// GetPostcode attempts to get the given postcode details from the DB.
func GetPostcode(postcodeArg string) (*model.Postcode, error) {
	if postcodeArg == "" {
		return nil, getPostcodeError(errors.New("postcode arg is required"))
	}

	db, err := sql.Open("sqlite3", db)
	if err != nil {
		return nil, getPostcodeError(fmt.Errorf("database error: %w", err))
	}

	rows, err := db.Query(getPostcode, postcodeArg)
	if err != nil {
		return nil, getPostcodeError(fmt.Errorf("query row error: %w", err))
	}

	var postcode model.Postcode
	if err = scan.RowStrict(&postcode, rows); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, getPostcodeError(fmt.Errorf("postcode '%s' does not exist", postcodeArg))
		}

		return nil, getPostcodeError(fmt.Errorf("query row scan error: %w", err))
	}

	return &postcode, nil
}

// QueryPostcodeRows is a generic query function that attempts to assign return a slice of model.Postcode
// values from a given query string and args.
func QueryPostcodeRows(queryStr string, args ...any) ([]*model.Postcode, error) {
	if queryStr == "" {
		return nil, errors.New("Repo.QueryPostcodeRows: query string is required")
	}

	db, err := sql.Open("sqlite3", db)
	if err != nil {
		return nil, fmt.Errorf("Repo.QueryPostcodeRows: database error: %w", err)
	}

	rows, err := db.Query(queryStr, args...)
	if err != nil {
		return nil, fmt.Errorf("Repo.QueryPostcodeRows: query row error: %w", err)
	}

	var output []model.Postcode
	if err = scan.Rows(&output, rows); err != nil {
		return nil, fmt.Errorf("Repo.QueryPostcodeRows: query row scan error: %w", err)
	}

	// Convert output to slice of pointers
	returnVals := make([]*model.Postcode, len(output))

	for i, outputVal := range output {
		outputVal := outputVal
		returnVals[i] = &outputVal
	}

	return returnVals, nil
}

func getPostcodeError(e error) error {
	return fmt.Errorf("Repo.GetPostcode: unable to get postcode from DB: %w", e)
}
