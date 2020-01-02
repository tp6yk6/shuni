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
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("7f429e5e0a1c48de9ffacb668eda4b56"), os.Getenv("fyHN9FqU0g2pOfX3d1jG0ZHxdU75biicfisoJTGBpuW/WwWghm1DwipvW9hG4fEo5GJGSGXKAsccP1wsqek+LEf08R0U5qIW1q7sK8DuYzME6erVbzOXmNOTLF+X2FvXJh6EXbVLpEt/RzXcLVP1qwdB04t89/1O/w1cDnyilFU="))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				quota, err := bot.GetMessageQuota().Do()
				if _, err = bot.ReplyMessage(event.ReplyToken,
				 linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain 
				 	message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
			}
		}
	}
}
}

