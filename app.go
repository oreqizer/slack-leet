package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)


const (
	URL     = "https://slack.com/api/chat.postMessage"
	scope   = "chat:write:user"
)

var (
	token   = os.Getenv("SLACK_TOKEN")
	channel = os.Getenv("SLACK_CHANNEL")
)

func main() {
	query := url.Values{}
	query.Add("token", token)
	query.Add("channel", channel)
	query.Add("scope", scope)
	query.Add("as_user", "true")
	query.Add("text", "13:37")

	resp, err := http.PostForm(URL, query)
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
	log.Printf("Code: %d, Response: %s\n", resp.StatusCode, string(body))
}
