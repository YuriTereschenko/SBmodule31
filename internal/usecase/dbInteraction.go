package usecase

import (
	"main/internal/models"
)

type DBinter struct {
	storage models.Storage
}

func NewDBinter(stor models.Storage) *DBinter {
	return &DBinter{storage: stor}
}
func (d *DBinter) CreateNewUser(u *models.User) (int, error) {
	return d.storage.InsertUser(u)
}

func (d *DBinter) MakeFriends(sourceID, targetID int) (string, string, error) {
	return d.storage.InsertFriend(sourceID, targetID)
}

func (d *DBinter) DelUser(id int) (string, error) {
	return d.storage.DelUser(id)
}

func (d *DBinter) GetFriendsList(userID int) (string, string, error) {
	return d.storage.GetFriends(userID)
}

func (d *DBinter) UpdateUserAge(userID, newAge int) (string, error) {
	return d.storage.UpdateAge(userID, newAge)
}

func (d *DBinter) GetUsers() ([]models.User, error) {
	return d.storage.GetUsers()
}
