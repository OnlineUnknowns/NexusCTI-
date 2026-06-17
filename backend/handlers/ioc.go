package handlers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/opencti-lite/backend/models"
	"github.com/opencti-lite/backend/services"
	"github.com/gin-gonic/gin"
)

func generateIOCListCacheKey(filters map[string]interface{}, page, limit int) string {
	raw := fmt.Sprintf("type:%v;tlp:%v;tag:%v;q:%v;p:%d;l:%d",
		filters["type"], filters["tlp"], filters["tag"], filters["q"], page, limit)
	hash := md5.Sum([]byte(raw))
	return fmt.Sprintf("opencti:iocs:list:%x", hash)
}

func ListIOCsHandler(c *gin.Context) {
	iocType := c.Query("type")
	tlp := c.Query("tlp")
	tag := c.Query("tag")
	q := c.Query("q")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 20
	}

	filters := map[string]interface{}{
		"type": iocType,
		"tlp":  tlp,
		"tag":  tag,
		"q":    q,
	}

	cacheKey := generateIOCListCacheKey(filters, page, limit)

	// Try to get from Redis cache
	if cachedVal, err := services.Get(cacheKey); err == nil && cachedVal != "" {
		var cachedResponse gin.H
		if err := json.Unmarshal([]byte(cachedVal), &cachedResponse); err == nil {
			c.JSON(http.StatusOK, cachedResponse)
			return
		}
	}

	// Fetch from DB
	iocs, total, err := models.ListIOCs(filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"data":  iocs,
		"total": total,
		"page":  page,
		"limit": limit,
	}

	// Save to Redis cache for 5 minutes
	services.Set(cacheKey, response, 5*time.Minute)

	c.JSON(http.StatusOK, response)
}

func CreateIOCHandler(c *gin.Context) {
	var ioc models.IOC
	if err := c.ShouldBindJSON(&ioc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	if ioc.Type == "" || ioc.Value == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "type and value are required"})
		return
	}

	if ioc.TLPLevel == "" {
		ioc.TLPLevel = "white"
	}

	if ioc.Confidence == 0 {
		ioc.Confidence = 50
	}

	if err := ioc.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Invalidate Cache
	services.InvalidatePattern("opencti:iocs:*")

	c.JSON(http.StatusCreated, ioc)
}

func GetIOCHandler(c *gin.Context) {
	id := c.Param("id")
	cacheKey := fmt.Sprintf("opencti:iocs:detail:%s", id)

	// Try Cache
	if cachedVal, err := services.Get(cacheKey); err == nil && cachedVal != "" {
		var ioc models.IOC
		if err := json.Unmarshal([]byte(cachedVal), &ioc); err == nil {
			c.JSON(http.StatusOK, ioc)
			return
		}
	}

	ioc, err := models.GetIOCByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "IOC not found"})
		return
	}

	services.Set(cacheKey, ioc, 10*time.Minute)

	c.JSON(http.StatusOK, ioc)
}

func UpdateIOCHandler(c *gin.Context) {
	id := c.Param("id")
	ioc, err := models.GetIOCByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "IOC not found"})
		return
	}

	if err := c.ShouldBindJSON(&ioc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ioc.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Invalidate cache entries
	services.Delete(fmt.Sprintf("opencti:iocs:detail:%s", id))
	services.InvalidatePattern("opencti:iocs:*")

	c.JSON(http.StatusOK, ioc)
}

func DeleteIOCHandler(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteIOC(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Invalidate Cache
	services.Delete(fmt.Sprintf("opencti:iocs:detail:%s", id))
	services.InvalidatePattern("opencti:iocs:*")

	c.JSON(http.StatusOK, gin.H{"message": "IOC deleted successfully"})
}

func BulkCreateIOCHandler(c *gin.Context) {
	var req struct {
		IOCs []models.IOC `json:"iocs" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count, err := models.BulkCreateIOCs(req.IOCs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "count": count})
		return
	}

	services.InvalidatePattern("opencti:iocs:*")

	c.JSON(http.StatusOK, gin.H{"message": "Bulk import completed", "inserted": count})
}
