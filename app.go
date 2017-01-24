package main

import (
	"os"
	"net/url"
	"net/http"
)

const (
	URL     = "https://slack.com/api/chat.postMessage"
	token   = os.Getenv("SLACK_TOKEN")
	channel = os.Getenv("SLACK_CHANNEL")
	scope   = "chat:write:user"
)


func leet() {
	var query url.Values
	query.Add("token", token)
	query.Add("channel", channel)
	query.Add("scope", scope)
	query.Add("as_user", true)
	query.Add("text", "13:37")

	resp, err := http.NewRequest(http.MethodPost, URL + query.Encode(), nil)
	if err != nil {
		log.Print("Error: " + err.Error())
	}
}

func main() {

}
