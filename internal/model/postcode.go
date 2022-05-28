package model

import "database/sql"

type Postcode struct {
	ID        int32          `db:"id"`
	Code      sql.NullString `db:"code"` // Postcode/ZIP code
	StateID   int32          `db:"state_id"`
	StateAbbr sql.NullString `db:"state_abbr"`
	StateName sql.NullString `db:"state_name"`
	City      sql.NullString `db:"city"`
	AreaCode  sql.NullString `db:"area_code"`
	Lat       float64        `db:"lat"`
	Lon       float64        `db:"lon"`
	Accuracy  int32          `db:"accuracy"`
}

func (p *Postcode) GetID() int32 {
	if p == nil {
		return 0
	}

	return p.ID
}

func (p *Postcode) GetCode() string {
	if p == nil {
		return ""
	}

	return nullStringToString(p.Code)
}

func (p *Postcode) GetStateID() int32 {
	if p == nil {
		return 0
	}

	return p.StateID
}

func (p *Postcode) GetStateAbbr() string {
	if p == nil {
		return ""
	}

	return nullStringToString(p.StateAbbr)
}

func (p *Postcode) GetStateName() string {
	if p == nil {
		return ""
	}

	return nullStringToString(p.StateName)
}

func (p *Postcode) GetCity() string {
	if p == nil {
		return ""
	}

	return nullStringToString(p.City)
}

func (p *Postcode) GetAreaCode() string {
	if p == nil {
		return ""
	}

	return nullStringToString(p.AreaCode)
}

func (p *Postcode) GetLat() float64 {
	if p == nil {
		return 0
	}

	return p.Lat
}

func (p *Postcode) GetLon() float64 {
	if p == nil {
		return 0
	}

	return p.Lon
}

func (p *Postcode) GetAccuracy() int32 {
	if p == nil {
		return 0
	}

	return p.Accuracy
}

func nullStringToString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}

	return ns.String
}
