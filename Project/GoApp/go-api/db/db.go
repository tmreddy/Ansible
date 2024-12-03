package db

import (
	"sync"
	"go-api/models"
)
var userStore = struct {
	sync.RWMutex
	users []models.User
}{}

// get all users
func GetUsers() []models.User {
	userStore.RLock()
	defer userStore.RUnlock()
	return userStore.users
}

// get user by id
func GetUser(id int) *models.User {
	userStore.RLock()
	defer userStore.RUnlock()
	for _, user := range userStore.users {
		if user.ID == id {
			return &user
		}
	}	
	return nil 
}

// add user
func AddUser(user models.User) {
	userStore.Lock()
	defer userStore.Unlock()
	userStore.users = append(userStore.users, user)
}

// update user
func UpdateUser(id int, user models.User) bool{
	userStore.Lock()
	defer userStore.Unlock()
	for i, u := range userStore.users {
		if u.ID == id {
			userStore.users[i] = user
			return true
		}
	}
	return false
}

// delete user
func DeleteUser(id int) bool {
	userStore.Lock()
	defer userStore.Unlock()
	for i, user := range userStore.users {
		if user.ID == id {
			userStore.users = append(userStore.users[:i], userStore.users[i+1:]...)
			return true
		}
	}
	return false
}