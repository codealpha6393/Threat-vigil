package ports

import "github.com/codealpha6393/threat-vigil/internal/core/domain"

// IncidentService defines the logic for handling security threats
type IncidentService interface {
	ReportIncident(incident domain.Incident) (domain.Incident, error)
	GetIncidents(severity string) ([]domain.Incident, error)
}

// IncidentRepository defines how we talk to the database
type IncidentRepository interface {
	Save(incident domain.Incident) (domain.Incident, error)
	List(severity string) ([]domain.Incident, error)
}