package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Players struct {
	Name     string `json:"name,omitempty"`
	Win      int    `json:"win,omitempty"`
	Loss     int    `json:"loss,omitempty"`
	Nickname string `json:"nickname"`
}

func SendPlayerToServer(player Players) {
	// Сериализуем структуру player в JSON
	playerData, err := json.Marshal(player)
	if err != nil {
		fmt.Println("Error marshalling player data:", err)
		return
	}

	// Отправляем POST-запрос с JSON-данными
	resp, err := http.Post("http://94.142.137.149:8080/add", "application/json", bytes.NewBuffer(playerData))
	if err != nil {
		fmt.Println("Error sending player to server:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Player added to server:", player.Nickname)
}

func GetPlayersFromServer() []Players {
	resp, err := http.Get("http://94.142.137.149:8080/players")
	if err != nil {
		fmt.Println("Error getting players from server:", err)
		return nil
	}
	defer resp.Body.Close()

	var players []Players
	json.NewDecoder(resp.Body).Decode(&players)

	return players
}

func DeletePlayerOnServer(nickname string) {
	req, err := http.NewRequest("DELETE", "http://94.142.137.149:8080/player/"+nickname, nil)
	if err != nil {
		fmt.Println("Error deleting player on server:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting player on server:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Player deleted on server:", nickname)
}
