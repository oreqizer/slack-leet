package main

import (
	"github.com/robfig/cron"
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


func leet() {
	var query url.Values
	query.Add("token", token)
	query.Add("channel", channel)
	query.Add("scope", scope)
	query.Add("as_user", "true")
	query.Add("text", "13:37")

	resp, err := http.NewRequest(http.MethodPost, URL+query.Encode(), nil)
	if err != nil {
		log.Print("Error: " + err.Error())
	}
}

func main() {
	c := cron.New()
	c.AddFunc("0 37 13 * * *", leet)
	c.Start()
}
