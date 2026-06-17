package services

import (
	"github.com/opencti-lite/backend/database"
	"github.com/opencti-lite/backend/models"
	"github.com/lib/pq"
)

func SearchIOCs(query string) ([]*models.IOC, error) {
	sqlQuery := `
		SELECT id, type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen, created_at, updated_at
		FROM iocs
		WHERE search_vector @@ plainto_tsquery('english', $1)
		ORDER BY ts_rank(search_vector, plainto_tsquery('english', $1)) DESC
		LIMIT 50
	`
	rows, err := database.DB.Query(sqlQuery, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var iocs []*models.IOC
	for rows.Next() {
		var i models.IOC
		err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Value,
			&i.TLPLevel,
			&i.Confidence,
			pq.Array(&i.Tags),
			&i.Source,
			&i.Description,
			&i.FirstSeen,
			&i.LastSeen,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		iocs = append(iocs, &i)
	}
	return iocs, nil
}
