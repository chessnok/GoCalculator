package agents

// Agent represents the table agents
type Agent struct {
	ID         string `json:"id"`
	LastPing   string `json:"last_ping"`
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	Status     string `json:"status"`
	IsUpToDate bool   `json:"config_is_up_to_date"`
}
