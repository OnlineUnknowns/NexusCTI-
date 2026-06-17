package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/opencti-lite/backend/database"
	"github.com/lib/pq"
)

type ATTACKMapping struct {
	ID            string    `json:"id"`
	TechniqueID   string    `json:"techniqueId"`
	TechniqueName string    `json:"techniqueName"`
	Tactic        string    `json:"tactic"`
	Platform      []string  `json:"platform"`
	EntityType    string    `json:"entityType"`
	EntityID      string    `json:"entityId"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (m *ATTACKMapping) Create() error {
	query := `
		INSERT INTO attack_mappings (technique_id, technique_name, tactic, platform, entity_type, entity_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	err := database.DB.QueryRow(
		query,
		m.TechniqueID,
		m.TechniqueName,
		m.Tactic,
		pq.Array(m.Platform),
		m.EntityType,
		m.EntityID,
	).Scan(&m.ID, &m.CreatedAt)
	return err
}

func DeleteATTACKMapping(id string) error {
	query := `DELETE FROM attack_mappings WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	return err
}

func ListATTACKMappings(filters map[string]interface{}) ([]*ATTACKMapping, error) {
	var whereClauses []string
	var args []interface{}
	argCounter := 1

	if val, ok := filters["entity_type"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("entity_type = $%d", argCounter))
		args = append(args, val)
		argCounter++
	}

	if val, ok := filters["entity_id"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("entity_id = $%d", argCounter))
		args = append(args, val)
		argCounter++
	}

	if val, ok := filters["tactic"]; ok && val != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("tactic = $%d", argCounter))
		args = append(args, val)
		argCounter++
	}

	whereQuery := ""
	if len(whereClauses) > 0 {
		whereQuery = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT id, technique_id, technique_name, tactic, platform, entity_type, entity_id, created_at
		FROM attack_mappings
		%s
		ORDER BY created_at DESC
	`, whereQuery)

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mappings []*ATTACKMapping
	for rows.Next() {
		var m ATTACKMapping
		err := rows.Scan(
			&m.ID,
			&m.TechniqueID,
			&m.TechniqueName,
			&m.Tactic,
			pq.Array(&m.Platform),
			&m.EntityType,
			&m.EntityID,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		mappings = append(mappings, &m)
	}
	return mappings, nil
}

func GetGroupedByTactic() (map[string][]*ATTACKMapping, error) {
	mappings, err := ListATTACKMappings(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	grouped := make(map[string][]*ATTACKMapping)
	for _, m := range mappings {
		grouped[m.Tactic] = append(grouped[m.Tactic], m)
	}
	return grouped, nil
}
