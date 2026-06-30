package store

import (
	"testing"

	"github.com/elton451/argus/internal/model"
)

func TestCreateIncident(t *testing.T) {
	s := setupTestDB(t)

	svc := &model.Service{
		Name:            "Tested",
		URL:             "https://placeholder.com/api/v1",
		Method:          "GET",
		IntervalSeconds: 30,
		TimeoutSeconds:  15,
	}

	if err := s.CreateService(svc); err != nil {
		t.Fatalf("CreateService: %v", err)
	}

	incident := &model.Incident{
		State:        "CLOSED",
		Acknowledged: false,
		ServiceID:    svc.ID,
		StartedAt:    24,
	}

	if err := s.CreateIncident(incident); err != nil {
		t.Fatalf("Create incident: %v", err)
	}

	got, err := s.GetIncident(incident.ID)
	if err != nil {
		t.Fatalf("Get incident: %v", err)
	}

	if got.State != "CLOSED" {
		t.Fatalf("Wrong state on incident creation")
	}

	if got.Acknowledged != false {
		t.Fatalf("Wrong acknowledged flag on incident creation")
	}
}
