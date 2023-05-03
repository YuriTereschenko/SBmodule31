package postgres

import (
	"database/sql"
	"fmt"
	"main/internal/models"
	"strconv"
	"strings"
)

type Storage struct {
	DB *sql.DB
}

func (s *Storage) GetUsers() ([]models.User, error) {
	rows, err := s.DB.Query("select name, age from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Storage) InsertUser(u *models.User) (int, error) {
	_, err := s.DB.Exec("insert into users(name, age) values ($1, $2)", u.Name, u.Age)
	if err != nil {
		return 0, err
	}
	var id int
	err = s.DB.QueryRow("select id from users order by id desc limit 1").Scan(&id)
	if err != nil {
		return 0, err
	}
	for _, friendID := range u.Friends {
		fmt.Println("What")
		_, _, err := s.InsertFriend(int(id), friendID)
		if err != nil {
			return 0, err
		}
	}
	fmt.Println("What id", id)
	return id, nil
}

func (s *Storage) InsertFriend(sourceID, targetID int) (string, string, error) {
	sourceName, err := s.getNameByID(sourceID)
	if err != nil {
		return "", "", err
	}
	targetName, err := s.getNameByID(targetID)
	if err != nil {
		return "", "", err
	}

	_, err = s.DB.Exec("insert into friends(user_id_one, user_id_two) values ($1, $2)", sourceID, targetID)
	if err != nil {
		return "", "", err
	}

	return sourceName, targetName, nil
}

func (s *Storage) getNameByID(userID int) (string, error) {
	var name string
	fmt.Println("i'm in get name, id:", userID)
	err := s.DB.QueryRow("select name from users where id = $1", userID).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (s *Storage) GetFriends(userID int) (string, string, error) {
	username, err := s.getNameByID(userID)
	if err != nil {
		return "", "", err
	}

	rows, err := s.DB.Query("select u.name, u.id from users u join friends f "+
		"on u.id = f.user_id_two where f.user_id_one = $1", userID)
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	var friends strings.Builder
	for rows.Next() {
		var (
			name string
			id   int
		)
		err := rows.Scan(&name, &id)
		if err != nil {
			return "", "", err
		}

		friends.WriteString(fmt.Sprintf("Name: %v, id: %v\n", name, strconv.Itoa(id)))

		err = rows.Err()
		if err != nil {
			return "", "", err
		}
	}
	return username, friends.String(), nil
}

func (s *Storage) DelUser(targetID int) (string, error) {
	userName, err := s.getNameByID(targetID)
	if err != nil {
		return "", err
	}
	_, err = s.DB.Exec("delete from friends where user_id_one = $1 or user_id_two = $1", targetID)
	if err != nil {
		return "", err
	}
	_, err = s.DB.Exec("delete from users where id = $1", targetID)
	if err != nil {
		return "", err
	}

	return userName, nil
}

func (s *Storage) UpdateAge(id, newAge int) (string, error) {
	userName, err := s.getNameByID(id)
	if err != nil {
		return "", err
	}
	_, err = s.DB.Exec("update users set age=$1 where id=$2", newAge, id)
	if err != nil {
		return "", err
	}

	return userName, nil
}
