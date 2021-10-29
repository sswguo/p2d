package main

type message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var messageList = []message{
	message{ID: 1, Title: "Message 1", Content: "Message info 1"},
	message{ID: 2, Title: "Message 2", Content: "Message info 2"},
  }

 func getAllMessages() []message {
	return messageList
 }