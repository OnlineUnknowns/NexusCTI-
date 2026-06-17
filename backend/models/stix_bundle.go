package models

import (
	"encoding/json"
	"time"

	"github.com/opencti-lite/backend/database"
)

type STIXBundle struct {
	ID          string          `json:"id"`
	SpecVersion string          `json:"specVersion"`
	BundleJSON  json.RawMessage `json:"bundleJson"`
	CreatedAt   time.Time       `json:"createdAt"`
}

func (s *STIXBundle) Create() error {
	query := `
		INSERT INTO stix_bundles (spec_version, bundle_json)
		VALUES ($1, $2)
		RETURNING id, created_at
	`
	err := database.DB.QueryRow(
		query,
		s.SpecVersion,
		s.BundleJSON,
	).Scan(&s.ID, &s.CreatedAt)
	return err
}

func GetSTIXBundleByID(id string) (*STIXBundle, error) {
	query := `
		SELECT id, spec_version, bundle_json, created_at
		FROM stix_bundles
		WHERE id = $1
	`
	row := database.DB.QueryRow(query, id)
	var s STIXBundle
	err := row.Scan(
		&s.ID,
		&s.SpecVersion,
		&s.BundleJSON,
		&s.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func ListSTIXBundles() ([]*STIXBundle, error) {
	query := `
		SELECT id, spec_version, bundle_json, created_at
		FROM stix_bundles
		ORDER BY created_at DESC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bundles []*STIXBundle
	for rows.Next() {
		var s STIXBundle
		err := rows.Scan(
			&s.ID,
			&s.SpecVersion,
			&s.BundleJSON,
			&s.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		bundles = append(bundles, &s)
	}
	return bundles, nil
}
