package handlers

import (
	"net/http"

	"github.com/opencti-lite/backend/models"
	"github.com/gin-gonic/gin"
)

func ListATTACKMappingsHandler(c *gin.Context) {
	entityType := c.Query("entity_type")
	entityID := c.Query("entity_id")
	tactic := c.Query("tactic")

	filters := map[string]interface{}{
		"entity_type": entityType,
		"entity_id":   entityID,
		"tactic":      tactic,
	}

	mappings, err := models.ListATTACKMappings(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mappings)
}

func CreateATTACKMappingHandler(c *gin.Context) {
	var mapping models.ATTACKMapping
	if err := c.ShouldBindJSON(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if mapping.TechniqueID == "" || mapping.TechniqueName == "" || mapping.Tactic == "" || mapping.EntityType == "" || mapping.EntityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "techniqueId, techniqueName, tactic, entityType, and entityId are required"})
		return
	}

	if err := mapping.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mapping)
}

func GetGroupedByTacticHandler(c *gin.Context) {
	grouped, err := models.GetGroupedByTactic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, grouped)
}

func DeleteATTACKMappingHandler(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteATTACKMapping(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ATT&CK mapping deleted successfully"})
}
