package dao

import (
	"encoding/json"
	"strconv"

	"github.com/natarajgootooru/Golang/Practice/FinalTask/restutil"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getUserPosts(userId string) []Post {

	// declare a Post slice variable to hold user posts json data
	var posts []Post
	// convert json bytes into Post struct type
	json.Unmarshal(restutil.DoGET("/users/"+userId+"/posts"), &posts)

	return posts
}

func GetUserPostCommentStats(userId string) (int, int) {

	commentCount := 0
	posts := getUserPosts(userId)
	c := make(chan int)
	for _, p := range posts {
		s := strconv.Itoa(p.Id)
		go getPostCommentCount(s, c)
	}
	chanCount := 0
	for i := range c {
		commentCount += i
		chanCount++
		if len(posts) == chanCount {
			close(c)
		}
	}
	return len(posts), commentCount
}
