// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

var conFile string

const RPC_SERVER_VERSION string = "v1.0"

const DefaultConfigFile string = `{
	"WebhookUrl" : "",
	"BotName" : "",
	"Channel" : "000",
	"Icon_emoji" : ":octocat:"
}`

type ServerConfig struct {
	//Server Initialize Variables
	WebhookUrl string //Get url from https://api.slack.com/incoming-webhooks
	BotName    string //The name display of your bot
	Channel    string //The channel your bot want to push message
	Icon_emoji string //The emoji of your bot, such `:octocat:`. Refer http://www.emoji-cheat-sheet.com/ for more
}

func LoadConfig() ServerConfig {
	usr, _ := user.Current()
	conFile = fmt.Sprintf("%v/slack-console.json", usr.HomeDir)
	if _, err := os.Stat(conFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(conFile, []byte(DefaultConfigFile), 0644)
		if err != nil {
			log.Panic("Cannot create config file")
		} else {
			fmt.Println("No configuration file, create a default config in ", conFile)
		}
	}

	file, _ := os.Open(conFile)
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := ServerConfig{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(" No config file.")
	}
	//log.Println(configurtion)
	return configuration
}
