package http

import (
	"github.com/codealpha6393/threat-vigil/internal/core/domain"
	"github.com/codealpha6393/threat-vigil/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type IncidentHandler struct {
	svc ports.IncidentService
}

func NewIncidentHandler(svc ports.IncidentService) *IncidentHandler {
	return &IncidentHandler{svc: svc}
}

func (h *IncidentHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/incidents", h.ReportIncident)
	router.GET("/incidents", h.GetIncidents)
}

func (h *IncidentHandler) ReportIncident(c *gin.Context) {
	var incident domain.Incident
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	created, _ := h.svc.ReportIncident(incident)
	c.JSON(201, created)
}

func (h *IncidentHandler) GetIncidents(c *gin.Context) {
	severity := c.Query("severity")
	incidents, _ := h.svc.GetIncidents(severity)
	c.JSON(200, incidents)
}
