package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/opencti-lite/backend/database"
	"github.com/lib/pq"
)

type IOC struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Value       string     `json:"value"`
	TLPLevel    string     `json:"tlpLevel"`
	Confidence  int        `json:"confidence"`
	Tags        []string   `json:"tags"`
	Source      string     `json:"source"`
	Description string     `json:"description"`
	FirstSeen   *time.Time `json:"firstSeen"`
	LastSeen    *time.Time `json:"lastSeen"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func (i *IOC) Create() error {
	query := `
		INSERT INTO iocs (type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`
	err := database.DB.QueryRow(
		query,
		i.Type,
		i.Value,
		i.TLPLevel,
		i.Confidence,
		pq.Array(i.Tags),
		i.Source,
		i.Description,
		i.FirstSeen,
		i.LastSeen,
	).Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return err
}

func GetIOCByID(id string) (*IOC, error) {
	query := `
		SELECT id, type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen, created_at, updated_at
		FROM iocs
		WHERE id = $1
	`
	row := database.DB.QueryRow(query, id)
	var i IOC
	err := row.Scan(
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
	return &i, nil
}

func (i *IOC) Update() error {
	query := `
		UPDATE iocs
		SET type = $1, value = $2, tlp_level = $3, confidence = $4, tags = $5, source = $6, description = $7, first_seen = $8, last_seen = $9, updated_at = NOW()
		WHERE id = $10
		RETURNING updated_at
	`
	err := database.DB.QueryRow(
		query,
		i.Type,
		i.Value,
		i.TLPLevel,
		i.Confidence,
		pq.Array(i.Tags),
		i.Source,
		i.Description,
		i.FirstSeen,
		i.LastSeen,
		i.ID,
	).Scan(&i.UpdatedAt)
	return err
}

func DeleteIOC(id string) error {
	query := `DELETE FROM iocs WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	return err
}

func ListIOCs(filters map[string]interface{}, page, limit int) ([]*IOC, int, error) {
	var whereClauses []string
	var args []interface{}
	argCounter := 1

	if val, ok := filters["type"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("type = $%d", argCounter))
		args = append(args, val)
		argCounter++
	}

	if val, ok := filters["tlp"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("tlp_level = $%d", argCounter))
		args = append(args, val)
		argCounter++
	}

	if val, ok := filters["tag"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("$%d = ANY(tags)", argCounter))
		args = append(args, val)
		argCounter++
	}

	if val, ok := filters["q"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("search_vector @@ plainto_tsquery('english', $%d)", argCounter))
		args = append(args, val)
		argCounter++
	}

	whereQuery := ""
	if len(whereClauses) > 0 {
		whereQuery = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Get total count
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM iocs %s", whereQuery)
	var total int
	err := database.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * limit
	limitQuery := fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
	args = append(args, limit, offset)

	query := fmt.Sprintf(`
		SELECT id, type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen, created_at, updated_at
		FROM iocs
		%s
		ORDER BY created_at DESC
		%s
	`, whereQuery, limitQuery)

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var iocs []*IOC
	for rows.Next() {
		var i IOC
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
			return nil, 0, err
		}
		iocs = append(iocs, &i)
	}

	return iocs, total, nil
}

func BulkCreateIOCs(iocs []IOC) (int, error) {
	if len(iocs) == 0 {
		return 0, nil
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO iocs (type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	count := 0
	for _, i := range iocs {
		_, err = stmt.Exec(
			i.Type,
			i.Value,
			i.TLPLevel,
			i.Confidence,
			pq.Array(i.Tags),
			i.Source,
			i.Description,
			i.FirstSeen,
			i.LastSeen,
		)
		if err != nil {
			return count, err
		}
		count++
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return count, nil
}
