package main

import (
  "fmt"
  "os"
  "path/filepath"
  "github.com/deiwin/interact"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "os/exec"
)

func main() {
  actor := interact.NewActor(os.Stdin, os.Stdout)
  // ~/.userenv/ の作成
  dir := filepath.Join(os.Getenv("HOME"), ".userenv")
  // dirがなかったら
  if _, err := os.Stat(dir); err != nil {
    err := os.Mkdir(dir, os.ModePerm)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
  }

  fp := filepath.Join(os.Getenv("HOME"), ".userenv", "config.yml")
  if _, err := os.Stat(fp); err == nil {
    // 設定ファイルがある場合
    users := []User{}
    file, _ := os.Open(fp)
    bs, _ := ioutil.ReadAll(file)
    err = yaml.Unmarshal(bs, &users)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
    fmt.Printf("%+v\n", users)
    user, unknown := SelectUser(actor, users)
    if unknown {
      users = append(users, *user)
      bs, err := yaml.Marshal(users)
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(0)
      }
      ioutil.WriteFile(fp, bs, os.ModePerm)
    }
    if user == nil {
      os.Exit(1)
    }

    errName := exec.Command("git", "config", "--global", "user.name", user.BitBucketName).Run()
    errEmail := exec.Command("git", "config", "--global", "user.email", user.BitBucketEmail).Run()
    if errName != nil && errEmail != nil {
      fmt.Printf("Name:%s\n", errName.Error())
      fmt.Printf("Email:%s\n", errEmail.Error())
    } else {
      fmt.Println("ログインに成功しました。")
    }
  } else {
    // 設定ファイルがない場合
    user, err := NewUserFromPrompt(actor)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(1)
    }
    fmt.Printf("%+v\n", user)
    users := []User{*user}
    bs, err := yaml.Marshal(users)
    if err != nil {
      fmt.Println(err.Error())
      os.Exit(0)
    }
    ioutil.WriteFile(fp, bs, os.ModePerm)
  }

}
