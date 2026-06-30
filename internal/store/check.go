package store

import (
	"time"

	"github.com/elton451/argus/internal/model"
)

func (s *Store) CreateCheck(c *model.Check) error {
	if c.CheckedAt == 0 {
		c.CheckedAt = time.Now().Unix()
	}

	result, err := s.db.Exec(`
		INSERT INTO checks (service_id, status_code, latency_ms, success, error, checked_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		c.ServiceID, c.StatusCode, c.LatencyMS,
		c.Success, c.Error, c.CheckedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.ID = id
	return nil
}

func (s *Store) GetCheck(id int64) (*model.Check, error) {
	row := s.db.QueryRow(`
		SELECT id, service_id, status_code, latency_ms, success, error, checked_at
		FROM checks WHERE id = ?
	`, id)

	c := &model.Check{}
	return c, row.Scan(
		&c.ID, &c.ServiceID, &c.StatusCode, &c.LatencyMS,
		&c.Success, &c.Error, &c.CheckedAt,
	)
}

func (s *Store) ListChecks() ([]*model.Check, error) {
	rows, err := s.db.Query(`
		SELECT id, service_id, status_code, latency_ms, success, error, checked_at
		FROM checks ORDER BY checked_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var checks []*model.Check
	for rows.Next() {
		c := &model.Check{}
		if err := rows.Scan(
			&c.ID, &c.ServiceID, &c.StatusCode, &c.LatencyMS,
			&c.Success, &c.Error, &c.CheckedAt,
		); err != nil {
			return nil, err
		}
		checks = append(checks, c)
	}

	return checks, rows.Err()
}

func (s *Store) DeleteCheck(id int64) error {
	_, err := s.db.Exec(`
		DELETE FROM checks WHERE id = ?
	`, id)

	return err
}
