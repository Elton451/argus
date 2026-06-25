// internal/store/services_test.go
package store

import (
	"testing"

	"github.com/elton451/argus/internal/model"
)

func setupTestDB(t *testing.T) *Store {
	t.Helper()

	s, err := Open(":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}

	if err := s.Migrate("../migrations"); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	t.Cleanup(func() { s.db.Close() })

	return s
}

func TestCreateService(t *testing.T) {
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

	got, err := s.GetService(svc.ID)
	if err != nil {
		t.Fatalf("GetService: %v", err)
	}

	if got.Name != svc.Name {
		t.Errorf("name: got %q, want %q", got.Name, svc.Name)
	}
	if got.State != "UP" {
		t.Errorf("state: got %q, want %q", got.State, "UP")
	}
}

func TestListServices(t *testing.T) {
	s := setupTestDB(t)

	for _, name := range []string{"Service A", "Service B"} {
		if err := s.CreateService(&model.Service{
			Name:            name,
			URL:             "https://placehold.com/" + name,
			Method:          "GET",
			IntervalSeconds: 30,
			TimeoutSeconds:  15,
		}); err != nil {
			t.Fatalf("Create service %s", err)
		}
	}

	svcs, err := s.ListServices()
	if err != nil {
		t.Fatalf("List services %s", err)
	}

	if len(svcs) != 2 {
		t.Fatalf("Got %d services, expected 2", len(svcs))
	}
}

func TestDeleteService(t *testing.T) {
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

	err := s.DeleteService(svc.ID)
	if err != nil {
		t.Fatalf("Delete service %s", err)
	}

	_, error := s.GetService(svc.ID)
	if error == nil {
		t.Error("expected error after delete, got nil")
	}
}
