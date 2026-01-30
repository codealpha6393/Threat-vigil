package memdb

import (
	"github.com/codealpha6393/threat-vigil/internal/core/domain"
)

type IncidentRepo struct {
	db []domain.Incident
}

func NewIncidentRepo() *IncidentRepo {
	return &IncidentRepo{
		db: []domain.Incident{},
	}
}

func (r *IncidentRepo) Save(incident domain.Incident) (domain.Incident, error) {
	r.db = append(r.db, incident)
	return incident, nil
}

func (r *IncidentRepo) List(severity string) ([]domain.Incident, error) {
	var filtered []domain.Incident
	for _, inc := range r.db {
		if severity == "" || inc.Severity == severity {
			filtered = append(filtered, inc)
		}
	}
	return filtered, nil
}