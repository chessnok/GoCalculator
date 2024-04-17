package agents

// Agent represents the table agent
type Agent struct {
	ID       string `json:"id"`
	LastPing string `json:"last_ping"`
	Pid      string `json:"pid"`
	//	LastExpression
}
