package store

import (
	"time"

	"github.com/elton451/argus/internal/model"
)

func (s *Store) CreateIncident(incident *model.Incident) error {
    if incident.StartedAt == 0 {
        incident.StartedAt = time.Now().Unix()
    }

    result, err := s.db.Exec(`
        INSERT INTO incidents (service_id, state, trigger_check_id, summary_ai, started_at, resolved_at, acknowledged)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `, incident.ServiceID, incident.State, incident.TriggerCheckID,
       incident.SummaryAI, incident.StartedAt, incident.ResolvedAt, incident.Acknowledged)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }

    incident.ID = id
    return nil
}


func (s *Store) ListIncidents() ([]*model.Incident, error) {
	rows, err := s.db.Query(`
		SELECT * FROM incidents
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []*model.Incident
	for rows.Next() {
		incident := &model.Incident{}
		if err := rows.Scan(&incident.ID,
			&incident.Acknowledged,
			&incident.ResolvedAt,
			&incident.ServiceID,
			&incident.StartedAt,
			&incident.State,
			&incident.SummaryAI,
			&incident.TriggerCheckID); err != nil {
			return nil, err
		}
		incidents = append(incidents, incident)
	}

	return incidents, rows.Err()
}

func (s *Store) GetIncident(id int64) (*model.Incident, error) {
	row := s.db.QueryRow(`
		SELECT * FROM incidents WHERE id = ?
	`, id)

	incident := &model.Incident{}

	return incident, row.Scan(
		&incident.ID,
		&incident.ServiceID,
		&incident.State,
		&incident.TriggerCheckID,
		&incident.SummaryAI,
		&incident.StartedAt,
		&incident.ResolvedAt,
		&incident.Acknowledged,
	)
}

func (s *Store) DeleteIncident(id int64) error {
	_, err := s.db.Exec(`
		DELETE FROM incidents WHERE id = ?
	`, id)

	return err
}
