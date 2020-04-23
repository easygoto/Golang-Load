package cg

import (
	"encoding/json"
	"errors"

	"demo/cgapp/ipc"
)

type CenterClient struct {
	*ipc.Client
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}
	response, err := client.Call("addPlayer", string(b))
	if err == nil && response.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	response, _ := client.Call("removePlayer", name)
	if response.Code == "200" {
		return nil
	}
	return errors.New(response.Code)
}

func (client *CenterClient) ListPlayer(params string) (players []*Player, err error) {
	response, _ := client.Call("listPlayer", params)
	if response.Code != "200" {
		err = errors.New(response.Code)
		return
	}
	err = json.Unmarshal([]byte(response.Body), &players)
	return
}

// 广播
func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	response, _ := client.Call("broadcast", string(b))
	if response.Code == "200" {
		return nil
	}
	return errors.New(response.Code)
}
