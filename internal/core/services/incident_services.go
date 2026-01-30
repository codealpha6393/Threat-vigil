package services

import (
	"time"

	"github.com/codealpha6393/threat-vigil/internal/core/domain"
	"github.com/codealpha6393/threat-vigil/internal/core/ports"
	"github.com/google/uuid"
)

type IncidentService struct {
	repo ports.IncidentRepository
}

func NewIncidentService(repo ports.IncidentRepository) *IncidentService {
	return &IncidentService{repo: repo}
}

func (s *IncidentService) ReportIncident(incident domain.Incident) (domain.Incident, error) {
	// Simple Logic: If no severity is provided, default to LOW
	if incident.Severity == "" {
		incident.Severity = "LOW"
	}

	// Generate ID and Timestamp
	incident.ID = uuid.New().String()
	incident.CreatedAt = time.Now()
	incident.Status = "OPEN"

	return s.repo.Save(incident)
}

func (s *IncidentService) GetIncidents(severity string) ([]domain.Incident, error) {
	return s.repo.List(severity)
}
