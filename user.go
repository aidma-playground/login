package main

import (
  "regexp"
  "github.com/deiwin/interact"
  "errors"
)

type User struct {
  Name string `yaml:"Name"`
  BitBucketName string `yaml:"BitBucketName"`
  BitBucketEmail string `yaml:"BitBucketEmail"`
}

func NewUserFromPrompt(actor interact.Actor) (*User, error) {
  var (
    err error
    name string
    bitBucketName string
    bitBucketEmail string
    checkNotEmpty = func(input string) error {
      if input == "" {
          return errors.New("入力が空白です。")
      }
      return nil
    }
    checkEmailAddress = func(input string) error {
      emailRegexp := regexp.MustCompile("^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:\\(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22)))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$")
      if !emailRegexp.Match([]byte(input)) {
				return errors.New("不正なメールアドレスです。")
			}
      return nil
    }
  )
  name, err = actor.Prompt("名前を入力してください。\t\t", checkNotEmpty)
  if err != nil {
    return nil, err
  }

  bitBucketName, err = actor.Prompt("BitBucketのユーザー名を入力してください。\t\t", checkNotEmpty)
  if err != nil {
    return nil, err
  }

  bitBucketEmail, err = actor.Prompt("BitBucket登録のEmailを入力してください。\t\t", checkNotEmpty, checkEmailAddress)
  if err != nil {
    return nil, err
  }
  return NewUser(name, bitBucketName, bitBucketEmail), nil
}

func NewUser(name, bitBucketName, bitBucketEmail string) *User {
  return &User{name, bitBucketName, bitBucketEmail}
}
