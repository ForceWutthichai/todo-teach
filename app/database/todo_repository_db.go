package database //เชื่อม database

import (
	"context"

	"todo/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, createTodoRequest *models.CreateTodoRequest) error
	ReadTodo(ctx context.Context, req *models.ReadTodoRequest) (*[]models.ResponseReadTodo, error)
	UpdateTodo(ctx context.Context, UpdateTodoRequest *models.UpdateTodoRequest) error
	DeleteTodo(ctx context.Context, req *models.DeleteTodo) error
	ReadTodoAll(ctx context.Context) (*[]models.ResponseReadTodoAll, error)
}

type TodoRepositoryDB struct {
	pool *pgxpool.Pool
}

func NewTodoRepositoryDB(pool *pgxpool.Pool) TodoRepository {
	return &TodoRepositoryDB{
		pool: pool,
	}
}

func (r *TodoRepositoryDB) CreateTodo(ctx context.Context, createTodoRequest *models.CreateTodoRequest) error {
	tx, err := r.pool.Begin(ctx) //เริ่มทำงาน
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit(ctx)
		default:
			_ = tx.Rollback(ctx)
		}
	}() //

	stmt := `INSERT INTO todo (task,completed)
		VALUES(@todo_name, @is_check);`
	args := pgx.NamedArgs{
		"todo_name": createTodoRequest.TodoName,
		"is_check":  createTodoRequest.IsCheck,
	}

	_, err = tx.Exec(ctx, stmt, args)
	if err != nil {
		return err
	}

	return err
}

func (r *TodoRepositoryDB) ReadTodo(ctx context.Context, req *models.ReadTodoRequest) (*[]models.ResponseReadTodo, error) {
	query := `SELECT t.id, t.task, t.completed
	FROM todo t
	WHERE 1=1`
	args := []interface{}{}

	if req.TodoName != "" {
		query += " AND t.task = $1"
		args = append(args, req.TodoName)
	}

	rows, err := r.pool.Query(ctx, query, args...) // pass args here
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responseReadTodoList []models.ResponseReadTodo
	for rows.Next() {
		var responseReadTodo models.ResponseReadTodo
		err := rows.Scan(
			&responseReadTodo.Id,
			&responseReadTodo.TodoName,
			&responseReadTodo.IsCheck,
		)
		if err != nil {
			return nil, err
		}
		responseReadTodoList = append(responseReadTodoList, responseReadTodo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(responseReadTodoList) == 0 {
		return &[]models.ResponseReadTodo{}, nil
	}

	return &responseReadTodoList, nil
}

// func (r *TodoRepositoryDB) ReadTodo(ctx context.Context, req *models.ReadTodoRequest) (*[]models.ResponseReadTodo, error) {
// 	query := `SELECT t.id, t.task, t.completed
// 	FROM todo t
// 	WHERE 1=1 `

// 	if req.IsCheck != nil {
// 		query += "AND t.completed = " + fmt.Sprint(*req.IsCheck)
// 	}
//
// 	rows, err := r.pool.Query(ctx, query) //เริ่มทำงานtodo_name
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var responseReadTodoList []models.ResponseReadTodo
// 	for rows.Next() {
// 		var responseReadTodo models.ResponseReadTodo
// 		err := rows.Scan(
// 			&responseReadTodo.Id,
// 			&responseReadTodo.TodoName,
// 			&responseReadTodo.IsCheck,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		responseReadTodoList = append(responseReadTodoList, responseReadTodo)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	if len(responseReadTodoList) == 0 {
// 		return &[]models.ResponseReadTodo{}, nil
// 	}

// 	return &responseReadTodoList, nil
// }

func (r *TodoRepositoryDB) UpdateTodo(ctx context.Context, req *models.UpdateTodoRequest) error {
	tx, err := r.pool.Begin(ctx) //เริ่มทำงาน
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit(ctx)
		default:
			_ = tx.Rollback(ctx)
		}
	}() //

	stmt := `UPDATE todo 
	SET task = @todo_name,
	Completed = @is_check
	WHERE id = @id;`
	args := pgx.NamedArgs{
		"id":        req.Id,
		"todo_name": req.TodoName,
		"is_check":  req.IsCheck,
	}

	_, err = tx.Exec(ctx, stmt, args)
	if err != nil {
		return err
	}

	return err
}

func (r *TodoRepositoryDB) DeleteTodo(ctx context.Context, req *models.DeleteTodo) error {
	tx, err := r.pool.Begin(ctx) //เริ่มทำงาน
	if err != nil {
		return err
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit(ctx)
		default:
			_ = tx.Rollback(ctx)
		}
	}() //

	stmt := `DELETE FROM todo
	WHERE id = @id;`
	args := pgx.NamedArgs{
		"id": req.Id,
	}

	_, err = tx.Exec(ctx, stmt, args)
	if err != nil {
		return err
	}

	return err
}

func (r *TodoRepositoryDB) ReadTodoAll(ctx context.Context) (*[]models.ResponseReadTodoAll, error) {
	query := `SELECT t.id,t.task,t.completed
	FROM todo t ;`

	rows, err := r.pool.Query(ctx, query) //เริ่มทำงาน
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responseReadTodoList []models.ResponseReadTodoAll
	for rows.Next() {
		var responseReadTodo models.ResponseReadTodoAll
		err := rows.Scan(
			&responseReadTodo.Id,
			&responseReadTodo.TodoName,
			&responseReadTodo.IsCheck,
		)
		if err != nil {
			return nil, err
		}
		responseReadTodoList = append(responseReadTodoList, responseReadTodo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(responseReadTodoList) == 0 {
		return &[]models.ResponseReadTodoAll{}, nil
	}

	return &responseReadTodoList, nil
}
