package cqrs

import d "github.com/calogxro/cqrs-es/domain"

type WriteRepository struct {
	store map[string]*d.User
}

func NewWriteRepository() *WriteRepository {
	return &WriteRepository{
		store: make(map[string]*d.User),
	}
}

func (r *WriteRepository) getUser(userId string) *d.User {
	user, exists := r.store[userId]
	if !exists {
		return nil
	}
	return user
}

func (r *WriteRepository) addUser(userId string, user *d.User) {
	r.store[userId] = user
}

// func (r *WriteRepository) updateUser(userId string, user *d.User) {
// 	r.store[userId] = user
// }
