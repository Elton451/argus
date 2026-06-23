-- +goose Up
CREATE TABLE alert_policies (
  id               INTEGER PRIMARY KEY,
  name             TEXT    NOT NULL,
  channel_type     TEXT    NOT NULL,
  channel_config   TEXT    NOT NULL DEFAULT '{}',
  failure_threshold INTEGER NOT NULL DEFAULT 3,
  cooldown_minutes  INTEGER NOT NULL DEFAULT 15,
  created_at        INTEGER NOT NULL DEFAULT (unixepoch())
);

CREATE TABLE services ( 
	id	INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	url TEXT NOT NULL,
	method TEXT NOT NULL DEFAULT "GET",
	interval_seconds INTEGER NOT NULL DEFAULT 60,
	timeout_seconds INTEGER NOT NULL DEFAULT 10,
	state TEXT NOT NULL DEFAULT "UP",
	failure_streak INTEGER DEFAULT 0,
	alert_policy_id INTEGER REFERENCES alert_policies(id) ON DELETE SET NULL,
  created_at        INTEGER NOT NULL DEFAULT (unixepoch()),
	updated_at INTEGER NOT NULL DEFAULT(unixepoch())
);

CREATE TABLE checks ( 
	id INTEGER PRIMARY KEY,
	service_id INTEGER REFERENCES services(id) ON DELETE CASCADE,
	status_code INTEGER,
	latency_ms INTEGER NOT NULL,
	success INTEGER NOT NULL DEFAULT 0,
	error TEXT,
	checked_at  INTEGER NOT NULL DEFAULT (unixepoch())
);

CREATE INDEX idx_checks_service_id_checked_at
  ON checks(service_id, checked_at DESC);

CREATE TABLE incidents (
	id INTEGER PRIMARY KEY,
	service_id INTEGER NOT NULL REFERENCES services(id) ON DELETE CASCADE,
	state TEXT NOT NULL DEFAULT 'OPEN',
	trigger_check_id INTEGER REFERENCES checks(id) ON DELETE SET NULL,
	summary_ai TEXT,
	started_at INTEGER NOT NULL DEFAULT (unixepoch()),
	resolved_at INTEGER,
	acknowledged INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX idx_incidents_service_id
  ON incidents(service_id, started_at DESC);

CREATE INDEX idx_incidents_open
  ON incidents(state)
  WHERE state = 'OPEN';

-- +goose Down
DROP TABLE IF EXISTS incidents;
DROP TABLE IF EXISTS checks;
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS alert_policies;
