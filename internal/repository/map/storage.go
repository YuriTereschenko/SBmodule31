package _map

import (
	"fmt"
	"main/internal/models"
	"strings"
)

type Service struct {
	maxID int
	store map[int]*models.User
}

func NewService() *Service {
	return &Service{
		maxID: 0,
		store: make(map[int]*models.User),
	}
}

func (s *Service) InsertUser(u *models.User) (int, error) {
	s.maxID++
	s.store[s.maxID] = u
	return s.maxID, nil
}

func (s *Service) InsertFriend(sourceID, targetID int) (string, string, error) {
	s.store[sourceID].Friends = append(s.store[sourceID].Friends, targetID)
	s.store[targetID].Friends = append(s.store[targetID].Friends, sourceID)
	return s.store[sourceID].Name, s.store[targetID].Name, nil
}

func (s *Service) GetFriends(userID int) (string, string, error) {
	var friendsList strings.Builder
	for _, friendID := range s.store[userID].Friends {
		friendsList.WriteString(fmt.Sprintf("Name: %v, id: %v\n", s.store[friendID].Name, friendID))
	}
	return s.store[userID].Name, friendsList.String(), nil
}

func (s *Service) DelUser(targetID int) (string, error) {
	for _, friendID := range s.store[targetID].Friends {
		for i, id := range s.store[friendID].Friends {
			if id == targetID {
				s.store[friendID].Friends = append(s.store[friendID].Friends[:i], s.store[friendID].Friends[i+1:]...)
				break
			}
		}
	}
	username := s.store[targetID].Name
	delete(s.store, targetID)
	return username, nil
}

func (s *Service) UpdateAge(userID, newAge int) (string, error) {
	s.store[userID].Age = newAge
	return s.store[userID].Name, nil
}
