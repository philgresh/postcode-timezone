package server

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Initialize go-sqlite3 library
	"github.com/philgresh/zipcode-timezone/internal/model"
)

const (
	file           = "../data/postal-codes-US/free_zipcode_data.sqlite3"
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

// GetPostcode
func GetPostcode(ctx context.Context, postcodeArg string) (*model.Postcode, error) {
	if postcodeArg == "" {
		return nil, fmt.Errorf("%s, postcode arg is required", commonErrorStr)
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("%s, database error: %s", commonErrorStr, err)
	}

	var (
		ID        string
		Code      string
		StateID   int32
		Accuracy  sql.NullInt16
		AreaCode  sql.NullString
		City      sql.NullString
		Lat       sql.NullFloat64
		Lon       sql.NullFloat64
		StateAbbr sql.NullString
		StateName sql.NullString
	)
	row := db.QueryRowContext(ctx, getPostcode, postcodeArg)
	if row.Err() != nil {
		return nil, fmt.Errorf("%s, query row error: %s", commonErrorStr, err)
	}

	if err = row.Scan(
		&ID,
		&Code,
		&StateID,
		&Accuracy,
		&AreaCode,
		&City,
		&Lat,
		&Lon,
		&StateAbbr,
		&StateName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s, postcode does not exist", commonErrorStr)
		}
		return nil, fmt.Errorf("%s, query row scan error: %s", commonErrorStr, err)
	}

	return &model.Postcode{
		ID:      ID,
		Code:    Code,
		StateID: StateID,
		// Accuracy:  int32(Accuracy),
		// AreaCode:  AreaCode,
		// City:      City,
		// Lat:       Lat,
		// Lon:       Lon,
		// StateAbbr: StateAbbr,
		// StateName: StateName,
	}, nil
}

//  func getRecentSearches(db *sql.DB, limit int) []Searches {
// 	var searches []Searches
// 	row, err := db.Query("SELECT * FROM search ORDER BY count LIMIT ?", limit)
// 	if err != nil {
// 			log.Fatal(err)
// 	}
// 	defer row.Close()
// 	for row.Next() { // Iterate and fetch the records from result cursor
// 			item := Searches{}
// 			err := row.Scan(&item.id, &item.count, &item.search)
// 			if err != nil {
// 					log.Fatal(err)
// 			}
// 			searches = append(searches, item)
// 	}
// 	return searches
// }
