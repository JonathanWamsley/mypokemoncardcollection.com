package main

import (
  "fmt"

  "mypokemoncardcollection.com/models"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "password"
  dbname   = "mypokemoncardcollection_dev"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  us, err := models.NewUserService(psqlInfo)
  if err != nil {
    panic(err)
  }
  defer us.Close()
  us.DestructiveReset()

  user := models.User{
	  Name: "test name",
	  Email: "test@email.com",
  }
  if err := us.Create(&user); err != nil {
	  panic(err)
  }

  foundUser, err := us.ByID(1)
  if err != nil {
    panic(err)
  }
  fmt.Println(foundUser)
}