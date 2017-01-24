package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	URLpost   = "https://slack.com/api/chat.postMessage"
	URLjoin   = "https://slack.com/api/channels.join"
	joinScope = "channels:write"
	postScope = "chat:write:user"
)

var (
	token   = os.Getenv("SLACK_TOKEN")
	channel = os.Getenv("SLACK_CHANNEL")
)

func main() {
	joinQuery := url.Values{}
	joinQuery.Add("token", token)
	joinQuery.Add("name", channel)
	joinQuery.Add("scope", joinScope)

	join, err := http.PostForm(URLjoin, joinQuery)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	defer join.Body.Close()

	body, err := ioutil.ReadAll(join.Body)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	log.Printf("Join: Code: %d, Response: %s\n", join.StatusCode, string(body))

	postQuery := url.Values{}
	postQuery.Add("token", token)
	postQuery.Add("channel", channel)
	postQuery.Add("postScope", postScope)
	postQuery.Add("as_user", "true")
	postQuery.Add("text", "13:37")

	post, err := http.PostForm(URLpost, postQuery)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	defer post.Body.Close()

	body, err = ioutil.ReadAll(post.Body)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	log.Printf("Post: Code: %d, Response: %s\n", post.StatusCode, string(body))
}
