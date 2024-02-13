package db

var schema = `
CREATE TABLE IF NOT EXISTS agents
(
    id        TEXT PRIMARY KEY,
    last_ping TIMESTAMP,
    ip       TEXT NOT NULL,
    port     INTEGER,
    status    TEXT NOT NULL DEFAULT 'active',
    config_is_up_to_date BOOLEAN DEFAULT TRUE
);
CREATE TABLE IF NOT EXISTS expressions
(
    id             TEXT PRIMARY KEY,
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    result_task_id TEXT NOT NULL,
    status         TEXT NOT NULL DEFAULT 'pending'
);
CREATE TABLE IF NOT EXISTS tasks
(
    id            TEXT PRIMARY KEY,
    operation     VARCHAR(1) NOT NULL,
    a             FLOAT,
    b             FLOAT,
    task_id 	  TEXT,
    a_is_numeral BOOLEAN,
    b_is_numeral BOOLEAN,
    next_task_id  TEXT,
    next_task_type TEXT,
    is_final      BOOLEAN NOT NULL
)
`

func (p *Postgres) Init() error {
	_, err := p.db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
