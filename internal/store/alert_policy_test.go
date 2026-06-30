package store

import (
	"testing"

	"github.com/elton451/argus/internal/model"
)

func TestCreateAlertPolicy(t *testing.T) {
	s := setupTestDB(t)

	p := &model.AlertPolicy{
		Name:             "PagerDuty Primary",
		ChannelType:      "pagerduty",
		ChannelConfig:    `{"routing_key":"abc123"}`,
		FailureThreshold: 5,
		CooldownMinutes:  30,
	}

	if err := s.CreateAlertPolicy(p); err != nil {
		t.Fatalf("CreateAlertPolicy: %v", err)
	}

	got, err := s.GetAlertPolicy(p.ID)
	if err != nil {
		t.Fatalf("GetAlertPolicy: %v", err)
	}

	if got.Name != p.Name {
		t.Errorf("name: got %q, want %q", got.Name, p.Name)
	}
	if got.ChannelType != p.ChannelType {
		t.Errorf("channel_type: got %q, want %q", got.ChannelType, p.ChannelType)
	}
	if got.ChannelConfig != p.ChannelConfig {
		t.Errorf("channel_config: got %q, want %q", got.ChannelConfig, p.ChannelConfig)
	}
	if got.FailureThreshold != p.FailureThreshold {
		t.Errorf("failure_threshold: got %d, want %d", got.FailureThreshold, p.FailureThreshold)
	}
	if got.CooldownMinutes != p.CooldownMinutes {
		t.Errorf("cooldown_minutes: got %d, want %d", got.CooldownMinutes, p.CooldownMinutes)
	}
	if got.CreatedAt == 0 {
		t.Error("created_at: expected non-zero timestamp")
	}
}

func TestListAlertPolicies(t *testing.T) {
	s := setupTestDB(t)

	for _, name := range []string{"Policy A", "Policy B"} {
		if err := s.CreateAlertPolicy(&model.AlertPolicy{
			Name:             name,
			ChannelType:      "slack",
			ChannelConfig:    `{}`,
			FailureThreshold: 3,
			CooldownMinutes:  15,
		}); err != nil {
			t.Fatalf("CreateAlertPolicy %s: %v", name, err)
		}
	}

	policies, err := s.ListAlertPolicies()
	if err != nil {
		t.Fatalf("ListAlertPolicies: %v", err)
	}

	if len(policies) != 2 {
		t.Fatalf("got %d policies, expected 2", len(policies))
	}
}

func TestDeleteAlertPolicy(t *testing.T) {
	s := setupTestDB(t)

	p := &model.AlertPolicy{
		Name:             "To Delete",
		ChannelType:      "email",
		ChannelConfig:    `{}`,
		FailureThreshold: 3,
		CooldownMinutes:  15,
	}
	if err := s.CreateAlertPolicy(p); err != nil {
		t.Fatalf("CreateAlertPolicy: %v", err)
	}

	if err := s.DeleteAlertPolicy(p.ID); err != nil {
		t.Fatalf("DeleteAlertPolicy: %v", err)
	}

	_, err := s.GetAlertPolicy(p.ID)
	if err == nil {
		t.Error("expected error after delete, got nil")
	}
}
