package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	boothNameRegex = regexp.MustCompile(`([あ-んス][\d]{2})`)
	consumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	token          = os.Getenv("TWITTER_TOKEN")
	tokenSecret    = os.Getenv("TWITTER_TOKEN_SECRET")
)

type TechBookUser struct {
	BoothName  string
	Name       string
	ScreenName string
}

func (tu TechBookUser) String() string {
	return fmt.Sprintf("%s, %s, %s", tu.BoothName, tu.ScreenName, tu.Name)
}

func main() {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	oauthToken := oauth1.NewToken(token, tokenSecret)

	httpClient := config.Client(oauth1.NoContext, oauthToken)
	client := twitter.NewClient(httpClient)

	var (
		cur   int64
		users []twitter.User
	)
	for {
		f, _, err := client.Friends.List(&twitter.FriendListParams{
			Cursor: cur,
			Count:  200,
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, u := range f.Users {
			if !strings.Contains(u.Name, "技術書") {
				continue
			}
			users = append(users, u)
		}
		cur = f.NextCursor
		if cur == 0 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	var tus []TechBookUser
	for _, u := range users {
		nameParts := strings.SplitAfter(u.Name, "技術書")
		matches := boothNameRegex.FindStringSubmatch(nameParts[1])
		if len(matches) == 0 {
			continue
		}
		tus = append(tus, TechBookUser{
			BoothName:  matches[0],
			Name:       u.Name,
			ScreenName: u.ScreenName,
		})
	}

	sort.Slice(tus, func(i, j int) bool {
		return tus[i].BoothName < tus[j].BoothName
	})

	for _, tu := range tus {
		fmt.Println(tu.String())
	}
}
