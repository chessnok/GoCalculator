package manager

import (
	"context"
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents/manager/http"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"log"
	"sync"
	"time"
)

type AgentManager struct {
	mu               sync.RWMutex
	currOperaion     map[string]string
	agents           *table.Agents
	timeout          time.Duration
	timer            *time.Timer
	stopped          bool
	CalculatorConfig *calculator.Config
}

func NewAgentManager(agents *table.Agents, timeout time.Duration, calculatorConfig *calculator.Config) *AgentManager {
	return &AgentManager{
		mu:               sync.RWMutex{},
		currOperaion:     make(map[string]string),
		agents:           agents,
		timeout:          timeout,
		timer:            nil,
		stopped:          false,
		CalculatorConfig: calculatorConfig,
	}
}

func (p *AgentManager) StartPing() {
	ctx := context.Background()
	go func() {
		for !p.stopped {
			p.timer = time.NewTimer(p.timeout)
			<-p.timer.C
			agnts, err := p.agents.GetAgentsList()
			if err != nil {
				log.Default().Println("Error getting agents: ", err)
				continue
			}
			wg := sync.WaitGroup{}
			wg.Add(len(agnts))
			for _, agent := range agnts {
				go func(agent agents.Agent) {
					defer wg.Done()
					var res bool
					var taskUuid string
					if agent.ConfigIsUpToDate {
						res, taskUuid = http.Ping(ctx, agent.IP)
					} else {
						res, taskUuid = http.NewConfig(ctx, agent.IP, p.CalculatorConfig)
					}
					if res {
						p.agents.SetAgentLastPing(agent.ID)
						if !agent.ConfigIsUpToDate {
							p.agents.SetAgentConfigIsUpToDate(agent.ID, true)
						}
						p.mu.Lock()
						p.currOperaion[agent.ID] = taskUuid
						p.mu.Unlock()
					} else {
						p.agents.SetAgentStatus(agent.ID, "offline")
					}
				}(agent)
			}
			wg.Wait()
			time.Sleep(p.timeout)
		}
	}()
}
func (p *AgentManager) GetCurrentOperations() map[string]string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.currOperaion
}

func (p *AgentManager) UpdateConfig() {
	ctx := context.Background()
	agnts, err := p.agents.GetAgentsList()
	if err != nil {
		log.Default().Println("Error getting agents: ", err)
		return
	}
	for _, agent := range agnts {
		s, _ := http.NewConfig(ctx, agent.IP, p.CalculatorConfig)
		if s {
			p.agents.SetAgentLastPing(agent.ID)
		} else {
			p.agents.SetAgentStatus(agent.ID, "offline")
			p.agents.SetAgentConfigIsUpToDate(agent.ID, false)
		}
	}
}
func (p *AgentManager) StopPing() {
	p.stopped = true
	p.timer.Stop()
}
