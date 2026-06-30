package store

import (
	"testing"

	"github.com/elton451/argus/internal/model"
)

func createTestService(t *testing.T, s *Store) *model.Service {
	t.Helper()

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
	return svc
}

func TestCreateCheck(t *testing.T) {
	s := setupTestDB(t)
	svc := createTestService(t, s)

	c := &model.Check{
		ServiceID:  svc.ID,
		StatusCode: 200,
		LatencyMS:  42,
		Success:    1,
		CheckedAt:  1700000000,
	}

	if err := s.CreateCheck(c); err != nil {
		t.Fatalf("CreateCheck: %v", err)
	}

	got, err := s.GetCheck(c.ID)
	if err != nil {
		t.Fatalf("GetCheck: %v", err)
	}

	if got.ServiceID != c.ServiceID {
		t.Errorf("service_id: got %d, want %d", got.ServiceID, c.ServiceID)
	}
	if got.StatusCode != c.StatusCode {
		t.Errorf("status_code: got %d, want %d", got.StatusCode, c.StatusCode)
	}
	if got.LatencyMS != c.LatencyMS {
		t.Errorf("latency_ms: got %d, want %d", got.LatencyMS, c.LatencyMS)
	}
	if got.Success != c.Success {
		t.Errorf("success: got %d, want %d", got.Success, c.Success)
	}
	if got.CheckedAt != c.CheckedAt {
		t.Errorf("checked_at: got %d, want %d", got.CheckedAt, c.CheckedAt)
	}
}

func TestListChecks(t *testing.T) {
	s := setupTestDB(t)
	svc := createTestService(t, s)

	for i, latency := range []int{10, 20} {
		if err := s.CreateCheck(&model.Check{
			ServiceID:  svc.ID,
			StatusCode: 200,
			LatencyMS:  latency,
			Success:    1,
			CheckedAt:  int64(1700000000 + i),
		}); err != nil {
			t.Fatalf("CreateCheck: %v", err)
		}
	}

	checks, err := s.ListChecks()
	if err != nil {
		t.Fatalf("ListChecks: %v", err)
	}

	if len(checks) != 2 {
		t.Fatalf("got %d checks, expected 2", len(checks))
	}
}

func TestDeleteCheck(t *testing.T) {
	s := setupTestDB(t)
	svc := createTestService(t, s)

	c := &model.Check{
		ServiceID:  svc.ID,
		StatusCode: 500,
		LatencyMS:  100,
		Success:    0,
		Error:      "connection refused",
		CheckedAt:  1700000000,
	}
	if err := s.CreateCheck(c); err != nil {
		t.Fatalf("CreateCheck: %v", err)
	}

	if err := s.DeleteCheck(c.ID); err != nil {
		t.Fatalf("DeleteCheck: %v", err)
	}

	_, err := s.GetCheck(c.ID)
	if err == nil {
		t.Error("expected error after delete, got nil")
	}
}
