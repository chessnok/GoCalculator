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
	_, err := r.db.Exec("INSERT INTO expressions(id, expression, normalized_expression, result_task_id, userid) VALUES ($1, $2, $3, $4)", exp.Id, exp.Expression, exp.NormalizedExpression, exp.ResultTaskId, exp.UserId)
	if err != nil {
		return err
	}
	return nil
}
func (r *Expressions) scanExpressionFromRows(row *sql.Rows) (*expressions.Expression, error) {
	var expression expressions.Expression
	if err := row.Scan(&expression.Id, &expression.Expression, &expression.NormalizedExpression, &expression.ResultTaskId, &expression.Status, &expression.Result, &expression.CreatedAt, &expression.UserId); err != nil {
		return nil, err
	}
	tsk, err := r.tasks.GetTasksByExpressionId(expression.Id)
	if err != nil {
		return nil, err
	}
	expression.Tasks = tsk
	return &expression, nil
}

func (r *Expressions) expressionsFromRows(rows *sql.Rows) ([]expressions.Expression, error) {
	ex := make([]expressions.Expression, 0)
	for rows.Next() {
		expression, err := r.scanExpressionFromRows(rows)
		if err != nil {
			return nil, err
		}
		ex = append(ex, *expression)
	}
	return ex, nil
}
func (r *Expressions) GetExpressionsList() ([]expressions.Expression, error) {
	rows, err := r.db.Query("SELECT id, expression,normalized_expression,result_task_id,status,result,created_at,userid FROM expressions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.expressionsFromRows(rows)
}

func (r *Expressions) GetExpressionsListByUserId(userid string) ([]expressions.Expression, error) {
	rows, err := r.db.Query("SELECT id, expression,normalized_expression,result_task_id,status,result,created_at,userid FROM expressions WHERE userid = $1", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return r.expressionsFromRows(rows)
}
