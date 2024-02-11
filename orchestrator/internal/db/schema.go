package db

var schema = `
CREATE TABLE IF NOT EXISTS agents
(
    id        UUID PRIMARY KEY,
    last_ping TIMESTAMP,
    ip       TEXT NOT NULL,
    port     INTEGER,
    status    TEXT NOT NULL DEFAULT 'active',
    config_is_up_to_date BOOLEAN DEFAULT TRUE
);
CREATE TABLE IF NOT EXISTS expressions
(
    id             UUID PRIMARY KEY,
    REQUEST_ID     TEXT UNIQUE NOT NULL,
    created_at     TIMESTAMP          NOT NULL,
    result_task_id UUID 			  NOT NULL,
    status         TEXT      NOT NULL
);
CREATE TABLE IF NOT EXISTS tasks
(
    id            UUID PRIMARY KEY,
    expression_id UUID,
    a             DOUBLE PRECISION,
    b             DOUBLE PRECISION,
    operation     TEXT NOT NULL,
    result        DOUBLE PRECISION,
    is_err        BOOLEAN DEFAULT FALSE,
    error         TEXT,
    status        TEXT NOT NULL,
    next_task_id  UUID,
    created_at    TIMESTAMP
)
`

func (p *Postgres) Init() error {
	_, err := p.db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
