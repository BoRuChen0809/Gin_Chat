package model

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Event   string `json:"event"`
	Name    string `json:"name"`
	ID      string `json:"id"`
	Content string `json:"content"`
}

func ParseMessage(msg []byte) (*Message, error) {
	Msg := &Message{}
	err := json.Unmarshal(msg, Msg)
	return Msg, err
}

func NewChatMessage(id, name, content string) *Message {
	return &Message{Event: "Message", ID: id, Name: name, Content: content}
}
func NewConnectMessage(id, name string) *Message {
	return &Message{Event: "Connect", ID: id, Name: name, Content: fmt.Sprintf("%s 加入聊天室", name)}
}
func NewCloseMessage(id, name string) *Message {
	return &Message{Event: "Close", ID: id, Name: name, Content: fmt.Sprintf("%s 離開聊天室", name)}
}

func (msg *Message) GetJSON() []byte {
	json, _ := json.Marshal(msg)
	return json
}
