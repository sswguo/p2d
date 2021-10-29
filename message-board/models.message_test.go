package main

import "testing"

// Test the function that fetches all messages
func TestGetAllMessages(t *testing.T) {
  mlist := getAllMessages()

  // Check that the length of the list of articles returned is the
  // same as the length of the global variable holding the list
  if len(mlist) != len(messageList) {
    t.Fail()
  }

  // Check that each member is identical
  for i, v := range mlist {
    if v.Content != messageList[i].Content ||
      v.ID != messageList[i].ID ||
      v.Title != messageList[i].Title {

      t.Fail()
      break
    }
  }
}