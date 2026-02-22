package repository

import "User/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	UpdateStatus(id int, status string) error
	ListDrivers(status string) ([]model.User, error)
	UserExists(Phone string) bool
	EmailExists(Email string) bool
}
