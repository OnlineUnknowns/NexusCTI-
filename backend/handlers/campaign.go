package handlers

import (
	"net/http"

	"github.com/opencti-lite/backend/models"
	"github.com/gin-gonic/gin"
)

func ListCampaignsHandler(c *gin.Context) {
	actorID := c.Query("threat_actor_id")
	var campaigns []*models.Campaign
	var err error

	if actorID != "" {
		campaigns, err = models.GetCampaignsByThreatActor(actorID)
	} else {
		campaigns, err = models.ListCampaigns()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func CreateCampaignHandler(c *gin.Context) {
	var campaign models.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if campaign.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	// Clean threat actor empty string
	if campaign.ThreatActorID != nil && *campaign.ThreatActorID == "" {
		campaign.ThreatActorID = nil
	}

	if err := campaign.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, campaign)
}

func GetCampaignHandler(c *gin.Context) {
	id := c.Param("id")
	campaign, err := models.GetCampaignByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func UpdateCampaignHandler(c *gin.Context) {
	id := c.Param("id")
	campaign, err := models.GetCampaignByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Clean threat actor empty string
	if campaign.ThreatActorID != nil && *campaign.ThreatActorID == "" {
		campaign.ThreatActorID = nil
	}

	if err := campaign.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func DeleteCampaignHandler(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteCampaign(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
