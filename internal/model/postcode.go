package model

type Postcode struct {
	ID string `db:"id"`
	Code string `db:"code"` // Postcode
	StateID int32 `db:"state_id"`
	StateAbbr string `db:"state_abbr"`
	StateName string `db:"state_name"`
	City string `db:"city"`
	AreaCode string `db:"area_code"`
	Lat float64 `db:"lat"`
	Lon	float64 `db:"lon"`
	Accuracy int32 `db:"accuracy"`
}

func (p *Postcode) GetCode() string {
	if p == nil {
		return ""
	}
	return p.Code
}

func (p *Postcode) GetCity() string {
	if p == nil {
		return ""
	}
	return p.City
}

func (p *Postcode) GetStateID() int32 {
	if p == nil {
		return 0
	}
	return p.StateID
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