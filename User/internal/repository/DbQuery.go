package repository

import (
	"User/internal/model"
	"context"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *model.User) error {
	query := `INSERT INTO users (name, mobile_no, email, user_type, cur_status)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`
	return r.db.QueryRowContext(context.Background(), query,
		user.Name, user.MobileNo, user.Email, user.UserType, user.CurStatus).
		Scan(&user.ID)
}

func (r *userRepo) GetByID(id int) (*model.User, error) {
	var u model.User
	query := `SELECT id, name, mobile_no, email, user_type, cur_status
			  FROM users WHERE id=$1`
	err := r.db.QueryRowContext(context.Background(), query, id).
		Scan(&u.ID, &u.Name, &u.MobileNo, &u.Email, &u.UserType, &u.CurStatus)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepo) UpdateStatus(id int, status string) error {
	query := `UPDATE users SET cur_status=$1 WHERE id=$2`
	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *userRepo) ListDrivers(status string) ([]model.User, error) {
	query := `SELECT id, name, mobile_no, email, user_type FROM users 
			  WHERE user_type='driver' AND cur_status=$1`
	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.MobileNo, &u.Email, &u.UserType)
		drivers = append(drivers, u)
	}
	return drivers, nil
}

func (r *userRepo) UserExists(phone string) bool {
	var id int
	query := "SELECT id FROM users WHERE mobile_no=$1"
	err := r.db.QueryRow(query, phone).Scan(&id)
	if err != nil {
		return false
	}
	return id > 0
}

func (r *userRepo) EmailExists(email string) bool {
	var id int
	query := "SELECT id FROM users WHERE email=$1"
	err := r.db.QueryRow(query, email).Scan(&id)
	if err != nil {
		return false
	}
	return id > 0
}
