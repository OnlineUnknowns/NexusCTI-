package main

import (
	"log"

	"github.com/opencti-lite/backend/config"
	"github.com/opencti-lite/backend/database"
	"github.com/opencti-lite/backend/handlers"
	"github.com/opencti-lite/backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load configuration
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading from environment variables directly")
	}

	config.LoadConfig()

	// 2. Initialize Database & Redis
	database.InitDB()

	// 3. Setup router
	r := gin.New()

	// 4. Register global middlewares
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Recovery())

	// 5. Auth public endpoints
	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", handlers.RegisterHandler)
		auth.POST("/login", handlers.LoginHandler)
	}

	// 6. Protected endpoints (requires token & rate limiting)
	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	api.Use(middleware.RateLimitMiddleware())
	{
		// Dashboard
		api.GET("/dashboard/stats", handlers.DashboardStatsHandler)

		// IOCs
		api.GET("/iocs", handlers.ListIOCsHandler)
		api.POST("/iocs", handlers.CreateIOCHandler)
		api.GET("/iocs/:id", handlers.GetIOCHandler)
		api.PUT("/iocs/:id", handlers.UpdateIOCHandler)
		api.DELETE("/iocs/:id", handlers.DeleteIOCHandler)
		api.POST("/iocs/bulk", handlers.BulkCreateIOCHandler)

		// Threat Actors
		api.GET("/threat-actors", handlers.ListThreatActorsHandler)
		api.POST("/threat-actors", handlers.CreateThreatActorHandler)
		api.GET("/threat-actors/:id", handlers.GetThreatActorHandler)
		api.PUT("/threat-actors/:id", handlers.UpdateThreatActorHandler)
		api.DELETE("/threat-actors/:id", handlers.DeleteThreatActorHandler)

		// Campaigns
		api.GET("/campaigns", handlers.ListCampaignsHandler)
		api.POST("/campaigns", handlers.CreateCampaignHandler)
		api.GET("/campaigns/:id", handlers.GetCampaignHandler)
		api.PUT("/campaigns/:id", handlers.UpdateCampaignHandler)
		api.DELETE("/campaigns/:id", handlers.DeleteCampaignHandler)

		// MITRE ATT&CK Mappings
		api.GET("/attack/mappings", handlers.ListATTACKMappingsHandler)
		api.POST("/attack/mappings", handlers.CreateATTACKMappingHandler)
		api.GET("/attack/techniques", handlers.GetGroupedByTacticHandler)
		api.DELETE("/attack/mappings/:id", handlers.DeleteATTACKMappingHandler)

		// STIX Import/Export
		api.POST("/stix/export", handlers.STIXExportHandler)
		api.POST("/stix/import", handlers.STIXImportHandler)
	}

	// 7. TAXII endpoints (public access for standard integration)
	taxii := r.Group("/taxii/v21")
	{
		taxii.GET("/collections", handlers.TAXIICollectionsHandler)
		taxii.GET("/collections/:id/objects", handlers.TAXIIObjectsHandler)
	}

	// 8. Start server
	log.Printf("Starting OpenCTI Clone Lite backend server on :%s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
