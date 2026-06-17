package services

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/opencti-lite/backend/models"
	"github.com/google/uuid"
)

type STIXObject map[string]interface{}

func GenerateSTIXBundle(iocs []*models.IOC, actors []*models.ThreatActor, campaigns []*models.Campaign, mappings []*models.ATTACKMapping) (map[string]interface{}, error) {
	var objects []STIXObject

	// 1. Convert IOCs to STIX Indicators
	for _, ioc := range iocs {
		pattern := ""
		switch ioc.Type {
		case "ip":
			if strings.Contains(ioc.Value, ":") {
				pattern = fmt.Sprintf("[ipv6-addr:value = '%s']", ioc.Value)
			} else {
				pattern = fmt.Sprintf("[ipv4-addr:value = '%s']", ioc.Value)
			}
		case "domain":
			pattern = fmt.Sprintf("[domain-name:value = '%s']", ioc.Value)
		case "url":
			pattern = fmt.Sprintf("[url:value = '%s']", ioc.Value)
		case "hash_md5":
			pattern = fmt.Sprintf("[file:hashes.'MD5' = '%s']", ioc.Value)
		case "hash_sha256":
			pattern = fmt.Sprintf("[file:hashes.'SHA-256' = '%s']", ioc.Value)
		case "email":
			pattern = fmt.Sprintf("[email-addr:value = '%s']", ioc.Value)
		default:
			pattern = fmt.Sprintf("[file:name = '%s']", ioc.Value)
		}

		stixIOC := STIXObject{
			"type":          "indicator",
			"spec_version":  "2.1",
			"id":            fmt.Sprintf("indicator--%s", ioc.ID),
			"created":       ioc.CreatedAt.Format(time.RFC3339),
			"modified":      ioc.UpdatedAt.Format(time.RFC3339),
			"indicator_types": []string{"malicious-activity"},
			"pattern":       pattern,
			"pattern_type":  "stix",
			"valid_from":    ioc.CreatedAt.Format(time.RFC3339),
			"name":          ioc.Value,
			"description":   ioc.Description,
			"labels":        ioc.Tags,
			"confidence":    ioc.Confidence,
		}

		// Set TLP marking definition in external references or custom field if needed
		stixIOC["x_opencti_tlp"] = ioc.TLPLevel
		stixIOC["x_opencti_source"] = ioc.Source

		objects = append(objects, stixIOC)
	}

	// 2. Convert Threat Actors
	for _, actor := range actors {
		stixActor := STIXObject{
			"type":                 "threat-actor",
			"spec_version":         "2.1",
			"id":                   fmt.Sprintf("threat-actor--%s", actor.ID),
			"created":              actor.CreatedAt.Format(time.RFC3339),
			"modified":             actor.UpdatedAt.Format(time.RFC3339),
			"name":                 actor.Name,
			"aliases":              actor.Aliases,
			"sophistication":       actor.Sophistication,
			"resource_level":       actor.ResourceLevel,
			"primary_motivation":   actor.PrimaryMotivation,
			"description":          actor.Description,
			"x_opencti_country":    actor.CountryCode,
		}
		objects = append(objects, stixActor)
	}

	// 3. Convert Campaigns + Relationships
	for _, camp := range campaigns {
		stixCamp := STIXObject{
			"type":         "campaign",
			"spec_version": "2.1",
			"id":           fmt.Sprintf("campaign--%s", camp.ID),
			"created":      camp.CreatedAt.Format(time.RFC3339),
			"modified":     camp.UpdatedAt.Format(time.RFC3339),
			"name":         camp.Name,
			"description":  camp.Description,
			"objective":    camp.Objective,
		}
		if camp.FirstSeen != nil {
			stixCamp["first_seen"] = camp.FirstSeen.Format(time.RFC3339)
		}
		if camp.LastSeen != nil {
			stixCamp["last_seen"] = camp.LastSeen.Format(time.RFC3339)
		}
		objects = append(objects, stixCamp)

		// Create relationship to threat actor if linked
		if camp.ThreatActorID != nil && *camp.ThreatActorID != "" {
			relID := uuid.New().String()
			stixRel := STIXObject{
				"type":              "relationship",
				"spec_version":      "2.1",
				"id":                fmt.Sprintf("relationship--%s", relID),
				"created":           camp.CreatedAt.Format(time.RFC3339),
				"modified":          camp.UpdatedAt.Format(time.RFC3339),
				"relationship_type": "attributed-to",
				"source_ref":        fmt.Sprintf("campaign--%s", camp.ID),
				"target_ref":        fmt.Sprintf("threat-actor--%s", *camp.ThreatActorID),
			}
			objects = append(objects, stixRel)
		}
	}

	// 4. Convert ATT&CK Mappings to Attack Patterns and Relationships
	for _, m := range mappings {
		apID := uuid.New().String()
		stixAP := STIXObject{
			"type":         "attack-pattern",
			"spec_version": "2.1",
			"id":           fmt.Sprintf("attack-pattern--%s", apID),
			"created":      m.CreatedAt.Format(time.RFC3339),
			"modified":     m.CreatedAt.Format(time.RFC3339),
			"name":         m.TechniqueName,
			"description":  fmt.Sprintf("MITRE ATT&CK Technique %s under %s tactic", m.TechniqueID, m.Tactic),
			"external_references": []map[string]string{
				{
					"source_name": "mitre-attack",
					"external_id": m.TechniqueID,
				},
			},
		}
		objects = append(objects, stixAP)

		// Create relationship between mapping target and the Attack Pattern
		relID := uuid.New().String()
		var targetType string
		switch m.EntityType {
		case "ioc":
			targetType = "indicator"
		case "threat_actor":
			targetType = "threat-actor"
		case "campaign":
			targetType = "campaign"
		}

		stixRel := STIXObject{
			"type":              "relationship",
			"spec_version":      "2.1",
			"id":                fmt.Sprintf("relationship--%s", relID),
			"created":           m.CreatedAt.Format(time.RFC3339),
			"modified":          m.CreatedAt.Format(time.RFC3339),
			"relationship_type": "uses",
			"source_ref":        fmt.Sprintf("%s--%s", targetType, m.EntityID),
			"target_ref":        fmt.Sprintf("attack-pattern--%s", apID),
		}
		objects = append(objects, stixRel)
	}

	bundleID := uuid.New().String()
	bundle := map[string]interface{}{
		"type":         "bundle",
		"id":           fmt.Sprintf("bundle--%s", bundleID),
		"spec_version": "2.1",
		"objects":      objects,
	}

	return bundle, nil
}

