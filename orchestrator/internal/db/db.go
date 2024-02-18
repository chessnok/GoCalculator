package db

import (
	"database/sql"
	"fmt"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db          *sql.DB
	Agents      *table.Agents
	Tasks       *table.Tasks
	Expressions *table.Expressions
}

func NewPostgres(cfg *Config) (*Postgres, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DB))
	if err != nil {
		return nil, err
	}
	return &Postgres{
		db:          db,
		Agents:      table.NewAgents(db),
		Tasks:       table.NewTasks(db),
		Expressions: table.NewExpressions(db),
	}, nil
}

func (p *Postgres) Close() {
	if p.db == nil {
		return
	}
	p.db.Close()
}
