package main

import (
  "github.com/robbiev/dilemma"
  "github.com/deiwin/interact"
  "fmt"
)

func SelectUser(actor interact.Actor, users []User) (user *User, unknown bool) {
  selection := make([]string, 0, len(users) + 1)
  for _, user := range users {
    selection = append(selection, user.Name)
  }
  selection = append(selection, "新規ユーザー")

  s := dilemma.Config{
  	Title:   "ログインするユーザーを選択してください。",
  	Help:    "Use arrow up and down, then enter to select.\n\rChoose wisely.",
  	Options: selection,
  }
  sekectedUser, exitKey, err := dilemma.Prompt(s)
  if err != nil || exitKey == dilemma.CtrlC {
  	fmt.Print("Exiting...\n")
  	return nil, false
  }
  for _, user := range users {
    if sekectedUser == user.Name {
      return &user, false
    }
  }
  user, err = NewUserFromPrompt(actor)
  if err != nil {
    return nil, false
  }
  return user, true
}
