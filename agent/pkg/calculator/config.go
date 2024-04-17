package calculator

import (
	"os"
	"strconv"
)

func GetWorkersCount() int {
	wk := os.Getenv("WORKERS_COUNT")
	parallelWorkers, err := strconv.Atoi(wk)
	if err != nil {
		parallelWorkers = 1
	}
	return parallelWorkers
}
