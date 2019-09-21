package followfilter

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type FollowFilter struct {
	httpCli    *http.Client
	twitterCli *twitter.Client
}

func NewFollowFilter(oauthConfig *oauth1.Config, oauthToken *oauth1.Token) *FollowFilter {
	httpCli := oauthConfig.Client(oauth1.NoContext, oauthToken)
	twitterCli := twitter.NewClient(httpCli)
	return &FollowFilter{
		httpCli:    httpCli,
		twitterCli: twitterCli,
	}
}

func (ff *FollowFilter) GetFriends() *twitter.Friends {
}
