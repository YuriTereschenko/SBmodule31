package models

type DBinteraction interface {
	CreateNewUser(*User) (newUserID int, err error)
	MakeFriends(sourceID, targetID int) (sourceName string, targetName string, err error)
	DelUser(id int) (username string, err error)
	GetFriendsList(userID int) (userName string, userFriends string, err error)
	UpdateUserAge(userID, newAge int) (userName string, err error)
	GetUsers() ([]User, error)
}
