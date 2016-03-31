Slack-console: A simple slack message tool in console.
==================
[![GoDoc](https://godoc.org/github.com/kkdai/skiplist?status.svg)](https://godoc.org/github.com/kkdai/slack-console)  [![Build Status](https://travis-ci.org/kkdai/slack-console.svg?branch=master)](https://travis-ci.org/kkdai/slack-console)


A simple console to let you send message to slack specific channel.

Before use this tool, remember you get your Incoming Webhook number in this [page](https://api.slack.com/incoming-webhooks).

Install
---------------
`go get github.com/kkdai/slack-console`


Change the default setting
---------------

Change the configuration file which will generate when your first time launch this application.

```
{
	"WebhookUrl" : "",	
	"BotName" : "",
	"Channel" : "000",
	"Icon_emoji" : ":octocat:"
}
```
Detail explaination as follow:

- `WebhookUrl`: Get url from https://api.slack.com/incoming-webhooks
- `BotName`: The name display of your bot
- `Channel`: The channel your bot want to push message
- `Icon_emoji`: The emoji of your bot, such `:octocat:`. Refer http://www.emoji-cheat-sheet.com/ for more

Fill all infor base on this [page](https://api.slack.com/incoming-webhooks)

Usage
---------------

```
//Pass message to slack
slack-console -m "test msg"

```

### Inspired By:

- [Slack incoming webhook](https://api.slack.com/incoming-webhooks)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.