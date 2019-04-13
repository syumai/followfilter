# followfilter

## Usage of `TechBook` command

* This command prints who joins `Tech book fest (技術書典)` in your Twitter friends (followees).
* Printed results are ordered by booth name.
* You have to get your Twitter app credentials by yourself.

```console
git clone https://github.com/syumai/followfilter
cd followfilter

GO111MODULE=on  \
TWITTER_CONSUMER_KEY=xxx \
TWITTER_CONSUMER_SECRET=xxx \
TWITTER_TOKEN=xxx \
TWITTER_TOKEN_SECRET=xxx \
go run cmd/techbook/main.go
```

## Output

```
か78, syumai, しゅーまい@技術書典か78
...
(All friends (followees) of your account ordered by booth name)
```
