package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"time"
	"os"
)

const (
	apiURL = "https://api.github.com"
	issuesEndpoint = "/users/tiagoalvesdulce"
)

type User struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          bool        `json:"hireable"`
	Bio               interface{} `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func getIssues() {
	resp, err := http.Get(apiURL + issuesEndpoint)

	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)

	fmt.Printf("%#v\n", user.Login)

	for _, pair := range os.Environ() {
    fmt.Println(pair)
  }
}

func main() {
	getIssues()
}