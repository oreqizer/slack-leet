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

	var body []byte
	_, err = resp.Body.Read(body)
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}
	log.Printf("Code: %d, Response: %s\n", resp.StatusCode, string(body))
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(365)  // do it for a year

	leet := &Leet{wg}
	c := cron.New()
	c.AddJob("0 41 17 * * *", leet)
	c.Start()

	wg.Wait()
}
