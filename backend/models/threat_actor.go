package models

import (
	"time"

	"github.com/opencti-lite/backend/database"
	"github.com/lib/pq"
)

type ThreatActor struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	Aliases           []string   `json:"aliases"`
	Sophistication    string     `json:"sophistication"`
	ResourceLevel     string     `json:"resourceLevel"`
	PrimaryMotivation string     `json:"primaryMotivation"`
	CountryCode       string     `json:"countryCode"`
	Description       string     `json:"description"`
	CampaignCount     int        `json:"campaignCount"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}

func (t *ThreatActor) Create() error {
	query := `
		INSERT INTO threat_actors (name, aliases, sophistication, resource_level, primary_motivation, country_code, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`
	err := database.DB.QueryRow(
		query,
		t.Name,
		pq.Array(t.Aliases),
		t.Sophistication,
		t.ResourceLevel,
		t.PrimaryMotivation,
		t.CountryCode,
		t.Description,
	).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
	return err
}

func GetThreatActorByID(id string) (*ThreatActor, error) {
	query := `
		SELECT ta.id, ta.name, ta.aliases, ta.sophistication, ta.resource_level, ta.primary_motivation, ta.country_code, ta.description, ta.created_at, ta.updated_at, COUNT(c.id) AS campaign_count
		FROM threat_actors ta
		LEFT JOIN campaigns c ON ta.id = c.threat_actor_id
		WHERE ta.id = $1
		GROUP BY ta.id
	`
	row := database.DB.QueryRow(query, id)
	var t ThreatActor
	err := row.Scan(
		&t.ID,
		&t.Name,
		pq.Array(&t.Aliases),
		&t.Sophistication,
		&t.ResourceLevel,
		&t.PrimaryMotivation,
		&t.CountryCode,
		&t.Description,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.CampaignCount,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *ThreatActor) Update() error {
	query := `
		UPDATE threat_actors
		SET name = $1, aliases = $2, sophistication = $3, resource_level = $4, primary_motivation = $5, country_code = $6, description = $7, updated_at = NOW()
		WHERE id = $8
		RETURNING updated_at
	`
	err := database.DB.QueryRow(
		query,
		t.Name,
		pq.Array(t.Aliases),
		t.Sophistication,
		t.ResourceLevel,
		t.PrimaryMotivation,
		t.CountryCode,
		t.Description,
		t.ID,
	).Scan(&t.UpdatedAt)
	return err
}

func DeleteThreatActor(id string) error {
	query := `DELETE FROM threat_actors WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	return err
}

func ListThreatActors() ([]*ThreatActor, error) {
	query := `
		SELECT ta.id, ta.name, ta.aliases, ta.sophistication, ta.resource_level, ta.primary_motivation, ta.country_code, ta.description, ta.created_at, ta.updated_at, COUNT(c.id) AS campaign_count
		FROM threat_actors ta
		LEFT JOIN campaigns c ON ta.id = c.threat_actor_id
		GROUP BY ta.id
		ORDER BY ta.name ASC
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []*ThreatActor
	for rows.Next() {
		var t ThreatActor
		err := rows.Scan(
			&t.ID,
			&t.Name,
			pq.Array(&t.Aliases),
			&t.Sophistication,
			&t.ResourceLevel,
			&t.PrimaryMotivation,
			&t.CountryCode,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.CampaignCount,
		)
		if err != nil {
			return nil, err
		}
		actors = append(actors, &t)
	}
	return actors, nil
}
