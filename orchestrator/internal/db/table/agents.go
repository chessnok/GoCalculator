package table

import (
	"database/sql"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents"
)

type Agents struct {
	db *sql.DB
}

func (a *Agents) GetList() ([]agents.Agent, error) {
	rows, err := a.db.Query("SELECT * FROM agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var age []agents.Agent
	for rows.Next() {
		var agent agents.Agent
		if err := rows.Scan(&agent.ID, &agent.LastPing, &agent.IP, &agent.Port, &agent.Status, &agent.IsUpToDate); err != nil {
			return nil, err
		}
		age = append(age, agent)
	}
	return age, nil
}

func NewAgents(db *sql.DB) *Agents {
	return &Agents{db: db}
}
