package crud

import d "github.com/calogxro/cqrs-es/domain"

type UserRepository struct {
	store map[string]*d.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		store: make(map[string]*d.User),
	}
}

func (r *UserRepository) getUser(userId string) *d.User {
	user, exists := r.store[userId]
	if !exists {
		return nil
	}
	return user
}

func (r *UserRepository) addUser(user *d.User) {
	r.store[user.GetId()] = user
}

// func (r *UserRepository) updateUser(user *d.User) {
// 	r.store[user.GetId()] = user
// }
