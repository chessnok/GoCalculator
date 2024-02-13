package table

import (
	"database/sql"
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions"
)

type Expressions struct {
	db *sql.DB
}

func NewExpressions(db *sql.DB) *Expressions {
	return &Expressions{db: db}
}

func (r *Expressions) NewExpression(exp *expressions.Expression) error {
	_, err := r.db.Exec("INSERT INTO expressions(id, result_task_id) VALUES ($1,$2)", exp.Id, exp.ResultTaskId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Expressions) UpdateExpression(exp *expressions.Expression) error {
	_, err := r.db.Exec("UPDATE expressions SET result_task_id=$1 WHERE id=$2", exp.ResultTaskId, exp.Id)
	if err != nil {
		return err
	}
	return nil
}
