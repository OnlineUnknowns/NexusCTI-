package handlers

import (
	"database/sql"
	"net/http"

	"github.com/opencti-lite/backend/database"
	"github.com/opencti-lite/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type DashboardStats struct {
	TotalIOCs              int                      `json:"totalIOCs"`
	TotalThreatActors      int                      `json:"totalThreatActors"`
	TotalCampaigns         int                      `json:"totalCampaigns"`
	TotalAttackMappings    int                      `json:"totalAttackMappings"`
	ConfidenceDistribution []ConfidenceBucket       `json:"confidenceDistribution"`
	TypeBreakdown          []TypeBreakdownItem      `json:"typeBreakdown"`
	RecentIOCs             []*models.IOC            `json:"recentIOCs"`
}

type ConfidenceBucket struct {
	Bucket string `json:"bucket"`
	Count  int    `json:"count"`
}

type TypeBreakdownItem struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

func DashboardStatsHandler(c *gin.Context) {
	var stats DashboardStats

	// 1. Get totals
	err := database.DB.QueryRow("SELECT COUNT(*) FROM iocs").Scan(&stats.TotalIOCs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.QueryRow("SELECT COUNT(*) FROM threat_actors").Scan(&stats.TotalThreatActors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.QueryRow("SELECT COUNT(*) FROM campaigns").Scan(&stats.TotalCampaigns)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.QueryRow("SELECT COUNT(*) FROM attack_mappings").Scan(&stats.TotalAttackMappings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2. Confidence Distribution
	buckets := []struct {
		Label string
		Min   int
		Max   int
	}{
		{"0-25", 0, 25},
		{"26-50", 26, 50},
		{"51-75", 51, 75},
		{"76-100", 76, 100},
	}

	stats.ConfidenceDistribution = make([]ConfidenceBucket, 0)
	for _, b := range buckets {
		var count int
		err = database.DB.QueryRow("SELECT COUNT(*) FROM iocs WHERE confidence >= $1 AND confidence <= $2", b.Min, b.Max).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stats.ConfidenceDistribution = append(stats.ConfidenceDistribution, ConfidenceBucket{
			Bucket: b.Label,
			Count:  count,
		})
	}

	// 3. Type Breakdown
	stats.TypeBreakdown = make([]TypeBreakdownItem, 0)
	rows, err := database.DB.Query("SELECT type, COUNT(*) FROM iocs GROUP BY type")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item TypeBreakdownItem
		if err := rows.Scan(&item.Type, &item.Count); err == nil {
			stats.TypeBreakdown = append(stats.TypeBreakdown, item)
		}
	}

	// 4. Recent IOCs (last 10)
	stats.RecentIOCs = make([]*models.IOC, 0)
	recentRows, err := database.DB.Query(`
		SELECT id, type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen, created_at, updated_at
		FROM iocs
		ORDER BY created_at DESC
		LIMIT 10
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer recentRows.Close()

	for recentRows.Next() {
		var i models.IOC
		err := recentRows.Scan(
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
		if err == nil {
			stats.RecentIOCs = append(stats.RecentIOCs, &i)
		}
	}

	c.JSON(http.StatusOK, stats)
}