type STIXImportStats struct {
	IOCsInserted         int `json:"iocsInserted"`
	ThreatActorsInserted int `json:"threatActorsInserted"`
	CampaignsInserted    int `json:"campaignsInserted"`
	RelationshipsLinked  int `json:"relationshipsLinked"`
}

func ImportSTIXBundle(bundleBytes []byte) (*STIXImportStats, error) {
	var bundle map[string]interface{}
	err := json.Unmarshal(bundleBytes, &bundle)
	if err != nil {
		return nil, err
	}

	objectsVal, ok := bundle["objects"]
	if !ok {
		return nil, fmt.Errorf("invalid STIX bundle: missing objects field")
	}

	objectsList, ok := objectsVal.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid STIX bundle: objects field is not an array")
	}

	stats := &STIXImportStats{}
	actorsMap := make(map[string]string) // Map STIX ID to database ID

	// First pass: import Threat Actors and store IDs
	for _, objVal := range objectsList {
		obj, ok := objVal.(map[string]interface{})
		if !ok {
			continue
		}
		objType, _ := obj["type"].(string)

		if objType == "threat-actor" {
			stixID, _ := obj["id"].(string)
			name, _ := obj["name"].(string)
			desc, _ := obj["description"].(string)
			soph, _ := obj["sophistication"].(string)
			res, _ := obj["resource_level"].(string)
			mot, _ := obj["primary_motivation"].(string)
			country, _ := obj["x_opencti_country"].(string)

			// Clean sophistication to match db check
			if soph != "minimal" && soph != "intermediate" && soph != "advanced" && soph != "expert" {
				soph = "intermediate" // Default fallback
			}

			var aliases []string
			if aliasVals, ok := obj["aliases"].([]interface{}); ok {
				for _, av := range aliasVals {
					if astr, ok := av.(string); ok {
						aliases = append(aliases, astr)
					}
				}
			}

			ta := models.ThreatActor{
				Name:              name,
				Aliases:           aliases,
				Sophistication:    soph,
				ResourceLevel:     res,
				PrimaryMotivation: mot,
				CountryCode:       country,
				Description:       desc,
			}

			err := ta.Create()
			if err == nil {
				stats.ThreatActorsInserted++
				actorsMap[stixID] = ta.ID
			}
		}
	}

	// Second pass: import Campaigns and IOCs
	campaignsMap := make(map[string]string)
	for _, objVal := range objectsList {
		obj, ok := objVal.(map[string]interface{})
		if !ok {
			continue
		}
		objType, _ := obj["type"].(string)

		if objType == "campaign" {
			stixID, _ := obj["id"].(string)
			name, _ := obj["name"].(string)
			desc, _ := obj["description"].(string)
			objec, _ := obj["objective"].(string)

			var firstSeen, lastSeen *time.Time
			if fsStr, ok := obj["first_seen"].(string); ok {
				if t, err := time.Parse(time.RFC3339, fsStr); err == nil {
					firstSeen = &t
				}
			}
			if lsStr, ok := obj["last_seen"].(string); ok {
				if t, err := time.Parse(time.RFC3339, lsStr); err == nil {
					lastSeen = &t
				}
			}

			cmp := models.Campaign{
				Name:        name,
				Description: desc,
				FirstSeen:   firstSeen,
				LastSeen:    lastSeen,
				Objective:   objec,
			}

			err := cmp.Create()
			if err == nil {
				stats.CampaignsInserted++
				campaignsMap[stixID] = cmp.ID
			}
		} else if objType == "indicator" {
			name, _ := obj["name"].(string)
			desc, _ := obj["description"].(string)
			confVal, _ := obj["confidence"].(float64)
			pattern, _ := obj["pattern"].(string)
			tlp, _ := obj["x_opencti_tlp"].(string)
			source, _ := obj["x_opencti_source"].(string)

			var tags []string
			if labelVals, ok := obj["labels"].([]interface{}); ok {
				for _, lv := range labelVals {
					if lstr, ok := lv.(string); ok {
						tags = append(tags, lstr)
					}
				}
			}

			// Parse type and value from pattern
			// Examples: [ipv4-addr:value = '1.2.3.4'] or [domain-name:value = 'evil.com']
			iocType := "domain"
			iocValue := name
			if strings.Contains(pattern, "ipv4-addr:value") {
				iocType = "ip"
			} else if strings.Contains(pattern, "ipv6-addr:value") {
				iocType = "ip"
			} else if strings.Contains(pattern, "url:value") {
				iocType = "url"
			} else if strings.Contains(pattern, "hashes.'MD5'") {
				iocType = "hash_md5"
			} else if strings.Contains(pattern, "hashes.'SHA-256'") {
				iocType = "hash_sha256"
			} else if strings.Contains(pattern, "email-addr:value") {
				iocType = "email"
			}

			// Extract value inside single quotes if name is empty
			if iocValue == "" {
				parts := strings.Split(pattern, "'")
				if len(parts) >= 3 {
					iocValue = parts[1]
				} else {
					iocValue = "unknown"
				}
			}

			// Clean TLP
			if tlp != "white" && tlp != "green" && tlp != "amber" && tlp != "red" {
				tlp = "white"
			}

			ioc := models.IOC{
				Type:        iocType,
				Value:       iocValue,
				TLPLevel:    tlp,
				Confidence:  int(confVal),
				Tags:        tags,
				Source:      source,
				Description: desc,
			}
			if ioc.Confidence == 0 {
				ioc.Confidence = 50
			}

			err := ioc.Create()
			if err == nil {
				stats.IOCsInserted++
			}
		}
	}

	// Third pass: Link campaigns and actors if relationship exists
	for _, objVal := range objectsList {
		obj, ok := objVal.(map[string]interface{})
		if !ok {
			continue
		}
		objType, _ := obj["type"].(string)

		if objType == "relationship" {
			relType, _ := obj["relationship_type"].(string)
			sourceRef, _ := obj["source_ref"].(string)
			targetRef, _ := obj["target_ref"].(string)

			if relType == "attributed-to" && strings.HasPrefix(sourceRef, "campaign--") && strings.HasPrefix(targetRef, "threat-actor--") {
				campID := campaignsMap[sourceRef]
				actorID := actorsMap[targetRef]

				if campID != "" && actorID != "" {
					campaign, err := models.GetCampaignByID(campID)
					if err == nil {
						campaign.ThreatActorID = &actorID
						err = campaign.Update()
						if err == nil {
							stats.RelationshipsLinked++
						}
					}
				}
			}
		}
	}

	return stats, nil
}
