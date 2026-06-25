package store

import "github.com/elton451/argus/internal/model"

func (s *Store) CreateService(svc *model.Service) error {
	result, err := s.db.Exec(`
			INSERT INTO services (name, url, method, interval_seconds, timeout_seconds, alert_policy_id)
			VALUES (?, ?, ?, ?, ?, ?)
		`,
		svc.Name, svc.URL, svc.Method,
		svc.IntervalSeconds, svc.TimeoutSeconds,
		svc.AlertPolicyID,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	svc.ID = id
	return nil
}

func (s *Store) GetService(id int64) (*model.Service, error) {
	row := s.db.QueryRow(`
		SELECT id, name, url, method, interval_seconds, timeout_seconds,
    state, failure_streak, alert_policy_id, created_at, updated_at
  	FROM services WHERE id = ?
	`, id)

	svc := &model.Service{}
	return svc, row.Scan(
		&svc.ID, &svc.Name, &svc.URL, &svc.Method,
		&svc.IntervalSeconds, &svc.TimeoutSeconds,
		&svc.State, &svc.FailureStreak, &svc.AlertPolicyID,
		&svc.CreatedAt, &svc.UpdatedAt,
	)
}

func (s *Store) ListServices() ([]*model.Service, error) {
	rows, err := s.db.Query(`
	  SELECT id, name, url, method, interval_seconds, timeout_seconds,
    state, failure_streak, alert_policy_id, created_at, updated_at
    FROM services ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var serviceList []*model.Service
	for rows.Next() {
		svc := &model.Service{}
		if err := rows.Scan(
			&svc.ID, &svc.Name, &svc.URL, &svc.Method,
			&svc.IntervalSeconds, &svc.TimeoutSeconds,
			&svc.State, &svc.FailureStreak, &svc.AlertPolicyID,
			&svc.CreatedAt, &svc.UpdatedAt,
		); err != nil {
			return nil, err
		}
		serviceList = append(serviceList, svc)
	}

	return serviceList, rows.Err()
}

func (s *Store) DeleteService(id int64) error {
	_, err := s.db.Exec(`
		DELETE FROM services WHERE id = ?
	`, id)

	return err
}
