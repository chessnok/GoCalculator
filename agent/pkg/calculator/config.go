package calculator

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	ParallelWorkers  int           `json:"-"`
	AddExecutionTime time.Duration `json:"add_execution_time"`
	SubExecutionTime time.Duration `json:"sub_execution_time"`
	MulExecutionTime time.Duration `json:"mul_execution_time"`
	DivExecutionTime time.Duration `json:"div_execution_time"`
}

func GetWorkersCount() int {
	wk := os.Getenv("WORKERS_COUNT")
	parallelWorkers, err := strconv.Atoi(wk)
	if err != nil {
		parallelWorkers = 1
	}
	return parallelWorkers
}
func NewConfigFromArgs() *Config {
	parallelWorkers := GetWorkersCount()
	defaultTime := time.Millisecond * 1
	addTime, subTime, mulTime, divTime := defaultTime, defaultTime, defaultTime, defaultTime
	if len(os.Args) > 5 {
		at, err := strconv.Atoi(os.Args[2])
		if err == nil {
			addTime = time.Duration(at)
		}
		st, err := strconv.Atoi(os.Args[3])
		if err == nil {
			subTime = time.Duration(st)
		}
		mt, err := strconv.Atoi(os.Args[4])
		if err == nil {
			mulTime = time.Duration(mt)
		}
		dt, err := strconv.Atoi(os.Args[5])
		if err == nil {
			divTime = time.Duration(dt)
		}
	}
	return &Config{
		ParallelWorkers:  parallelWorkers,
		AddExecutionTime: addTime,
		SubExecutionTime: subTime,
		MulExecutionTime: mulTime,
		DivExecutionTime: divTime,
	}
}

func NewConfig() *Config {
	return &Config{
		ParallelWorkers:  1,
		AddExecutionTime: 1 * time.Millisecond,
		SubExecutionTime: 1 * time.Millisecond,
		MulExecutionTime: 1 * time.Millisecond,
		DivExecutionTime: 1 * time.Millisecond,
	}
}
