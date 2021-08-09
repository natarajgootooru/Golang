package dao

import (
	"encoding/json"

	"github.com/natarajgootooru/Golang/Practice/FinalTask/restutil"
)

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"title"`
	Body   string `json:"body"`
}

func getPostComments(postId string) []Comment {

	// declare a Post slice variable to hold user posts json data
	var comments []Comment
	// convert json bytes into Post struct type
	json.Unmarshal(restutil.DoGET("/posts/"+postId+"/comments"), &comments)

	return comments
}

func getPostCommentCount(postId string, c chan int) {

	c <- len(getPostComments(postId))
}
