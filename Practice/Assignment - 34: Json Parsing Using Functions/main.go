package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	// declare a User slice variable to hold json data
	var users []User
	// declare a Post slice variable to hold json data
	var posts []Post
	// read users.json file content
	bytes, err := readFileContent("users.json")
	if err != nil {
		log.Fatal(err)
	}
	// parse the json data into User slice
	json.Unmarshal(bytes, &users)
	// Print no of users
	fmt.Printf("No of users: %d\n", len(users))
	// read posts.json file content
	bytes, err = readFileContent("posts.json")
	if err != nil {
		log.Fatal(err)
	}
	// parse the json data into Post slice
	json.Unmarshal(bytes, &posts)
	// print no of posts
	fmt.Printf("No of posts: %d\n", len(posts))
	// identify post per user
	var userPostCountMap = map[int]int{}
	for _, p := range posts {
		userPostCountMap[p.UserId] = userPostCountMap[p.UserId] + 1
	}
	// print user details and their post count
	fmt.Printf("%-5s | %-30s | %-3s\n", "ID", "User Name", "No of Posts")
	fmt.Println(strings.Repeat("-", 60))
	for _, u := range users {
		fmt.Printf("%-5d | %-30s | %-3d\n", u.Id, u.Name, userPostCountMap[u.Id])
	}
}

func readFileContent(filePath string) ([]byte, error) {

	// open file to read data
	userJsonTxt, err := os.Open(filePath)
	// handle errors
	if err != nil {
		return nil, err
	}
	// close the file once the operations are done
	defer userJsonTxt.Close()
	// read all file data in bytes
	bytes, err := ioutil.ReadAll(userJsonTxt)
	// handle errors
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
