package models

type Message struct {
	Text   string `json:"msg"`
	Sender string `json:"senderName"`
}
