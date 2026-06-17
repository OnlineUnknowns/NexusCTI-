package handlers

import (
	"net/http"

	"github.com/opencti-lite/backend/models"
	"github.com/gin-gonic/gin"
)

func ListThreatActorsHandler(c *gin.Context) {
	actors, err := models.ListThreatActors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actors)
}

func CreateThreatActorHandler(c *gin.Context) {
	var actor models.ThreatActor
	if err := c.ShouldBindJSON(&actor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if actor.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if err := actor.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, actor)
}

func GetThreatActorHandler(c *gin.Context) {
	id := c.Param("id")
	actor, err := models.GetThreatActorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Threat actor not found"})
		return
	}
	c.JSON(http.StatusOK, actor)
}

func UpdateThreatActorHandler(c *gin.Context) {
	id := c.Param("id")
	actor, err := models.GetThreatActorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Threat actor not found"})
		return
	}

	if err := c.ShouldBindJSON(&actor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := actor.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actor)
}

func DeleteThreatActorHandler(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteThreatActor(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Threat actor deleted successfully"})
}
