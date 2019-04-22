package main

import (
	"github.com/nlopes/slack"
)

var api = slack.New("TOKEN")

func main() {
  api.PostMessage("#general",
            slack.MsgOptionText("Hello @rbmrclo, say hi to @everyone in #general", false),
            slack.MsgOptionParse(true))

  api.PostMessage("#general",
            slack.MsgOptionText("Hello @rbmrclo, say hi to @everyone in #general", false),
            slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{Parse: "none"}))

  api.PostMessage("#general",
            slack.MsgOptionText("Hello @rbmrclo, say hi to @everyone in #general", false),
            slack.MsgOptionPostMessageParameters(slack.PostMessageParameters{Parse: "full"}))
}
