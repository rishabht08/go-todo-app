package repositories

import (
	"context"
	"database/sql"
	"time"
)

type PostgresRepo struct {
	DbConn *sql.DB
}

var DbRepo *PostgresRepo

func NewRepo(db *sql.DB) *PostgresRepo {
	DbRepo = &PostgresRepo{
		DbConn: db,
	}

	return DbRepo
}

func (m *PostgresRepo) CreateDummyUser() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `insert into users (first_name , last_name , email , password, created_at , updated_at)
		values ('Rishab' , 'Tiwari' , 'rishabht123@gmail.com' , 'rishabh3t' , $1 , $2)`

	_, err := DbRepo.DbConn.ExecContext(ctx, stmt, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
