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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
)

var cfg ServerConfig

type SlackObj struct {
	Channel    string `json:"channel"`
	Username   string `json:"username"`
	Text       string `json:"text"`
	Icon_emoji string `json:"icon_emoji"`
}

func SendServerWarningMsg(msg string) {
	url := cfg.WebhookUrl
	if url == "" {
		fmt.Println("Your configuration is empty, go to ", conFile, " to modifty it.")
		return
	}

	slack_obj := SlackObj{}
	slack_obj.Channel = cfg.Channel
	slack_obj.Username = cfg.BotName
	slack_obj.Text = msg
	slack_obj.Icon_emoji = cfg.Icon_emoji
	jsonStr, _ := json.Marshal(slack_obj)
	Server_connect(url, "POST", jsonStr)
}

func Server_connect(url string, reqest_mode string, jsonstr []byte) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(reqest_mode, url, bytes.NewBuffer(jsonstr))
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Server log failed. err:", err)
		return
	}

	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
}

func main() {
	cfg = LoadConfig()

	var format = logging.MustStringFormatter("%{level} %{message}")
	logging.SetFormatter(format)
	logging.SetLevel(logging.INFO, "slack-console")

	//Some param here
	var chatMsg string
	rootCmd := &cobra.Command{
		Use:   "slack-console",
		Short: "Post message to your slack channel",
		Run: func(cmd *cobra.Command, args []string) {
			SendServerWarningMsg(chatMsg)
		},
	}
	rootCmd.Flags().StringVarP(&chatMsg, "msg", "m", "some text", "Text content")
	rootCmd.Execute()
}
