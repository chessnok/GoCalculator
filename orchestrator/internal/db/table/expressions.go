package table

import (
	"database/sql"
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions"
)

type Expressions struct {
	db    *sql.DB
	tasks *Tasks
}

func NewExpressions(db *sql.DB) *Expressions {
	return &Expressions{
		db:    db,
		tasks: NewTasks(db),
	}
}

func (r *Expressions) New(exp *expressions.Expression) error {
	_, err := r.db.Exec("INSERT INTO expressions(id, expression,normalized_expression,result_task_id) VALUES ($1,$2,$3,$4)", exp.Id, exp.Expression, exp.NormalizedExpression, exp.ResultTaskId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Expressions) GetExpressionsList() ([]expressions.Expression, error) {
	rows, err := r.db.Query("SELECT id, expression,normalized_expression,result_task_id,status,result,created_at FROM expressions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ex := make([]expressions.Expression, 0)
	for rows.Next() {
		var expression expressions.Expression
		if err := rows.Scan(&expression.Id, &expression.Expression, &expression.NormalizedExpression, &expression.ResultTaskId, &expression.Status, &expression.Result, &expression.CreatedAt); err != nil {
			return nil, err
		}
		tsk, err := r.tasks.GetTasksByExpressionId(expression.Id)
		if err != nil {
			return nil, err
		}
		expression.Tasks = tsk
		ex = append(ex, expression)
	}
	return ex, nil
}
