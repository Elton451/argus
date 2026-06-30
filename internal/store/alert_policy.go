package store

import "github.com/elton451/argus/internal/model"

func (s *Store) CreateAlertPolicy(p *model.AlertPolicy) error {
	result, err := s.db.Exec(`
		INSERT INTO alert_policies (name, channel_type, channel_config, failure_threshold, cooldown_minutes)
		VALUES (?, ?, ?, ?, ?)
	`,
		p.Name, p.ChannelType, p.ChannelConfig,
		p.FailureThreshold, p.CooldownMinutes,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = id
	return nil
}

func (s *Store) GetAlertPolicy(id int64) (*model.AlertPolicy, error) {
	row := s.db.QueryRow(`
		SELECT id, name, channel_type, channel_config, failure_threshold, cooldown_minutes, created_at
		FROM alert_policies WHERE id = ?
	`, id)

	p := &model.AlertPolicy{}
	return p, row.Scan(
		&p.ID, &p.Name, &p.ChannelType, &p.ChannelConfig,
		&p.FailureThreshold, &p.CooldownMinutes, &p.CreatedAt,
	)
}

func (s *Store) ListAlertPolicies() ([]*model.AlertPolicy, error) {
	rows, err := s.db.Query(`
		SELECT id, name, channel_type, channel_config, failure_threshold, cooldown_minutes, created_at
		FROM alert_policies ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var policies []*model.AlertPolicy
	for rows.Next() {
		p := &model.AlertPolicy{}
		if err := rows.Scan(
			&p.ID, &p.Name, &p.ChannelType, &p.ChannelConfig,
			&p.FailureThreshold, &p.CooldownMinutes, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		policies = append(policies, p)
	}

	return policies, rows.Err()
}

func (s *Store) DeleteAlertPolicy(id int64) error {
	_, err := s.db.Exec(`
		DELETE FROM alert_policies WHERE id = ?
	`, id)

	return err
}
