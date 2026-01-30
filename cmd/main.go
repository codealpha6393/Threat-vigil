package main

import (
	"github.com/gin-gonic/gin"
	"github.com/codealpha6393/threat-vigil/internal/adapter/handler/http"
	"github.com/codealpha6393/threat-vigil/internal/adapter/repository/memdb"
	"github.com/codealpha6393/threat-vigil/internal/core/services"
)

func main() {
	// 1. Initialize Repository (The Memory DB)
	repo := memdb.NewIncidentRepo()

	// 2. Initialize Service (The Logic)
	svc := services.NewIncidentService(repo)

	// 3. Initialize Handler (The API)
	handler := http.NewIncidentHandler(svc)

	// 4. Run Server
	router := gin.Default()
	handler.RegisterRoutes(router)
	
	router.Run(":8080") // Listen on port 8080
}