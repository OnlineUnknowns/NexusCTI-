package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/opencti-lite/backend/models"
	"github.com/opencti-lite/backend/services"
	"github.com/gin-gonic/gin"
)

func STIXExportHandler(c *gin.Context) {
	// Fetch all IOCs
	iocs, _, err := models.ListIOCs(map[string]interface{}{}, 1, 1000)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load IOCs: " + err.Error()})
		return
	}

	// Fetch all actors
	actors, err := models.ListThreatActors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load Threat Actors: " + err.Error()})
		return
	}

	// Fetch all campaigns
	campaigns, err := models.ListCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load Campaigns: " + err.Error()})
		return
	}

	// Fetch all mappings
	mappings, err := models.ListATTACKMappings(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load ATT&CK mappings: " + err.Error()})
		return
	}

	// Generate STIX 2.1 Bundle
	bundle, err := services.GenerateSTIXBundle(iocs, actors, campaigns, mappings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate STIX bundle: " + err.Error()})
		return
	}

	bundleJSON, err := json.Marshal(bundle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal STIX bundle: " + err.Error()})
		return
	}

	// Save to database
	stixBundle := models.STIXBundle{
		SpecVersion: "2.1",
		BundleJSON:  json.RawMessage(bundleJSON),
	}

	if err := stixBundle.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save STIX bundle to db: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bundle)
}

func STIXImportHandler(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body: " + err.Error()})
		return
	}

	stats, err := services.ImportSTIXBundle(bodyBytes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse STIX bundle: " + err.Error()})
		return
	}

	// Invalidate caches since we imported new records
	services.InvalidatePattern("opencti:iocs:*")

	c.JSON(http.StatusOK, gin.H{
		"message": "STIX bundle imported successfully",
		"stats":   stats,
	})
}

func TAXIICollectionsHandler(c *gin.Context) {
	collections := gin.H{
		"collections": []gin.H{
			{
				"id":          "91a7b524-c2cf-434b-93f0-1216bca238d8",
				"title":       "OpenCTI Lite Default Collection",
				"description": "Standard intelligence collection containing threat actors, campaigns, and indicators.",
				"can_read":    true,
				"can_write":   false,
				"media_types": []string{
					"application/stix+json;version=2.1",
				},
			},
		},
	}
	c.Header("Content-Type", "application/taxii+json;version=2.1")
	c.JSON(http.StatusOK, collections)
}

func TAXIIObjectsHandler(c *gin.Context) {
	// Fetch all data to construct latest intelligence package
	iocs, _, err := models.ListIOCs(map[string]interface{}{}, 1, 1000)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	actors, err := models.ListThreatActors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	campaigns, err := models.ListCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mappings, err := models.ListATTACKMappings(map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bundle, err := services.GenerateSTIXBundle(iocs, actors, campaigns, mappings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/stix+json;version=2.1")
	c.JSON(http.StatusOK, bundle)
}
