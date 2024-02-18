package db

var schema = `
CREATE TABLE IF NOT EXISTS Agents
(
    id        TEXT PRIMARY KEY,
    last_ping TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip       TEXT NOT NULL,
    port     INTEGER NOT NULL,
    status    TEXT NOT NULL DEFAULT 'online',
    config_is_up_to_date BOOLEAN DEFAULT FALSE
);
CREATE TABLE IF NOT EXISTS Expressions
(
    id             TEXT PRIMARY KEY,
    expression     TEXT,
    normalized_expression TEXT,
    result_task_id TEXT NOT NULL,
    status         TEXT NOT NULL DEFAULT 'pending',
    result         FLOAT DEFAULT 0.0,
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS Tasks
(
    id            TEXT PRIMARY KEY,
    operation     VARCHAR(1) NOT NULL,
    a             FLOAT,
    b             FLOAT,
    expression_id 	  TEXT,
    a_is_numeral BOOLEAN,
    b_is_numeral BOOLEAN,
    next_task_id  TEXT,
    next_task_type TEXT,
    is_final      BOOLEAN NOT NULL,
    status        TEXT NOT NULL DEFAULT 'pending'
)
`

func (p *Postgres) Init() error {
	_, err := p.db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
