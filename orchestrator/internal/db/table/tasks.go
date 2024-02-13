package table

import (
	"database/sql"
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/task"
)

type Tasks struct {
	db *sql.DB
}

func (t *Tasks) newTask(tsk *task.Task) error {
	_, err := t.db.Exec("INSERT INTO tasks (id, operation, a, b, a_is_numeral, b_is_numeral, next_task_id, next_task_type, is_final, task_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", tsk.Id, tsk.Operation, tsk.A, tsk.B, tsk.AIsNumeral, tsk.BIsNumeral, tsk.NextTaskId, tsk.NextTaskType, tsk.IsFinal, tsk.ExprId)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasks) NewTasks(tasks []*task.Task) error {
	for _, tsk := range tasks {
		err := t.newTask(tsk)
		if err != nil {
			return err
		}
	}
	return nil
}
func NewTasks(db *sql.DB) *Tasks {
	return &Tasks{db: db}
}
