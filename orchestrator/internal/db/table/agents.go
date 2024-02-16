package table

import (
	"database/sql"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents"
	"time"
)

type Agents struct {
	db *sql.DB
}

func (a *Agents) GetAgentsList() ([]agents.Agent, error) {
	rows, err := a.db.Query("SELECT * FROM agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	age := make([]agents.Agent, 0)
	for rows.Next() {
		var agent agents.Agent
		if err := rows.Scan(&agent.ID, &agent.LastPing, &agent.IP, &agent.Port, &agent.Status, &agent.ConfigIsUpToDate); err != nil {
			return nil, err
		}
		age = append(age, agent)
	}
	return age, nil
}

func (a *Agents) SetAgentConfigIsUpToDate(id string, agentConfigIsUpToDate bool) error {
	_, err := a.db.Exec("UPDATE agents SET config_is_up_to_date = $1 WHERE id = $2", agentConfigIsUpToDate, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Agents) SetAgentStatus(id string, status string) error {
	_, err := a.db.Exec("UPDATE agents SET status = $1 WHERE id = $2", status, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Agents) SetAgentLastPing(id string) error {
	_, err := a.db.Exec("UPDATE agents SET last_ping = $1, status='online' WHERE id = $2", time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
func NewAgents(db *sql.DB) *Agents {
	return &Agents{db: db}
}
