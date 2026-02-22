package repository

import (
	"User/internal/model"
	"User/internal/utils"
	"log"
)

type UserQuery struct {
	repo UserRepository
}

func NewUserQuery(repo UserRepository) *UserQuery {
	return &UserQuery{repo: repo}
}

func (q *UserQuery) GetUser(id int) (*model.UserResponse, error) {
	cached, err := GetUserCache(id)
	if err == nil && cached != nil {
		log.Println("Cache Hit for User ID:", id)
		return cached, nil
	}

	log.Printf("No Cache User ID: %d", id)
	user, err := q.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := utils.CreateQueryResponse(user)

	if err := SetUserCache(id, response); err != nil {
		log.Printf("[CACHE ERROR] Failed to cache: %v", err)
	}

	return response, nil
}
