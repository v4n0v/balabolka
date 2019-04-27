package models

type Message struct {
	Text   string `json:"text"`
	Sender string `json:"sender"`
	Time   string `json:"dt"`
}
