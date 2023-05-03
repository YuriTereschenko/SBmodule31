package models

type Storage interface {
	InsertUser(user2 *User) (int, error)
	InsertFriend(sourceID, targetID int) (string, string, error)
	GetFriends(userID int) (username string, userFriends string, err error)
	DelUser(targetID int) (userName string, err error)
	UpdateAge(id, newAge int) (userName string, err error)
	GetUsers() ([]User, error)
}
