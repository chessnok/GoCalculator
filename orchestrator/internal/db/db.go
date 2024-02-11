package db

import (
	"database/sql"
	"fmt"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table/agents"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

type Postgres struct {
	db *sql.DB
	*agents.Agents
}

func NewPostgres(cfg *Config) (*Postgres, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DB))
	if err != nil {
		return nil, err
	}
	return &Postgres{
		db:     db,
		Agents: agents.NewAgents(db),
	}, nil
}
func (p *Postgres) OnNewResult(message *amqp.Delivery) {
	fmt.Println(string(message.Body))
}

func (p *Postgres) Close() {
	if p.db == nil {
		return
	}
	p.db.Close()
}
