package agents

import "database/sql"

type Agents struct {
	db *sql.DB
}

// Agent represents the table agents
type Agent struct {
	ID         string `json:"id"`
	LastPing   string `json:"last_ping"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	Status     string `json:"status"`
	IsUpToDate bool   `json:"config_is_up_to_date"`
}

func (a *Agents) GetList() ([]Agent, error) {
	rows, err := a.db.Query("SELECT * FROM agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var agents []Agent
	for rows.Next() {
		var agent Agent
		if err := rows.Scan(&agent.ID, &agent.LastPing, &agent.IP, &agent.Port, &agent.Status, &agent.IsUpToDate); err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}
	return agents, nil
}

func NewAgents(db *sql.DB) *Agents {
	return &Agents{db: db}
}
