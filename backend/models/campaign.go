package models

import (
	"time"

	"github.com/opencti-lite/backend/database"
)

type Campaign struct {
	ID                       string     `json:"id"`
	Name                     string     `json:"name"`
	Description              string     `json:"description"`
	FirstSeen                *time.Time `json:"firstSeen"`
	LastSeen                 *time.Time `json:"lastSeen"`
	Objective                string     `json:"objective"`
	ThreatActorID            *string    `json:"threatActorId"`
	ThreatActorName          *string    `json:"threatActorName"`
	ThreatActorSophistication *string   `json:"threatActorSophistication"`
	CreatedAt                time.Time  `json:"createdAt"`
	UpdatedAt                time.Time  `json:"updatedAt"`
}

func (c *Campaign) Create() error {
	query := `
		INSERT INTO campaigns (name, description, first_seen, last_seen, objective, threat_actor_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`
	err := database.DB.QueryRow(
		query,
		c.Name,
		c.Description,
		c.FirstSeen,
		c.LastSeen,
		c.Objective,
		c.ThreatActorID,
	).Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt)
	return err
}

func GetCampaignByID(id string) (*Campaign, error) {
	query := `
		SELECT c.id, c.name, c.description, c.first_seen, c.last_seen, c.objective, c.threat_actor_id, ta.name, ta.sophistication, c.created_at, c.updated_at
		FROM campaigns c
		LEFT JOIN threat_actors ta ON c.threat_actor_id = ta.id
		WHERE c.id = $1
	`
	row := database.DB.QueryRow(query, id)
	var cmp Campaign
	err := row.Scan(
		&cmp.ID,
		&cmp.Name,
		&cmp.Description,
		&cmp.FirstSeen,
		&cmp.LastSeen,
		&cmp.Objective,
		&cmp.ThreatActorID,
		&cmp.ThreatActorName,
		&cmp.ThreatActorSophistication,
		&cmp.CreatedAt,
		&cmp.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &cmp, nil
}

func (c *Campaign) Update() error {
	query := `
		UPDATE campaigns
		SET name = $1, description = $2, first_seen = $3, last_seen = $4, objective = $5, threat_actor_id = $6, updated_at = NOW()
		WHERE id = $7
		RETURNING updated_at
	`
	err := database.DB.QueryRow(
		query,
		c.Name,
		c.Description,
		c.FirstSeen,
		c.LastSeen,
		c.Objective,
		c.ThreatActorID,
		c.ID,
	).Scan(&c.UpdatedAt)
	return err
}

func DeleteCampaign(id string) error {
	query := `DELETE FROM campaigns WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	return err
}

func ListCampaigns() ([]*Campaign, error) {
	query := `
		SELECT c.id, c.name, c.description, c.first_seen, c.last_seen, c.objective, c.threat_actor_id, ta.name, ta.sophistication, c.created_at, c.updated_at
		FROM campaigns c
		LEFT JOIN threat_actors ta ON c.threat_actor_id = ta.id
		ORDER BY c.first_seen DESC NULLS LAST
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []*Campaign
	for rows.Next() {
		var cmp Campaign
		err := rows.Scan(
			&cmp.ID,
			&cmp.Name,
			&cmp.Description,
			&cmp.FirstSeen,
			&cmp.LastSeen,
			&cmp.Objective,
			&cmp.ThreatActorID,
			&cmp.ThreatActorName,
			&cmp.ThreatActorSophistication,
			&cmp.CreatedAt,
			&cmp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, &cmp)
	}
	return campaigns, nil
}

func GetCampaignsByThreatActor(threatActorID string) ([]*Campaign, error) {
	query := `
		SELECT id, name, description, first_seen, last_seen, objective, threat_actor_id, created_at, updated_at
		FROM campaigns
		WHERE threat_actor_id = $1
		ORDER BY first_seen DESC NULLS LAST
	`
	rows, err := database.DB.Query(query, threatActorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []*Campaign
	for rows.Next() {
		var cmp Campaign
		err := rows.Scan(
			&cmp.ID,
			&cmp.Name,
			&cmp.Description,
			&cmp.FirstSeen,
			&cmp.LastSeen,
			&cmp.Objective,
			&cmp.ThreatActorID,
			&cmp.CreatedAt,
			&cmp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, &cmp)
	}
	return campaigns, nil
}
