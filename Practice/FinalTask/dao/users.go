package dao

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/natarajgootooru/Golang/Practice/FinalTask/restutil"
)

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Address  Address
	Company  Company
}

func (u User) printUserShortDetails() {
	fmt.Printf("%-4d | %-25s | %-15s\n", u.Id, u.Name, u.Email)
}

func (u User) printUserDetails() {

	fmt.Printf("Id: %d | Name: %s | UserName: %s | Email: %s\n", u.Id, u.Name, u.UserName, u.Email)
	fmt.Printf("Phone: %s | Website: %s\n", u.Phone, u.Website)
}

func getUserList() []User {

	// declare a User slice variable to hold users json data
	var users []User
	// convert json bytes into User struct type
	json.Unmarshal(restutil.DoGET("/users"), &users)

	return users
}

func getUserDetails(userId string) User {

	var user User
	json.Unmarshal(restutil.DoGET("/users/"+userId), &user)
	return user
}

func PrintAllUserOnConsole() {

	users := getUserList()
	fmt.Println(strings.Repeat("-", 65))
	fmt.Printf("%-4s | %-25s | %-15s\n", "Id", "Name", "Email Id")
	fmt.Println(strings.Repeat("-", 65))
	for _, u := range users {
		u.printUserShortDetails()
	}
	fmt.Println(strings.Repeat("-", 65))
}

func PrintUserStats(userId string) {

	user := getUserDetails(userId)
	if user.Name == "" {
		log.Fatal("User Id " + userId + " does not exists")
	}
	postCount, commentCount := GetUserPostCommentStats(userId)
	user.printUserDetails()
	fmt.Printf("Total User Posts: %d\n", postCount)
	fmt.Printf("Total Comments Received: %d\n", commentCount)
}
