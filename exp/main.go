package main

import (
	"fmt"

	"mypokemoncardcollection.com/models"
	"mypokemoncardcollection.com/rand"
	"mypokemoncardcollection.com/hash"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mypokemoncardcollection_dev"
)

func expQueries() {
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
		Name:  "test name",
		Email: "test@email.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	foundUser, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("found by id ", foundUser)

	foundUser, err = us.ByEmail("test@email.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("found by email ", foundUser)

	user.Name = "updated name"
	err = us.Update(&user)
	if err != nil {
		panic(err)
	}

	foundUser, err = us.ByEmail("test@email.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("found by email after updating name ", foundUser)

	if err := us.Delete(foundUser.ID); err != nil {
		panic(err)
	}
	foundUser, err = us.ByID(foundUser.ID)
	if err != models.ErrNotFound {
		panic("user was not deleted!")
	}
	fmt.Println("found nothing after delete ", foundUser)
}

func expHMAC(){
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}

func expHash(){
	hmac := hash.NewHMAC("my-secret-key")
	fmt.Println("got      =", hmac.Hash("this is my string to hash"))
	fmt.Println("expected = 4waUFc1cnuxoM2oUOJfpGZLGP1asj35y7teuweSFgPY=")
}

func expRememberToken() {
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
	  Name:     "james bond",
	  Email:    "bondjames@spy.com",
	  Password: "doubleohseven",
	}
	err = us.Create(&user)
	if err != nil {
	  panic(err)
	}
	// Verify that the user has a Remember and RememberHash
	fmt.Printf("%+v\n", user)
	if user.Remember == "" {
	  panic("Invalid remember token")
	}
  
	// Now verify that we can lookup a user with that remember
	// token
	user2, err := us.ByRemember(user.Remember)
	if err != nil {
	  panic(err)
	}
	fmt.Printf("%+v\n", *user2)
  }

func main(){
	expRememberToken()
}