package table

import (
	"database/sql"
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/task"
	pb "github.com/chessnok/GoCalculator/proto"
	"time"
)

type Tasks struct {
	db *sql.DB
}

func (t *Tasks) newTask(tsk *task.Task) error {
	_, err := t.db.Exec("INSERT INTO tasks (id, operation, a, b, a_is_numeral, b_is_numeral, next_task_id, next_task_type, is_final, expression_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", tsk.Id, tsk.Operation, tsk.A, tsk.B, tsk.AIsNumeral, tsk.BIsNumeral, tsk.NextTaskId, tsk.NextTaskType, tsk.IsFinal, tsk.ExprId)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasks) New(tasks []*task.Task) error {
	for _, tsk := range tasks {
		err := t.newTask(tsk)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t *Tasks) GetTaskById(id string) (*task.Task, error) {
	row := t.db.QueryRow("SELECT id, operation, a, b, a_is_numeral, b_is_numeral, next_task_id, next_task_type, is_final, expression_id FROM tasks WHERE id = $1", id)
	var tsk task.Task
	if err := row.Scan(&tsk.Id, &tsk.Operation, &tsk.A, &tsk.B, &tsk.AIsNumeral, &tsk.BIsNumeral, &tsk.NextTaskId, &tsk.NextTaskType, &tsk.IsFinal, &tsk.Status); err != nil {
		return nil, err
	}
	return &tsk, nil
}

func (t *Tasks) GetTasksByExpressionId(id string) ([]*task.Task, error) {
	rows, err := t.db.Query("SELECT id, operation, a, b, a_is_numeral, b_is_numeral, next_task_id, next_task_type, is_final, status FROM tasks WHERE expression_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tsks := make([]*task.Task, 0)
	for rows.Next() {
		tsk := task.Task{ExprId: id}
		if err := rows.Scan(&tsk.Id, &tsk.Operation, &tsk.A, &tsk.B, &tsk.AIsNumeral, &tsk.BIsNumeral, &tsk.NextTaskId, &tsk.NextTaskType, &tsk.IsFinal, &tsk.Status); err != nil {
			return nil, err
		}
		tsks = append(tsks, &tsk)
	}
	return tsks, nil
}

func (t *Tasks) SelectTasksToSendToQueue(config *pb.Config) ([]*task.Task, error) {
	var maxTime int64
	maxTime = config.AddExecutionTime
	if config.MulExecutionTime > maxTime {
		maxTime = config.MulExecutionTime
	}
	if config.DivExecutionTime > maxTime {
		maxTime = config.DivExecutionTime
	}
	if config.DivExecutionTime > maxTime {
		maxTime = config.DivExecutionTime
	}
	rows, err := t.db.Query(`
    SELECT id, operation, a, b 
    FROM tasks 
    WHERE 
        (a_is_numeral = true AND b_is_numeral = true AND status = 'pending') OR 
        (status = 'in_queue' AND time < NOW() - INTERVAL '3 minutes' * $1)
`, 3*maxTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tsks := make([]*task.Task, 0)
	for rows.Next() {
		var tsk task.Task
		if err := rows.Scan(&tsk.Id, &tsk.Operation, &tsk.A, &tsk.B); err != nil {
			return nil, err
		}
		tsks = append(tsks, &tsk)
	}
	return tsks, nil
}
func NewTasks(db *sql.DB) *Tasks {
	return &Tasks{db: db}
}

func (t *Tasks) UpdateTaskStatus(id string, status string) error {
	_, err := t.db.Exec("UPDATE tasks SET status = $1, time = $3 WHERE id = $2", status, id, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasks) TaskResult(id string, result float64, isErr bool) error {
	err := t.UpdateTaskStatus(id, "done")
	if err != nil {
		return err
	}
	task, err := t.GetTaskById(id)
	if err != nil {
		return err
	}
	if isErr {
		err := t.UpdateTaskStatus(id, "error")
		if err != nil {
			return err
		}
		_, err = t.db.Exec("UPDATE expressions SET status = 'error' WHERE id = $1", task.ExprId)
		if err != nil {
			return err
		}
		return nil
	}
	if task.IsFinal {
		_, err = t.db.Exec("UPDATE expressions SET status = 'done', result= $1 WHERE id = $2", result, task.ExprId)
		if err != nil {
			return err
		}
	} else {
		if task.NextTaskType {
			_, err := t.db.Exec("UPDATE tasks SET b = $1, b_is_numeral = true WHERE id = $2", result, task.NextTaskId)
			if err != nil {
				return err
			}
		} else {
			_, err := t.db.Exec("UPDATE tasks SET a = $1, a_is_numeral = true WHERE id = $2", result, task.NextTaskId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
