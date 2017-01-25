package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	token   = os.Getenv("SLACK_TOKEN")
	channel = os.Getenv("SLACK_CHANNEL")
)

func join() {
	query := url.Values{}
	query.Add("token", token)
	query.Add("name", channel)
	query.Add("scope", "channels:write")

	resp, err := http.PostForm("https://slack.com/api/channels.join", query)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	log.Printf("Join: Code: %d, Response: %s\n", resp.StatusCode, string(body))
}

func post() {
	query := url.Values{}
	query.Add("token", token)
	query.Add("channel", channel)
	query.Add("scope", "chat:write:user")
	query.Add("as_user", "true")
	query.Add("text", "13:37")

	resp, err := http.PostForm("https://slack.com/api/chat.postMessage", query)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	log.Printf("Post: Code: %d, Response: %s\n", resp.StatusCode, string(body))
}

func main() {
	join()
	post()
}
