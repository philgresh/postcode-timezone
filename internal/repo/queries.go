package repo

const (
	getPostcode = `
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
