package calculator

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ParallelWorkers  int           `json:"parallel_workers" `
	AddExecutionTime time.Duration `json:"add_execution_time"`
	SubExecutionTime time.Duration `json:"sub_execution_time"`
	MulExecutionTime time.Duration `json:"mul_execution_time"`
	DivExecutionTime time.Duration `json:"div_execution_time"`
}

func NewConfig() *Config {
	wk := os.Getenv("WORKERS_COUNT")
	parallelWorkers, err := strconv.Atoi(wk)
	if err != nil {
		parallelWorkers = 1
	}
	addTime, subTime, mulTime, divTime := 0, 0, 0, 0
	if len(os.Args) > 5 {
		addTime, err = strconv.Atoi(os.Args[2])
		if err != nil {
			addTime = 1
		}
		subTime, _ = strconv.Atoi(os.Args[3])
		if err != nil {
			subTime = 1
		}
		mulTime, _ = strconv.Atoi(os.Args[4])
		if err != nil {
			mulTime = 1
		}
		divTime, _ = strconv.Atoi(os.Args[5])
		if err != nil {
			divTime = 1
		}
	}
	return &Config{
		ParallelWorkers:  parallelWorkers,
		AddExecutionTime: time.Duration(addTime) * time.Millisecond,
		SubExecutionTime: time.Duration(subTime) * time.Millisecond,
		MulExecutionTime: time.Duration(mulTime) * time.Millisecond,
		DivExecutionTime: time.Duration(divTime) * time.Millisecond,
	}
}
