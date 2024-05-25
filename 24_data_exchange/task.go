package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Friends []string `json:"friends"`
}

var (
	users map[string]User
	mu    sync.RWMutex
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	users[newUser.Name] = newUser

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Пользователь %s создан.", newUser.Name)
}

func makeFriends(w http.ResponseWriter, r *http.Request) {
	var friendship struct {
		InitiatorID string `json:"initiator_id"`
		FriendID    string `json:"friend_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&friendship)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	initiator, foundInitiator := users[friendship.InitiatorID]
	newFriend, foundFriend := users[friendship.FriendID]

	if !foundInitiator || !foundFriend {
		http.Error(w, "Пользователь не найден.", http.StatusNotFound)
		return
	}

	initiator.Friends = append(initiator.Friends, friendship.FriendID)
	newFriend.Friends = append(newFriend.Friends, friendship.InitiatorID)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Дружба установлена.")
}

func main() {
	users = make(map[string]User)

	http.HandleFunc("/createUser", createUser)
	http.HandleFunc("/makeFriends", makeFriends)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
