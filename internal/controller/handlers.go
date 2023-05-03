package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"main/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request, dbi models.DBinteraction) {
	fmt.Println("Data was received from this server")
	users, err := dbi.GetUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var usersData strings.Builder
	for _, user := range users {
		usersData.WriteString(fmt.Sprintf("Name: %v, Age: %v\n", user.Name, user.Age))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(usersData.String()))
}

func CreateHandler(w http.ResponseWriter, r *http.Request, dbi models.DBinteraction) {
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	var u models.User

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	userID, err := dbi.CreateNewUser(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User %v created. ID: %v\n", u.Name, userID)))

}

func MakeFriendsHandler(w http.ResponseWriter, r *http.Request, dbi models.DBinteraction) {
	sourceID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect URL\n"))
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	defer r.Body.Close()

	target := struct {
		ID int `json:"targetID"`
	}{}

	if err := json.Unmarshal(content, &target); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	fmt.Println("In make friendsHandler, target id:", target.ID)
	fmt.Printf("handler\n s id:%v\n t id:%v\n", sourceID, target.ID)
	sourceName, targetName, err := dbi.MakeFriends(sourceID, target.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v and %v are friends now\n", sourceName, targetName)))
}

func GetFriendsHandler(w http.ResponseWriter, r *http.Request, dbi models.DBinteraction) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect url: " + err.Error() + "\n"))
		return
	}

	userName, userFriends, err := dbi.GetFriendsList(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v's friends:\n%v", userName, userFriends)))
}

func DelUserHandler(w http.ResponseWriter, r *http.Request, dbi models.DBinteraction) {
	targetID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect URL" + err.Error() + "\n"))
		return
	}

	username, err := dbi.DelUser(targetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User %v deleted\n", username)))
}

func AgeUpdateHandler(w http.ResponseWriter, r *http.Request, dbi models.DBinteraction) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("incorrect URL\n"))
		return
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	data := struct {
		NewAge int `json:"new age"`
	}{}

	if err := json.Unmarshal(content, &data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	userName, err := dbi.UpdateUserAge(userID, data.NewAge)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v's age has been updated\n", userName)))
}
