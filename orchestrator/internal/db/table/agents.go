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
		if err := rows.Scan(&agent.ID, &agent.LastPing, &agent.Pid); err != nil {
			return nil, err
		}
		age = append(age, agent)
	}
	return age, nil
}

func (a *Agents) GetAgentById(id string) (*agents.Agent, error) {
	row := a.db.QueryRow("SELECT last_ping, pid FROM agents WHERE id = $1", id)
	agnt := agents.Agent{ID: id}
	err := row.Scan(&agnt.LastPing, &agnt.Pid)
	if err != nil {
		return nil, err
	}
	return &agnt, nil
}
func (a *Agents) NewAgent(id, pid string) error {
	_, err := a.db.Exec("INSERT INTO agents (id, pid) VALUES ($1, $2)", id, pid)
	if err != nil {
		return err
	}
	return nil
}
func (a *Agents) SetAgentLastPing(id string) error {
	_, err := a.db.Exec("UPDATE agents SET last_ping = $1 WHERE id = $2", time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}
func (a *Agents) DeleteAgent(id string) error {
	_, err := a.db.Exec("DELETE FROM agents WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func NewAgents(db *sql.DB) *Agents {
	return &Agents{db: db}
}
