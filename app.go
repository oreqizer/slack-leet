package main

import (
	"github.com/robfig/cron"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)


const (
	URL     = "https://slack.com/api/chat.postMessage"
	scope   = "chat:write:user"
)

var (
	token   = os.Getenv("SLACK_TOKEN")
	channel = os.Getenv("SLACK_CHANNEL")
)


type Leet struct {
	wg *sync.WaitGroup
}

func (l *Leet) Run() {
	defer l.wg.Done()

	var query url.Values
	query.Add("token", token)
	query.Add("channel", channel)
	query.Add("scope", scope)
	query.Add("as_user", "true")
	query.Add("text", "13:37")

	resp, err := http.NewRequest(http.MethodPost, URL+query.Encode(), nil)
	if err != nil {
		log.Println("Error: " + err.Error())
	}
	defer resp.Body.Close()

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		log.Println("Error: " + err.Error())
	}
	log.Println("Response: " + string(body))
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(365)  // do it for a year

	leet := &Leet{wg}
	c := cron.New()
	c.AddJob("0 37 13 * * *", leet)
	c.Start()

	wg.Wait()
}
